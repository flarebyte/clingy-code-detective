package aggregator

import "encoding/json"

// JSONRenderer implements Renderer for JSON output.
type JSONRenderer struct{}

// Render renders dependencies as JSON.
func (r *JSONRenderer) Render(deps []FlatDependency) ([]byte, error) {
	return json.MarshalIndent(deps, "", "  ")
}
