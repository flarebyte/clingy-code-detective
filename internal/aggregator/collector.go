package aggregator

import (
	"fmt"
	"os"

	"github.com/flarebyte/clingy-code-detective/internal/parser"
)

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
	done <- flatDependencies
}
