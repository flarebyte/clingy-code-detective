package parser

import (
	"errors"
	"os"
	"path/filepath"
)

// ParseDependencyFile opens and parses a file using the appropriate parser.
func ParseDependencyFile(path string) DependencyFile {
	var parser Parser

	switch filepath.Base(path) {
	case "package.json":
		parser = nodeParser{}
	case "pubspec.yaml":
		parser = dartParser{}
	case "go.mod":
		parser = goModParser{}
	case "requirements.txt":
		parser = pythonParser{}
	default:
		return DependencyFile{Path: path, Err: errors.New("unsupported file type")}
	}

	content, ferr := os.ReadFile(path)
	if ferr != nil {
		return DependencyFile{
			Path: path,
			Err:  ferr,
		}
	}

	deps, err := parser.Parse(content)
	return DependencyFile{
		Path:         path,
		Dependencies: deps,
		Err:          err,
	}
}

func ProduceDependencyFile(filePathChan <-chan string, resultChan chan<- DependencyFile) {
	for path := range filePathChan {
		depFile := ParseDependencyFile(path)
		resultChan <- depFile
	}
}
