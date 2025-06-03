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
	var deps []Dependency
	scanner := bufio.NewScanner(bytes.NewReader(content))

	inRequireBlock := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "require (") {
			inRequireBlock = true
			continue
		}
		if inRequireBlock && line == ")" {
			inRequireBlock = false
			continue
		}

		if !inRequireBlock && !strings.HasPrefix(line, "require ") {
			continue
		}

		var depLine string
		if inRequireBlock {
			depLine = line
		} else {
			depLine = strings.TrimPrefix(line, "require ")
		}

		// Remove inline comments, but detect `// indirect`
		category := "prod"
		if strings.Contains(depLine, "// indirect") {
			category = "dev"
			depLine = strings.Split(depLine, "//")[0]
		} else if idx := strings.Index(depLine, "//"); idx != -1 {
			depLine = depLine[:idx]
		}

		fields := strings.Fields(depLine)
		if len(fields) != 2 {
			continue // malformed line; skip
		}

		deps = append(deps, Dependency{
			Name:     fields[0],
			Version:  fields[1],
			Category: category,
		})
	}

	return deps, scanner.Err()
}
