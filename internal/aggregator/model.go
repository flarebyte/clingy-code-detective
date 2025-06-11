package aggregator

// A single dependency
type FlatDependency struct {
	Name      string
	Version   string
	Category  string // e.g., "prod", "dev"
	Path      string
	Packaging string // e.g., "node", "python"
}

// An aggregated dependency representeing all the dependency with the same name
type AggregatedDependency struct {
	Name       string
	MinVersion string
	MaxVersion string
	Count      uint
	Category   string // e.g., "prod", "dev"
	Packaging  string // e.g., "node", "python"
}
