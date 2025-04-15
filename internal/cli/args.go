package cli

import (
	"errors"
	"flag"
	"fmt"
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

// ParseArgs parses command-line arguments into Config.
func ParseArgs() (*Config, error) {
	var includes parseIncludes
	var jsonOut, csvOut, aggregate bool

	flag.Var(&includes, "include", "Comma-separated list of ecosystems to include (e.g., node,dart)")
	flag.BoolVar(&jsonOut, "json", false, "Output in JSON format")
	flag.BoolVar(&csvOut, "csv", false, "Output in CSV format")
	flag.BoolVar(&aggregate, "aggregate", false, "Aggregate results across all directories")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] <paths...>\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	// Require at least one path
	paths := flag.Args()
	if len(paths) == 0 {
		return nil, errors.New("at least one directory path must be specified")
	}

	// Enforce mutual exclusivity between JSON and CSV formats
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
