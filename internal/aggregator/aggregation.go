package aggregator

import (
	"sort"
)

// AggregateDependencies aggregates a slice of FlatDependency into AggregatedDependency.
func AggregateDependencies(deps []FlatDependency) []AggregatedDependency {
	type aggState struct {
		minVersion string
		maxVersion string
		count      uint
	}

	// Grouping key: name + category + packaging
	keyFor := func(d FlatDependency) string {
		return d.Name + "|" + d.Category + "|" + d.Packaging
	}

	state := make(map[string]*aggState)
	meta := make(map[string]struct {
		Name      string
		Category  string
		Packaging string
	})

	for _, d := range deps {
		key := keyFor(d)

		if _, ok := state[key]; !ok {
			state[key] = &aggState{}
			meta[key] = struct {
				Name      string
				Category  string
				Packaging string
			}{
				Name:      d.Name,
				Category:  d.Category,
				Packaging: d.Packaging,
			}
		}

		agg := state[key]

		if agg.minVersion == "" {
			agg.minVersion = d.Version
		} else {
			minV, err := MinVersion(agg.minVersion, d.Version)
			if err == nil {
				agg.minVersion = minV.String()
			}
		}

		if agg.maxVersion == "" {
			agg.maxVersion = d.Version
		} else {
			maxV, err := MaxVersion(agg.maxVersion, d.Version)
			if err == nil {
				agg.maxVersion = maxV.String()
			}
		}

		agg.count++
	}

	// Build final result
	var result []AggregatedDependency
	for key, agg := range state {
		m := meta[key]

		ad := AggregatedDependency{
			Name:       m.Name,
			Category:   m.Category,
			Packaging:  m.Packaging,
			Count:      agg.count,
			MinVersion: agg.minVersion,
			MaxVersion: agg.maxVersion,
		}

		result = append(result, ad)
	}

	// Optional: sort result by name for stability
	sort.Slice(result, func(i, j int) bool {
		if result[i].Name != result[j].Name {
			return result[i].Name < result[j].Name
		}
		if result[i].Category != result[j].Category {
			return result[i].Category < result[j].Category
		}
		return result[i].Packaging < result[j].Packaging
	})

	return result
}
