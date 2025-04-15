type dartParser struct{}

type pubspecYAML struct {
	Dependencies    map[string]interface{} `yaml:"dependencies"`
	DevDependencies map[string]interface{} `yaml:"dev_dependencies"`
}

func (p dartParser) Parse(path string) ([]Dependency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var spec pubspecYAML
	if err := yaml.Unmarshal(data, &spec); err != nil {
		return nil, err
	}

	parseMap := func(m map[string]interface{}, cat string) []Dependency {
		var deps []Dependency
		for name, val := range m {
			switch v := val.(type) {
			case string:
				deps = append(deps, Dependency{Name: name, Version: v, Category: cat})
			default:
				deps = append(deps, Dependency{Name: name, Version: "", Category: cat})
			}
		}
		return deps
	}

	var deps []Dependency
	deps = append(deps, parseMap(spec.Dependencies, "prod")...)
	deps = append(deps, parseMap(spec.DevDependencies, "dev")...)

	return deps, nil
}
