type goModParser struct{}

func (p goModParser) Parse(path string) ([]Dependency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	f, err := modfile.Parse(path, data, nil)
	if err != nil {
		return nil, err
	}

	var deps []Dependency
	for _, r := range f.Require {
		cat := "prod"
		if r.Indirect {
			cat = "indirect"
		}
		deps = append(deps, Dependency{
			Name:     r.Mod.Path,
			Version:  r.Mod.Version,
			Category: cat,
		})
	}
	return deps, nil
}
