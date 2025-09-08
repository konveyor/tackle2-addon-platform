package main

import (
	"path"
	"testing"

	"github.com/goccy/go-json"
	"github.com/konveyor/tackle2-hub/api"
	"github.com/onsi/gomega"
)

func TestManifestMerge(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	v := Values{Manifest: api.Map{
		"a": api.Map{
			"b": 2,
		},
		"c": 25,
	}}

	d := api.Map{
		"manifest.a.b":   100,
		"manifest.a.n.x": 200,
		"manifest.a.n.y": 200,
		"manifest.c":     300,
		"port":           8080,
	}

	injected, _ := v.inject(d)

	v2 := Values{Manifest: api.Map{
		"a": api.Map{
			"b": 100,
			"n": api.Map{
				"x": 200,
				"y": 200,
			},
		},
		"c": 300,
	}}
	expected := v2.asMap()
	expected["port"] = 8080

	b, _ := json.Marshal(expected)
	var mA map[string]any
	_ = json.Unmarshal(b, &mA)
	b, _ = json.Marshal(injected)
	var mB map[string]any
	_ = json.Unmarshal(b, &mB)

	g.Expect(mA).To(gomega.BeEquivalentTo(mB))
}

func TestAssetDir(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	rootDir := "/tmp"
	a := Generate{}

	gen := &api.Generator{}
	gen.Name = "templates"
	gen.Repository = &api.Repository{Path: "test"}
	p := a.genAssetDir("/tmp", gen)
	g.Expect(p).To(gomega.Equal(path.Join(rootDir, gen.Name, gen.Repository.Path)))

	gen = &api.Generator{}
	gen.ID = 18
	gen.Repository = &api.Repository{Path: "test"}
	p = a.genAssetDir("/tmp", gen)
	g.Expect(p).To(gomega.Equal(path.Join(rootDir, "18", gen.Repository.Path)))

	gen.Repository = &api.Repository{URL: "http://host:8080/r/dog.git"}
	p = a.genAssetDir("/tmp", gen)
	g.Expect(p).To(gomega.Equal(path.Join(rootDir, "18", "dog.git")))

	gen.Repository = &api.Repository{URL: "http://host:8080"}
	p = a.genAssetDir("/tmp", gen)
	g.Expect(p).To(gomega.Equal(path.Join(rootDir, "18", "host")))

	gen.Repository = &api.Repository{URL: ""}
	p = a.genAssetDir("/tmp", gen)
	g.Expect(p).To(gomega.Equal(path.Join(rootDir, "18", "templates")))

	gen.Repository = &api.Repository{Path: "."}
	p = a.genAssetDir("/tmp", gen)
	g.Expect(p).To(gomega.Equal(path.Join(rootDir, "18", "templates")))
	
	gen.Repository = &api.Repository{Path: "/"}
	p = a.genAssetDir("/tmp", gen)
	g.Expect(p).To(gomega.Equal(path.Join(rootDir, "18", "templates")))
}
