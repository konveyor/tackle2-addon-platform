package main

import (
	"errors"

	"github.com/konveyor/tackle2-hub/api"
)

// Action provides and addon action.
type Action interface {
	Run(*Data) error
}

// NewAction returns and action based the data provided.
func NewAction(d *Data) (a Action, err error) {
	switch d.Action {
	case "fetch":
		a = &Fetch{}
	case "import":
		a = &Import{}
	case "generate":
		a = &Generate{}
	default:
		err = errors.New("action not supported")
	}
	return
}

// BaseAction provides base functionality.
type BaseAction struct {
	application api.Application
	platform    api.Platform
}

// setApplication fetches and sets `application` referenced by the task.
// The associated `platform` will be set when as appropriate.
func (r *BaseAction) setApplication() (err error) {
	defer func() {
		if err != nil {
			if errors.Is(err, &api.NotFound{}) {
				err = nil
			}
		}
	}()
	app, err := addon.Task.Application()
	if err == nil {
		r.application = *app
	} else {
		return
	}
	if app.Platform == nil {
		return
	}
	p, err := addon.Platform.Get(app.Platform.ID)
	if err == nil {
		r.platform = *p
	}
	return
}

// setPlatform fetches and sets `platform` referenced by the task.
func (r *BaseAction) setPlatform() (err error) {
	defer func() {
		if err != nil {
			if errors.Is(err, &api.NotFound{}) {
				err = nil
			}
		}
	}()
	p, err := addon.Task.Platform()
	if err == nil {
		r.platform = *p
	}
	return
}
