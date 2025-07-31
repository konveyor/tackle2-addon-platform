package helm

import (
	liberr "github.com/jortel/go-utils/error"
	hp "github.com/konveyor/asset-generation/pkg/providers/generators/helm"
	"github.com/konveyor/tackle2-hub/api"
)

var (
	Wrap = liberr.Wrap
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
		err = Wrap(err)
		return
	}
	return
}
