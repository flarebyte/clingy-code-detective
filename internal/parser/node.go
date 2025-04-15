type nodeParser struct{}

type packageJSON struct {
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

func (p nodeParser) Parse(path string) ([]Dependency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var pkg packageJSON
	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, err
	}

	var deps []Dependency
	for name, ver := range pkg.Dependencies {
		deps = append(deps, Dependency{Name: name, Version: ver, Category: "prod"})
	}
	for name, ver := range pkg.DevDependencies {
		deps = append(deps, Dependency{Name: name, Version: ver, Category: "dev"})
	}

	return deps, nil
}
