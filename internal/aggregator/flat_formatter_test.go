package aggregator

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"strings"
	"testing"
)

func sampleDependencies() []FlatDependency {
	return []FlatDependency{
		{
			Name:      "dep1",
			Version:   "1.0.0",
			Category:  "prod",
			Path:      "/some/path",
			Packaging: "node",
		},
		{
			Name:      "dep2",
			Version:   "2.3.4",
			Category:  "dev",
			Path:      "/another/path",
			Packaging: "python",
		},
	}
}

func TestJSONRenderer(t *testing.T) {
	renderer := &JSONRenderer{}
	output, err := renderer.Render(sampleDependencies())
	if err != nil {
		t.Fatalf("JSONRenderer.Render returned error: %v", err)
	}

	var result []FlatDependency
	if err := json.Unmarshal(output, &result); err != nil {
		t.Fatalf("JSONRenderer.Render output is not valid JSON: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("expected 2 dependencies, got %d", len(result))
	}
}

func TestCSVRenderer(t *testing.T) {
	renderer := &CSVRenderer{}
	output, err := renderer.Render(sampleDependencies())
	if err != nil {
		t.Fatalf("CSVRenderer.Render returned error: %v", err)
	}

	reader := csv.NewReader(bytes.NewReader(output))
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("CSVRenderer.Render output is not valid CSV: %v", err)
	}

	if len(records) != 3 { // header + 2 rows
		t.Errorf("expected 3 CSV records, got %d", len(records))
	}

	expectedHeader := []string{"Name", "Version", "Category", "Path", "Packaging"}
	for i, field := range expectedHeader {
		if records[0][i] != field {
			t.Errorf("CSV header mismatch at column %d: got %q, want %q", i, records[0][i], field)
		}
	}
}

func TestMarkdownRenderer(t *testing.T) {
	renderer := &MarkdownRenderer{}
	output, err := renderer.Render(sampleDependencies())
	if err != nil {
		t.Fatalf("MarkdownRenderer.Render returned error: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) < 3 {
		t.Fatalf("expected at least 3 lines in Markdown output, got %d", len(lines))
	}

	if !strings.HasPrefix(lines[0], "| Name |") {
		t.Errorf("unexpected Markdown header: %q", lines[0])
	}

	if !strings.HasPrefix(lines[2], "| dep1 |") {
		t.Errorf("unexpected first row: %q", lines[2])
	}
}
