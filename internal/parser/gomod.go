package parser

import (
	"bufio"
	"bytes"
	"strings"
)

type goModParser struct{}

// Parse extracts dependencies from a go.mod file. It identifies "prod" or "dev"
// based on the presence of "// indirect" comments in the require block.
func (p goModParser) Parse(content []byte) ([]Dependency, error) {
	deps := make([]Dependency, 0)
	scanner := bufio.NewScanner(bytes.NewReader(content))

	inRequireBlock := false

	for scanner.Scan() {
		rawLine := scanner.Text()
		line := strings.TrimSpace(rawLine)

		if strings.HasPrefix(line, "require (") {
			inRequireBlock = true
			continue
		}
		if inRequireBlock && line == ")" {
			inRequireBlock = false
			continue
		}

		var depLine string
		if inRequireBlock {
			depLine = line
		} else if strings.HasPrefix(line, "require ") {
			depLine = strings.TrimSpace(strings.TrimPrefix(line, "require "))
		} else {
			continue
		}

		category := "prod"
		if strings.Contains(depLine, "// indirect") {
			category = "dev"
			depLine = strings.Split(depLine, "//")[0]
		} else if idx := strings.Index(depLine, "//"); idx != -1 {
			depLine = depLine[:idx]
		}

		fields := strings.Fields(depLine)
		if len(fields) != 2 {
			continue
		}

		deps = append(deps, Dependency{
			Name:     fields[0],
			Version:  fields[1],
			Category: category,
		})
	}

	return deps, scanner.Err()
}
