package parser

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNodeParser_Parse(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected []Dependency
	}{
		{
			name: "Only prod dependencies",
			input: map[string]interface{}{
				"dependencies": map[string]string{
					"express": "4.17.1",
					"lodash":  "4.17.21",
				},
			},
			expected: []Dependency{
				{Name: "express", Version: "4.17.1", Category: "prod"},
				{Name: "lodash", Version: "4.17.21", Category: "prod"},
			},
		},
		{
			name: "Only dev dependencies",
			input: map[string]interface{}{
				"devDependencies": map[string]string{
					"mocha": "9.0.0",
				},
			},
			expected: []Dependency{
				{Name: "mocha", Version: "9.0.0", Category: "dev"},
			},
		},
		{
			name: "Both prod and dev dependencies",
			input: map[string]interface{}{
				"dependencies": map[string]string{
					"react": "17.0.2",
				},
				"devDependencies": map[string]string{
					"eslint": "7.32.0",
				},
			},
			expected: []Dependency{
				{Name: "react", Version: "17.0.2", Category: "prod"},
				{Name: "eslint", Version: "7.32.0", Category: "dev"},
			},
		},
		{
			name:     "Empty package.json",
			input:    map[string]interface{}{},
			expected: nil,
		},
	}

	parser := nodeParser{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("failed to marshal input: %v", err)
			}

			got, err := parser.Parse(content)
			if err != nil {
				t.Fatalf("Parse() error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Parse() got = %v, want = %v", got, tt.expected)
			}
		})
	}
}
