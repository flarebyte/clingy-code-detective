package aggregator

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
)

// JSONAggregateRenderer implements AggregateRenderer for JSON output.
type JSONAggregateRenderer struct{}

func (r *JSONAggregateRenderer) Render(deps []AggregatedDependency) ([]byte, error) {
	return json.MarshalIndent(deps, "", "  ")
}

// CSVAggregateRenderer implements AggregateRenderer for CSV output.
type CSVAggregateRenderer struct{}

func (r *CSVAggregateRenderer) Render(deps []AggregatedDependency) ([]byte, error) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	header := []string{"Name", "MinVersion", "MaxVersion", "Count", "Category", "Packaging"}
	if err := writer.Write(header); err != nil {
		return nil, fmt.Errorf("error writing CSV header: %w", err)
	}

	for _, dep := range deps {
		record := []string{
			dep.Name,
			dep.MinVersion,
			dep.MaxVersion,
			fmt.Sprint(dep.Count),
			dep.Category,
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

// MarkdownAggregateRenderer implements AggregateRenderer for Markdown output.
type MarkdownAggregateRenderer struct{}

func (r *MarkdownAggregateRenderer) Render(deps []AggregatedDependency) ([]byte, error) {
	var buf bytes.Buffer

	buf.WriteString("| Name | MinVersion | MaxVersion | Count | Category | Packaging |\n")
	buf.WriteString("| ---- | ---------- | ---------- | ----- | -------- | --------- |\n")

	for _, dep := range deps {
		row := fmt.Sprintf("| %s | %s | %s | %d | %s | %s |\n",
			EscapeMarkdown(dep.Name),
			EscapeMarkdown(dep.MinVersion),
			EscapeMarkdown(dep.MaxVersion),
			dep.Count,
			EscapeMarkdown(dep.Category),
			EscapeMarkdown(dep.Packaging),
		)
		buf.WriteString(row)
	}

	return buf.Bytes(), nil
}
