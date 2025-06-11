package aggregator

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"
)

// JSONRenderer implements Renderer for JSON output.
type JSONRenderer struct{}

// Render renders dependencies as JSON.
func (r *JSONRenderer) Render(deps []FlatDependency) ([]byte, error) {
	return json.MarshalIndent(deps, "", "  ")
}

// CSVRenderer implements Renderer for CSV output.
type CSVRenderer struct{}

// Render renders dependencies as CSV.
func (r *CSVRenderer) Render(deps []FlatDependency) ([]byte, error) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// Write header
	header := []string{"Name", "Version", "Category", "Path", "Packaging"}
	if err := writer.Write(header); err != nil {
		return nil, fmt.Errorf("error writing CSV header: %w", err)
	}

	// Write records
	for _, dep := range deps {
		record := []string{
			dep.Name,
			dep.Version,
			dep.Category,
			dep.Path,
			dep.Packaging,
		}
		if err := writer.Write(record); err != nil {
			return nil, fmt.Errorf("error writing CSV record: %w", err)
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return nil, fmt.Errorf("error flushing CSV writer: %w", err)
	}

	return buf.Bytes(), nil
}

// MarkdownRenderer implements FlatRenderer for Markdown table output.
type MarkdownRenderer struct{}

// Render renders dependencies as a Markdown table.
func (r *MarkdownRenderer) Render(deps []FlatDependency) ([]byte, error) {
	var buf bytes.Buffer

	// Write header
	buf.WriteString("| Name | Version | Category | Path | Packaging |\n")
	buf.WriteString("| ---- | ------- | -------- | ---- | --------- |\n")

	// Write rows
	for _, dep := range deps {
		row := fmt.Sprintf("| %s | %s | %s | %s | %s |\n",
			escapeMarkdown(dep.Name),
			escapeMarkdown(dep.Version),
			escapeMarkdown(dep.Category),
			escapeMarkdown(dep.Path),
			escapeMarkdown(dep.Packaging),
		)
		buf.WriteString(row)
	}

	return buf.Bytes(), nil
}

func escapeMarkdown(s string) string {
	return strings.ReplaceAll(s, "|", "\\|")
}
