package main

import (
	"os"
	"path"

	"github.com/konveyor/tackle2-addon/ssh"
	hub "github.com/konveyor/tackle2-hub/addon"
	"github.com/konveyor/tackle2-hub/api"
	"github.com/konveyor/tackle2-hub/nas"
)

var (
	addon       = hub.Addon
	SourceDir   = ""
	TemplateDir = ""
	AssetDir    = ""
	Dir         = ""
)

func init() {
	Dir, _ = os.Getwd()
	SourceDir = path.Join(Dir, "source")
	TemplateDir = path.Join(Dir, "templates")
	AssetDir = path.Join(Dir, "assets")
}

// Data Addon data passed in the secret.
type Data struct {
	// Action (fetch|import|generate)
	Action string `json:"action"`
	// Import
	// Filter applications.
	Filter api.Map `json:"filter"`
	// Asset Generation
	// Profiles
	Profiles Profiles `json:"profiles"`
	// Params generator params
	Params api.Map `json:"params"`
}

// main
func main() {
	addon.Run(func() (err error) {
		addon.Activity("SourceDir: %s", SourceDir)
		addon.Activity("TemplateDir: %s", TemplateDir)
		addon.Activity("AssetDir: %s", AssetDir)
		//
		// Get the addon data associated with the task.
		d := &Data{}
		err = addon.DataWith(d)
		if err != nil {
			return
		}
		//
		// Create directories.
		for _, dir := range []string{SourceDir, TemplateDir, AssetDir} {
			err = nas.MkDir(dir, 0755)
			if err != nil {
				err = wrap(err)
				return
			}
		}
		//
		// SSH
		agent := ssh.Agent{}
		err = agent.Start()
		if err != nil {
			return
		}
		//
		// action
		action, err := NewAction(d)
		if err != nil {
			return
		}
		//
		// Run action
		err = action.Run(d)
		if err != nil {
			return
		}

		addon.Activity("Done.")

		return
	})
}
