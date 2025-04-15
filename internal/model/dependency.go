// Dependency represents a single declared dependency.
type Dependency struct {
	Name     string
	Version  string
	Category string // e.g., "prod", "dev"
}

// DependencyFile holds the metadata and results of parsing a dependency file.
type DependencyFile struct {
	Path         string
	Dependencies []Dependency
	Err          error
}

// Parser is implemented by each language-specific dependency file parser.
type Parser interface {
	Parse(path string) ([]Dependency, error)
}
