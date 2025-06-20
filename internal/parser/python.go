package parser

import (
	"strings"
)

type pythonParser struct{}

func (p pythonParser) Parse(content []byte) ([]Dependency, error) {

	lines := strings.Split(string(content), "\n")
	deps := make([]Dependency, 0, len(lines)) // ensures non-nil slice
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
