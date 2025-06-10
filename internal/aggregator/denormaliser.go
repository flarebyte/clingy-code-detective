package aggregator

import "github.com/flarebyte/clingy-code-detective/internal/parser"

// Denormalise the dependency file into an array of FlatDependency
func DenormaliseDependencyFile(file parser.DependencyFile) []FlatDependency {
	if file.Err != nil || len(file.Dependencies) == 0 {
		return []FlatDependency{}
	}

	flatDeps := make([]FlatDependency, 0, len(file.Dependencies))
	for _, dep := range file.Dependencies {
		flatDeps = append(flatDeps, FlatDependency{
			Name:     dep.Name,
			Version:  dep.Version,
			Category: dep.Category,
			Path:     file.Path,
		})
	}

	return flatDeps
}
