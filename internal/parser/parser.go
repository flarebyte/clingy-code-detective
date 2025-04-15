import (
	"errors"
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

	deps, err := parser.Parse(path)
	return DependencyFile{
		Path:         path,
		Dependencies: deps,
		Err:          err,
	}
}
