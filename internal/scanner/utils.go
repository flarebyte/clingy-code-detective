package scanner

import "strings"

// categoryToFiles maps categories to supported files.
var categoryToFiles = map[string][]string{
	"dart":   {"pubspec.yaml"},
	"go":     {"go.mod"},
	"node":   {"package.json"},
	"python": {"requirements.txt"},
}

// aliasToCategory maps aliases to canonical categories.
var aliasToCategory = map[string]string{
	"js":   "node",
	"ts":   "node",
	"node": "node", // for consistency
}

// isFileRequired returns true if the filename is required based on includes.
func IsFileRequired(filename string, includes []string) bool {
	// Build set of allowed files.
	allowedFiles := make(map[string]struct{})

	if len(includes) == 0 {
		// If includes is empty, all supported files are allowed.
		for _, files := range categoryToFiles {
			for _, f := range files {
				allowedFiles[f] = struct{}{}
			}
		}
	} else {
		for _, inc := range includes {
			// Resolve alias if needed.
			category := resolveCategory(inc)
			if files, ok := categoryToFiles[category]; ok {
				for _, f := range files {
					allowedFiles[f] = struct{}{}
				}
			}
		}
	}

	_, ok := allowedFiles[strings.ToLower(filename)]
	return ok
}

// resolveCategory resolves an alias to its canonical category.
func resolveCategory(alias string) string {
	if cat, ok := aliasToCategory[strings.ToLower(alias)]; ok {
		return cat
	}
	// If not an alias, treat it as its own category.
	return strings.ToLower(alias)
}
