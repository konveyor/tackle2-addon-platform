package helm

import (
	hp "github.com/konveyor/asset-generation/pkg/providers/generators/helm"
	"github.com/konveyor/tackle2-hub/api"
)

type Files = map[string]string

// Engine is a helm template engine.
type Engine struct {
}

// Generate generates assets.
// Returns a list of files (content).
func (g *Engine) Generate(templateDir string, values api.Map) (files Files, err error) {
	files = make(Files)
	config := hp.Config{
		ChartPath: templateDir,
		Values:    values,
	}
	provider := hp.New(config)
	files, err = provider.Generate()
	if err != nil {
		return
	}
	return
}
