package cli

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

// helper to temporarily override os.Args and reset flag.CommandLine
func withArgs(args []string, fn func()) {
	origArgs := os.Args
	defer func() { os.Args = origArgs }()
	os.Args = append([]string{origArgs[0]}, args...)

	// reset flags between tests
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	fn()
}

func TestParseArgs_Basic(t *testing.T) {
	withArgs([]string{"project1", "project2", "--json", "--include=node,dart"}, func() {
		cfg, err := ParseArgs()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got, want := cfg.Format, "json"; got != want {
			t.Errorf("Format = %q, want %q", got, want)
		}
		if !cfg.Aggregate {
			t.Errorf("expected Aggregate to be false by default")
		}
		if !reflect.DeepEqual(cfg.Includes, []string{"node", "dart"}) {
			t.Errorf("Includes = %v, want [node dart]", cfg.Includes)
		}
		if len(cfg.Paths) != 2 {
			t.Errorf("Paths = %v, want 2 paths", cfg.Paths)
		}
	})
}

func TestParseArgs_MutuallyExclusiveFormats(t *testing.T) {
	withArgs([]string{"dir", "--json", "--csv"}, func() {
		_, err := ParseArgs()
		if err == nil || err.Error() != "only one of --json or --csv can be used" {
			t.Errorf("expected mutually exclusive error, got: %v", err)
		}
	})
}

func TestParseArgs_NoPaths(t *testing.T) {
	withArgs([]string{"--json"}, func() {
		_, err := ParseArgs()
		if err == nil || err.Error() != "at least one directory path must be specified" {
			t.Errorf("expected missing path error, got: %v", err)
		}
	})
}

func TestParseArgs_Defaults(t *testing.T) {
	withArgs([]string{"some-dir"}, func() {
		cfg, err := ParseArgs()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if cfg.Format != "" {
			t.Errorf("expected Format to be empty, got %q", cfg.Format)
		}
		if cfg.Aggregate {
			t.Errorf("expected Aggregate to be false by default")
		}
		if len(cfg.Includes) != 0 {
			t.Errorf("expected no Includes, got %v", cfg.Includes)
		}
	})
}
