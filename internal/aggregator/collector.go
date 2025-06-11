package aggregator

import (
	"fmt"
	"os"
	"sort"

	"github.com/flarebyte/clingy-code-detective/internal/parser"
)

// SortFlatDependencies sorts the given slice by Packaging, then Category, then Name.
func sortFlatDependencies(deps []FlatDependency) {
	sort.Slice(deps, func(i, j int) bool {
		if deps[i].Packaging != deps[j].Packaging {
			return deps[i].Packaging < deps[j].Packaging
		}
		if deps[i].Category != deps[j].Category {
			return deps[i].Category < deps[j].Category
		}
		return deps[i].Name < deps[j].Name
	})
}

// collectResults reads DependencyFile results and processes them.
// Once resultChan is closed and drained, it signals completion on done chan.
func CollectDependencies(resultChan <-chan parser.DependencyFile, done chan<- []FlatDependency) {
	var flatDependencies []FlatDependency
	for depFile := range resultChan {
		if depFile.Err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing %s: %v\n", depFile.Path, depFile.Err)
			continue
		} else {
			flatDependencies = append(flatDependencies, DenormaliseDependencyFile(depFile)...)
		}
	}
	sortFlatDependencies(flatDependencies)
	done <- flatDependencies
}
