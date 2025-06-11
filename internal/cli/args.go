package cli

import (
	"errors"
	"flag"
	"os"
	"strings"
)

// Config holds the parsed CLI arguments.
type Config struct {
	Paths     []string // Directories to scan
	Format    string   // Output format: json, csv
	Aggregate bool     // Whether to aggregate results
	Includes  []string // Filter by ecosystems (e.g., node, dart)
	Excludes  []string // Discard path segments (e.g. /node_modules/)
}

// parseIncludes and parseExcludes are custom flag.Value for comma-separated lists.
type parseIncludes []string
type parseExcludes []string

func (i *parseIncludes) String() string {
	return strings.Join(*i, ",")
}

func (i *parseIncludes) Set(value string) error {
	*i = strings.Split(value, ",")
	return nil
}

func (i *parseExcludes) String() string {
	return strings.Join(*i, ",")
}

func (i *parseExcludes) Set(value string) error {
	*i = strings.Split(value, ",")
	return nil
}

// ParseArgsFrom parses command-line arguments into Config.
func ParseArgsFrom(args []string) (*Config, error) {
	var includes parseIncludes
	var excludes parseExcludes
	var jsonOut, csvOut, mdOut, aggregate bool

	fs := flag.NewFlagSet("clingy", flag.ContinueOnError)
	fs.Var(&includes, "include", "Comma-separated list of ecosystems to include (e.g., node,dart)")
	fs.Var(&excludes, "exclude", "Comma-separated list of path segments to exclude (e.g., /node_modules/,'/dist/')")
	fs.BoolVar(&jsonOut, "json", false, "Output in JSON format")
	fs.BoolVar(&csvOut, "csv", false, "Output in CSV format")
	fs.BoolVar(&mdOut, "md", false, "Output in markdown format")
	fs.BoolVar(&aggregate, "aggregate", false, "Aggregate results across all directories")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	paths := fs.Args()
	if len(paths) == 0 {
		return nil, errors.New("at least one directory path must be specified")
	}

	var format string
	switch {
	case (jsonOut && csvOut) || (jsonOut && mdOut) || (csvOut && mdOut):
		return nil, errors.New("only one of --json or --csv or --md can be used")
	case jsonOut:
		format = "json"
	case csvOut:
		format = "csv"
	case mdOut:
		format = "md"
	}

	return &Config{
		Paths:     paths,
		Format:    format,
		Aggregate: aggregate,
		Includes:  includes,
		Excludes:  excludes,
	}, nil
}

// ParseArgs parses command-line arguments into Config.
func ParseArgs() (*Config, error) {
	return ParseArgsFrom(os.Args[1:])
}
