package parser

import (
	"encoding/json"
	"sort"
)

type nodeParser struct{}

type packageJSON struct {
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

func (p nodeParser) Parse(content []byte) ([]Dependency, error) {
	var pkg packageJSON
	if err := json.Unmarshal(content, &pkg); err != nil {
		return nil, err
	}

	collectDeps := func(m map[string]string, cat string) []Dependency {
		var deps []Dependency
		keys := make([]string, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, name := range keys {
			deps = append(deps, Dependency{Name: name, Version: m[name], Category: cat})
		}
		return deps
	}

	var deps []Dependency
	deps = append(deps, collectDeps(pkg.Dependencies, "prod")...)
	deps = append(deps, collectDeps(pkg.DevDependencies, "dev")...)

	return deps, nil
}
