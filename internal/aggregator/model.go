package aggregator

// A single dependency
type FlatDependency struct {
	Name      string
	Version   string
	Category  string // e.g., "prod", "dev"
	Path      string
	Packaging string // e.g., "node", "python"
}
