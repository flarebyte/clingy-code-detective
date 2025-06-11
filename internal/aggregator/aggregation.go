package aggregator

import (
	"sort"

	"github.com/Masterminds/semver/v3"
)

// AggregateDependencies aggregates a slice of FlatDependency into AggregatedDependency.
func AggregateDependencies(deps []FlatDependency) []AggregatedDependency {
	type aggState struct {
		minV  *semver.Version
		maxV  *semver.Version
		count uint
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

		v, err := semver.NewVersion(d.Version)
		if err != nil {
			// Invalid version — skip it
			continue
		}

		agg := state[key]

		if agg.minV == nil || v.LessThan(agg.minV) {
			agg.minV = v
		}
		if agg.maxV == nil || v.GreaterThan(agg.maxV) {
			agg.maxV = v
		}

		agg.count++
	}

	// Build final result
	var result []AggregatedDependency
	for key, agg := range state {
		m := meta[key]

		ad := AggregatedDependency{
			Name:      m.Name,
			Category:  m.Category,
			Packaging: m.Packaging,
			Count:     agg.count,
		}

		if agg.minV != nil {
			ad.MinVersion = agg.minV.String()
		}
		if agg.maxV != nil {
			ad.MaxVersion = agg.maxV.String()
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
