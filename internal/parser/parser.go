package parser

import (
	"errors"
	"os"
	"path/filepath"
)

// ParseDependencyFile opens and parses a file using the appropriate parser.
func ParseDependencyFile(path string) DependencyFile {
	var parser Parser
	var packaging string

	switch filepath.Base(path) {
	case "package.json":
		parser = nodeParser{}
		packaging = "node"
	case "pubspec.yaml":
		parser = dartParser{}
		packaging = "dart"
	case "go.mod":
		parser = goModParser{}
		packaging = "go"
	case "requirements.txt":
		parser = pythonParser{}
		packaging = "python"
	default:
		return DependencyFile{Path: path, Packaging: "", Err: errors.New("unsupported file type")}
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
		Packaging:    packaging,
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
