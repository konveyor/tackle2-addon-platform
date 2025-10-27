package main

import "github.com/konveyor/tackle2-hub/api"

const (
	TagSource   = "platform-discovery"
	TagCategory = "Platform"
	Tag         = "Cloud Foundry"
)

// Fetch application manifest action.
type Fetch struct {
	BaseAction
}

// Run executes the action.
func (a *Fetch) Run(d *Data) (err error) {
	err = a.setApplication()
	if err != nil {
		return
	}
	addon.Activity(
		"[Fetch] Fetch manifest for application (id=%d): %s",
		a.application.ID,
		a.application.Name)
	addon.Activity(
		"[Fetch] Using platform (id=%d): %s",
		a.platform.ID,
		a.platform.Name)
	provider, err := a.selectProvider(a.platform.Kind)
	if err != nil {
		return
	}
	err = a.setTag()
	if err != nil {
		return
	}
	err = a.fetch(provider, &a.application)
	return
}

// fetch manifest.
func (a *Fetch) fetch(p Provider, app *api.Application) (err error) {
	manifest, err := p.Fetch(app)
	if err != nil {
		return
	}
	manifest.Application.ID = app.ID
	err = addon.Manifest.Create(manifest)
	if err == nil {
		addon.Activity(
			"Manifest (id=%d) created.",
			manifest.ID)
	}
	return
}

// setTag replaces the platform tag.
func (a *Fetch) setTag() (err error) {
	cat := &api.TagCategory{Name: TagCategory}
	err = addon.TagCategory.Ensure(cat)
	if err != nil {
		return
	}
	tag := &api.Tag{
		Category: api.Ref{ID: cat.ID},
		Name:     Tag,
	}
	err = addon.Tag.Ensure(tag)
	if err != nil {
		return
	}
	appTags := addon.Application.Tags(a.application.ID)
	appTags.Source(TagSource)
	err = appTags.Replace([]uint{tag.ID})
	if err != nil {
		addon.Activity("Application tagged.")
	}
	return
}
