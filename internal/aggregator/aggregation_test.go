package aggregator

import (
	"reflect"
	"testing"
)

func TestAggregateDependencies(t *testing.T) {
	tests := []struct {
		name  string
		input []FlatDependency
		want  []AggregatedDependency
	}{
		{
			name: "basic aggregation single group",
			input: []FlatDependency{
				{Name: "foo", Version: "1.0.0", Category: "prod", Packaging: "node"},
				{Name: "foo", Version: "1.2.0", Category: "prod", Packaging: "node"},
				{Name: "foo", Version: "1.1.0", Category: "prod", Packaging: "node"},
			},
			want: []AggregatedDependency{
				{
					Name: "foo", Category: "prod", Packaging: "node",
					Count: 3, MinVersion: "1.0.0", MaxVersion: "1.2.0",
				},
			},
		},
		{
			name: "multiple groups",
			input: []FlatDependency{
				{Name: "foo", Version: "1.0.0", Category: "prod", Packaging: "node"},
				{Name: "bar", Version: "2.0.0", Category: "dev", Packaging: "python"},
				{Name: "foo", Version: "1.2.0", Category: "prod", Packaging: "node"},
				{Name: "bar", Version: "2.1.0", Category: "dev", Packaging: "python"},
			},
			want: []AggregatedDependency{
				{
					Name: "bar", Category: "dev", Packaging: "python",
					Count: 2, MinVersion: "2.0.0", MaxVersion: "2.1.0",
				},
				{
					Name: "foo", Category: "prod", Packaging: "node",
					Count: 2, MinVersion: "1.0.0", MaxVersion: "1.2.0",
				},
			},
		},
		{
			name: "handles invalid version",
			input: []FlatDependency{
				{Name: "foo", Version: "1.0.0", Category: "prod", Packaging: "node"},
				{Name: "foo", Version: "bad.version", Category: "prod", Packaging: "node"},
				{Name: "foo", Version: "1.1.0", Category: "prod", Packaging: "node"},
			},
			want: []AggregatedDependency{
				{
					Name: "foo", Category: "prod", Packaging: "node",
					Count: 2, MinVersion: "1.0.0", MaxVersion: "1.1.0",
				},
			},
		},
		{
			name:  "empty input",
			input: []FlatDependency{},
			want:  []AggregatedDependency{},
		},
		{
			name: "grouping respects category and packaging",
			input: []FlatDependency{
				{Name: "foo", Version: "1.0.0", Category: "prod", Packaging: "node"},
				{Name: "foo", Version: "1.1.0", Category: "dev", Packaging: "node"},
				{Name: "foo", Version: "1.2.0", Category: "prod", Packaging: "python"},
			},
			want: []AggregatedDependency{
				{
					Name: "foo", Category: "dev", Packaging: "node",
					Count: 1, MinVersion: "1.1.0", MaxVersion: "1.1.0",
				},
				{
					Name: "foo", Category: "prod", Packaging: "node",
					Count: 1, MinVersion: "1.0.0", MaxVersion: "1.0.0",
				},
				{
					Name: "foo", Category: "prod", Packaging: "python",
					Count: 1, MinVersion: "1.2.0", MaxVersion: "1.2.0",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AggregateDependencies(tt.input)

			// Compare ignoring order
			if !equalAggregatedDependencies(got, tt.want) {
				t.Errorf("AggregateDependencies() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

// Helper to compare slices ignoring order
func equalAggregatedDependencies(a, b []AggregatedDependency) bool {
	if len(a) != len(b) {
		return false
	}

	aMap := make(map[string]AggregatedDependency)
	bMap := make(map[string]AggregatedDependency)

	for _, x := range a {
		key := x.Name + "|" + x.Category + "|" + x.Packaging
		aMap[key] = x
	}
	for _, x := range b {
		key := x.Name + "|" + x.Category + "|" + x.Packaging
		bMap[key] = x
	}

	return reflect.DeepEqual(aMap, bMap)
}
