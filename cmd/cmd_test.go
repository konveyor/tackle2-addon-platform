package main

import (
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
