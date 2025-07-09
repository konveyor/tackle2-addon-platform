package main

import (
	"errors"

	cf "github.com/konveyor/tackle2-addon-platform/cmd/cloudfoundry"
	"github.com/konveyor/tackle2-hub/api"
)

// Import applications action.
type Import struct {
	BaseAction
}

// Run executes the action.
func (a *Import) Run(d *Data) (err error) {
	var applications []api.Application
	err = a.setPlatform()
	if err != nil {
		return
	}
	addon.Activity(
		"[Import] Using platform (id=%d): %s",
		a.platform.ID,
		a.platform.Name)
	switch a.platform.Kind {
	case "cloudfoundry":
		applications, err = a.cloudfoundry(d.Filter)
	default:
		err = errors.New("platform.kind not supported")
		return
	}
	addon.Activity(
		"[Import] Found %d applications.",
		len(applications))
	for _, app := range applications {
		err := addon.Application.Create(&app)
		if err != nil {
			addon.Errorf(
				"warn",
				"[Import] Application: %s, create failed: %s",
				app.Name,
				err.Error())
			continue
		}
		addon.Activity(
			"[Import] Application: %s, created.",
			app.Name)
	}
	return
}

// cloudfoundry implementation.
func (a *Import) cloudfoundry(filter api.Map) (applications []api.Application, err error) {
	p := cf.Provider{
		URL: a.platform.URL,
	}
	if a.platform.Identity.ID != 0 {
		p.Identity, err = addon.Identity.Get(a.platform.Identity.ID)
		if err == nil {
			addon.Activity(
				"[Import] Using credentials (id=%d): %s",
				p.Identity.ID,
				p.Identity.Name)
		} else {
			return
		}
	}
	applications, err = p.Find(filter)
	return
}
