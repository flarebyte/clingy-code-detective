type pythonParser struct{}

func (p pythonParser) Parse(path string) ([]Dependency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var deps []Dependency
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		// Rough split for `package==version` style
		parts := strings.SplitN(line, "==", 2)
		if len(parts) == 2 {
			deps = append(deps, Dependency{Name: parts[0], Version: parts[1], Category: "prod"})
		} else {
			deps = append(deps, Dependency{Name: line, Version: "", Category: "prod"})
		}
	}
	return deps, nil
}
