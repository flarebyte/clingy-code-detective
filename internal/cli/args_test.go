package cli

import (
	"flag"
	"os"
	"reflect"
	"strings"
	"testing"
)

// helper to temporarily override os.Args and reset flag.CommandLine
func withArgs(args []string, fn func()) {
	origArgs := os.Args
	defer func() { os.Args = origArgs }()
	os.Args = append([]string{origArgs[0]}, args...)

	// Reset flags
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	fn()
}

func TestParseArgs_Basic(t *testing.T) {
	args := []string{
		"--json",
		"--include=node,dart",
		"--exclude=/node-modules/,/dist/",
		"project1",
		"project2",
	}

	cfg, err := ParseArgsFrom(args)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Format != "json" {
		t.Errorf("Format = %q, want \"json\"", cfg.Format)
	}

	if cfg.Aggregate {
		t.Errorf("Aggregate = true, want false by default")
	}

	wantIncludes := []string{"node", "dart"}
	if !reflect.DeepEqual(cfg.Includes, wantIncludes) {
		t.Errorf("Includes = %v, want %v", cfg.Includes, wantIncludes)
	}

	wantExcludes := []string{"/node-modules/", "/dist/"}
	if !reflect.DeepEqual(cfg.Excludes, wantExcludes) {
		t.Errorf("Excludes = %v, want %v", cfg.Excludes, wantExcludes)
	}

	if len(cfg.Paths) != 2 || cfg.Paths[0] != "project1" || cfg.Paths[1] != "project2" {
		t.Errorf("Paths = %v, want [project1 project2]", cfg.Paths)
	}
}

func TestParseArgs_MutuallyExclusiveFormats(t *testing.T) {
	args := []string{
		"--json",
		"--csv",
		"dir",
	}

	_, err := ParseArgsFrom(args)
	if err == nil || !strings.Contains(err.Error(), "only one of --json, --csv, --md") {
		t.Errorf("expected mutually exclusive error, got: %v", err)
	}
}

func TestParseArgs_NoPaths(t *testing.T) {
	withArgs([]string{"--json"}, func() {
		_, err := ParseArgs()
		if err == nil || !strings.Contains(err.Error(), "at least one directory path") {
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

func TestParseArgs_HelpFlag(t *testing.T) {
	cfg, err := ParseArgsFrom([]string{"--help"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.ShowHelp {
		t.Error("expected ShowHelp to be true")
	}
}

func TestParseArgs_VersionFlag(t *testing.T) {
	cfg, err := ParseArgsFrom([]string{"--version"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.ShowVer {
		t.Error("expected ShowVer to be true")
	}
}
