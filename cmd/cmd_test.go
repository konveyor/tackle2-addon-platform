package main

import (
	"testing"

	"github.com/konveyor/tackle2-hub/api"
	"github.com/onsi/gomega"
)

func TestManifestMerge(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	m := api.Map{
		"a": Map{
			"b": 2,
		},
		"c": 25,
	}
	d := api.Map{
		"a.b":   100,
		"a.n.x": 200,
		"a.n.y": 200,
		"c":     300,
	}

	gen := Generate{}
	gen.inject(m, d)

	expected := api.Map{
		"a": Map{
			"b": 100,
			"n": Map{
				"x": 200,
				"y": 200,
			},
		},
		"c": 300,
	}
	g.Expect(expected).To(gomega.BeEquivalentTo(m))
}
