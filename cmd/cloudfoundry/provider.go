package cloudfoundry

import (
	"path/filepath"

	cf "github.com/cloudfoundry/go-cfclient/v3/config"
	cfp "github.com/konveyor/asset-generation/pkg/providers/discoverers/cloud_foundry"
	hub "github.com/konveyor/tackle2-hub/addon"
	"github.com/konveyor/tackle2-hub/api"
	"github.com/konveyor/tackle2-hub/api/jsd"
	"github.com/konveyor/tackle2-hub/migration/json"
	"github.com/pkg/errors"
)

var (
	addon = hub.Addon
)

// Provider is a cloudfoundry provider.
type Provider struct {
	URL      string
	Identity *api.Identity
}

// Use identity.
func (p *Provider) Use(identity *api.Identity) {
	p.Identity = identity
}

// Fetch the manifest for the application.
func (p *Provider) Fetch(application *api.Application) (m *api.Manifest, err error) {
	if application.Coordinates == nil {
		err = errors.Errorf("Coordinates required.")
		return
	}
	coordinates := Coordinates{}
	err = application.Coordinates.As(&coordinates)
	if err != nil {
		return
	}
	ref := cfp.AppReference{
		SpaceName: coordinates.Space,
		AppName:   coordinates.Name,
	}
	client, err := p.client(ref.SpaceName)
	if err != nil {
		return
	}
	manifest, err := client.Discover(ref)
	if err != nil {
		return
	}
	m = &api.Manifest{}
	m.Content = manifest.Content
	m.Secret = manifest.Secret
	return
}

// Find applications on the platform.
func (p *Provider) Find(filter api.Map) (found []api.Application, err error) {
	f := Filter{}
	err = filter.As(&f)
	if err != nil {
		return
	}
	client, err := p.client(f.Spaces...)
	if err != nil {
		return
	}
	spaces, err := client.ListApps()
	if err != nil {
		return
	}
	schema, err := addon.Schema.Find("platform", "cloudfoundry", "coordinates")
	if err != nil {
		return
	}
	for space, applications := range spaces {
		if !f.MatchSpace(space) {
			continue
		}
		for _, ref := range applications {
			appRef := ref.(cfp.AppReference)
			if !f.MatchName(appRef.AppName) {
				continue
			}
			r := api.Application{}
			r.Name = appRef.AppName
			r.Coordinates = &jsd.Document{
				Schema: schema.Name,
				Content: json.Map{
					"name":  appRef.AppName,
					"space": space,
				},
			}
			found = append(found, r)
		}
	}
	return
}

// client returns a cloudfoundry client.
func (p *Provider) client(spaces ...string) (client *cfp.CloudFoundryProvider, err error) {
	options := []cf.Option{
		cf.SkipTLSValidation(),
	}
	if p.Identity != nil {
		options = append(
			options,
			cf.UserPassword(
				p.Identity.User,
				p.Identity.Password))
	}
	cfConfig, err := cf.New(p.URL, options...)
	if err != nil {
		return
	}
	pConfig := &cfp.Config{
		CloudFoundryConfig: cfConfig,
		SpaceNames:         spaces,
	}
	client, err = cfp.New(pConfig, &addon.Log, true)
	if err != nil {
		return
	}
	return
}

// Coordinates - platform coordinates.
type Coordinates struct {
	Space string `json:"space"`
	Name  string `json:"name"`
}

// Filter applications.
type Filter struct {
	Spaces []string `json:"spaces"`
	Names  []string `json:"names"`
}

// MatchSpace returns true when the application name matches the filter.
func (f *Filter) MatchSpace(name string) (match bool) {
	for _, s := range f.Spaces {
		match = s == name
		if match {
			break
		}
	}
	return
}

// MatchName returns true when the name matches the filter.
// The name may be a glob.
func (f *Filter) MatchName(name string) (match bool) {
	var err error
	if len(f.Names) == 0 {
		match = true
		return
	}
	for _, pattern := range f.Names {
		match, err = filepath.Match(pattern, name)
		if err != nil {
			addon.Log.Error(err, "Invalid glob pattern", "pattern", pattern)
			continue
		}
		if match {
			break
		}
	}
	return
}
