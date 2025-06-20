package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Version should be set at build time using -ldflags "-X 'cli.Version=1.2.3'"
var (
	Version string
	Date    string
)
var niceDate = strings.ReplaceAll(Date, "-", " ")

const helpIntro = `clingy - expose the code that's a little too attached
Also known as clingy-code-detective, a command-line tool to scan project directories for dependencies across multiple ecosystems, aggregating and reporting them.	
Copyright (c) 2025 Flarebyte.com - MIT License
`

const helpUsage = `
Usage: clingy [options] <paths>

Options:
  --include    Comma-separated ecosystems to include (e.g. node,dart)
  --exclude    Comma-separated path segments to exclude (e.g. /node_modules/)
  --json       Output in JSON format
  --csv        Output in CSV format
  --md         Output in Markdown format
  --aggregate  Aggregate results across all directories
  --version    Show version information
  --help       Show this help message
`

// Config holds the parsed CLI arguments.
type Config struct {
	Paths     []string
	Format    string
	Aggregate bool
	Includes  []string
	Excludes  []string
	ShowHelp  bool
	ShowVer   bool
}

type parseIncludes []string
type parseExcludes []string

func (i *parseIncludes) String() string         { return strings.Join(*i, ",") }
func (i *parseIncludes) Set(value string) error { *i = strings.Split(value, ","); return nil }
func (i *parseExcludes) String() string         { return strings.Join(*i, ",") }
func (i *parseExcludes) Set(value string) error { *i = strings.Split(value, ","); return nil }

// ParseArgsFrom parses CLI arguments into a Config.
func ParseArgsFrom(args []string) (*Config, error) {
	var includes parseIncludes
	var excludes parseExcludes
	var jsonOut, csvOut, mdOut, aggregate, showHelp, showVer bool

	fs := flag.NewFlagSet("clingy", flag.ContinueOnError)
	fs.Usage = func() {
		fmt.Printf("%sVersion: %s, Date: %s \n%s", helpIntro, Version, niceDate, helpUsage)
	}

	fs.Var(&includes, "include", "Ecosystems to include")
	fs.Var(&excludes, "exclude", "Path segments to exclude")
	fs.BoolVar(&jsonOut, "json", false, "Output JSON")
	fs.BoolVar(&csvOut, "csv", false, "Output CSV")
	fs.BoolVar(&mdOut, "md", false, "Output Markdown")
	fs.BoolVar(&aggregate, "aggregate", false, "Aggregate results")
	fs.BoolVar(&showVer, "version", false, "Show version")
	fs.BoolVar(&showHelp, "help", false, "Show help")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if showHelp {
		fs.Usage()
		return &Config{ShowHelp: true}, nil
	}

	if showVer {
		fmt.Println("clingy version", Version)
		return &Config{ShowVer: true}, nil
	}

	paths := fs.Args()
	if len(paths) == 0 {
		fs.Usage()
		return &Config{ShowHelp: true}, nil
	}

	var format string
	switch {
	case (jsonOut && csvOut) || (jsonOut && mdOut) || (csvOut && mdOut):
		return nil, fmt.Errorf("only one of --json, --csv, --md may be used")
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

// ParseArgs uses os.Args and defaults to help on empty args.
func ParseArgs() (*Config, error) {
	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"--help"}
	}
	return ParseArgsFrom(args)
}
