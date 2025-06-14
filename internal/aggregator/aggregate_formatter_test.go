package aggregator

import (
	"encoding/json"
	"strings"
	"testing"
)

var sampleDeps = []AggregatedDependency{
	{
		Name:       "dep1",
		MinVersion: "1.0.0",
		MaxVersion: "2.0.0",
		Count:      3,
		Category:   "prod",
		Packaging:  "node",
	},
	{
		Name:       "dep2",
		MinVersion: "0.1.0",
		MaxVersion: "1.2.3",
		Count:      1,
		Category:   "dev",
		Packaging:  "python",
	},
}

func TestJSONAggregateRenderer_Render(t *testing.T) {
	r := &JSONAggregateRenderer{}
	out, err := r.Render(sampleDeps)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var unmarshalled []AggregatedDependency
	if err := json.Unmarshal(out, &unmarshalled); err != nil {
		t.Fatalf("failed to unmarshal output: %v", err)
	}

	if len(unmarshalled) != len(sampleDeps) {
		t.Errorf("expected %d dependencies, got %d", len(sampleDeps), len(unmarshalled))
	}
}

func TestCSVAggregateRenderer_Render(t *testing.T) {
	r := &CSVAggregateRenderer{}
	out, err := r.Render(sampleDeps)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) != len(sampleDeps)+1 {
		t.Errorf("expected %d lines, got %d", len(sampleDeps)+1, len(lines))
	}

	header := "Name,MinVersion,MaxVersion,Count,Category,Packaging"
	if lines[0] != header {
		t.Errorf("expected header %q, got %q", header, lines[0])
	}
}

func TestMarkdownAggregateRenderer_Render(t *testing.T) {
	r := &MarkdownAggregateRenderer{}
	out, err := r.Render(sampleDeps)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	expectedLineCount := len(sampleDeps) + 2 // header + separator + rows
	if len(lines) != expectedLineCount {
		t.Errorf("expected %d lines, got %d", expectedLineCount, len(lines))
	}

	if !strings.HasPrefix(lines[0], "| Name |") {
		t.Errorf("unexpected markdown header: %s", lines[0])
	}
}
