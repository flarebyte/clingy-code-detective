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
}

// parseIncludes is a custom flag.Value for comma-separated lists.
type parseIncludes []string

func (i *parseIncludes) String() string {
	return strings.Join(*i, ",")
}

func (i *parseIncludes) Set(value string) error {
	*i = strings.Split(value, ",")
	return nil
}

// ParseArgsFrom parses command-line arguments into Config.
func ParseArgsFrom(args []string) (*Config, error) {
	var includes parseIncludes
	var jsonOut, csvOut, aggregate bool

	fs := flag.NewFlagSet("clingy", flag.ContinueOnError)
	fs.Var(&includes, "include", "Comma-separated list of ecosystems to include (e.g., node,dart)")
	fs.BoolVar(&jsonOut, "json", false, "Output in JSON format")
	fs.BoolVar(&csvOut, "csv", false, "Output in CSV format")
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
	case jsonOut && csvOut:
		return nil, errors.New("only one of --json or --csv can be used")
	case jsonOut:
		format = "json"
	case csvOut:
		format = "csv"
	}

	return &Config{
		Paths:     paths,
		Format:    format,
		Aggregate: aggregate,
		Includes:  includes,
	}, nil
}

// ParseArgs parses command-line arguments into Config.
func ParseArgs() (*Config, error) {
	return ParseArgsFrom(os.Args[1:])
}
