package main

import (
	"github.com/konveyor/tackle2-hub/api"
)

// Import applications action.
type Import struct {
	BaseAction
}

// Run executes the action.
func (a *Import) Run(d *Data) (err error) {
	err = a.setPlatform()
	if err != nil {
		return
	}
	addon.Activity(
		"[Import] Using platform (id=%d): %s",
		a.platform.ID,
		a.platform.Name)
	provider, err := a.selectProvider(a.platform.Kind)
	if err != nil {
		return
	}
	created := make([]api.Application, 0)
	applications, err := provider.Find(d.Filter)
	if err != nil {
		return
	}
	addon.Activity(
		"[Import] Found %d applications.",
		len(applications))
	for _, app := range applications {
		app.Platform = &api.Ref{
			ID: a.platform.ID,
		}
		err := addon.Application.Create(&app)
		if err != nil {
			addon.Errorf(
				"warn",
				"[Import] Application: %s, create failed: %s",
				app.Name,
				err.Error())
			continue
		}
		created = append(created, app)
		addon.Activity(
			"[Import] Application: %s, created.",
			app.Name)
		err = (&Fetch{}).fetch(provider, &app)
		if err != nil {
			addon.Errorf(
				"warn",
				"[Import] Application: %s, fetch manifest failed: %s",
				app.Name,
				err.Error())
			continue
		}
	}
	return
}
