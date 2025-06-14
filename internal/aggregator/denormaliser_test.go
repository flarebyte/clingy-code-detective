package aggregator

import (
	"errors"
	"testing"

	"github.com/flarebyte/clingy-code-detective/internal/parser"
)

func TestDenormaliseDependencyFile(t *testing.T) {
	t.Run("returns flat dependencies when input is valid", func(t *testing.T) {
		file := parser.DependencyFile{
			Path: "deps.txt",
			Dependencies: []parser.Dependency{
				{Name: "foo", Version: "1.0.0", Category: "prod"},
				{Name: "bar", Version: "2.3.4", Category: "dev"},
			},
		}

		got := DenormaliseDependencyFile(file)

		want := []FlatDependency{
			{Name: "foo", Version: "1.0.0", Category: "prod", Path: "deps.txt"},
			{Name: "bar", Version: "2.3.4", Category: "dev", Path: "deps.txt"},
		}

		if len(got) != len(want) {
			t.Fatalf("expected %d dependencies, got %d", len(want), len(got))
		}

		for i, dep := range got {
			if dep != want[i] {
				t.Errorf("unexpected dependency at index %d: got %+v, want %+v", i, dep, want[i])
			}
		}
	})

	t.Run("returns empty slice when dependencies are empty", func(t *testing.T) {
		file := parser.DependencyFile{
			Path:         "empty.txt",
			Dependencies: []parser.Dependency{},
		}

		got := DenormaliseDependencyFile(file)

		if len(got) != 0 {
			t.Errorf("expected empty slice, got %v", got)
		}
	})

	t.Run("returns empty slice when error is present", func(t *testing.T) {
		file := parser.DependencyFile{
			Path: "broken.txt",
			Err:  errors.New("parse failed"),
			Dependencies: []parser.Dependency{
				{Name: "should-not-be-used", Version: "9.9.9", Category: "prod"},
			},
		}

		got := DenormaliseDependencyFile(file)

		if len(got) != 0 {
			t.Errorf("expected empty slice due to error, got %v", got)
		}
	})
}
