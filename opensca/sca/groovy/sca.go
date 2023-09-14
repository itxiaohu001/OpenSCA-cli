package groovy

import (
	"context"

	"github.com/xmirrorsecurity/opensca-cli/opensca/model"
	"github.com/xmirrorsecurity/opensca-cli/opensca/sca/filter"
)

type Sca struct{}

func (sca Sca) Language() model.Language {
	return model.Lan_Java
}

func (sca Sca) Filter(relpath string) bool {
	return filter.GroovyGradle(relpath) || filter.GroovyFile(relpath)
}

func (sca Sca) Sca(ctx context.Context, parent *model.File, files []*model.File) []*model.DepGraph {

	roots := GradleTree(ctx, parent)
	if len(roots) == 0 {
		roots = ParseGradle(files)
	}

	for _, f := range files {
		if filter.GroovyFile(f.Relpath()) {
			root := ParseGroovy(f)
			if root != nil {
				roots = append(roots, root)
			}
		}
	}

	return roots
}
