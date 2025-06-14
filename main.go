package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/flarebyte/clingy-code-detective/internal/aggregator"
	"github.com/flarebyte/clingy-code-detective/internal/cli"
	"github.com/flarebyte/clingy-code-detective/internal/parser"
	"github.com/flarebyte/clingy-code-detective/internal/scanner"
)

func main() {
	cfg, err := cli.ParseArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	numWorkers := runtime.NumCPU()

	filePathChan := make(chan string)
	resultChan := make(chan parser.DependencyFile)

	var wg sync.WaitGroup

	var root = cfg.Paths[0]

	go scanner.WalkDirectories(root, cfg.Includes, cfg.Excludes, filePathChan)

	//Parse each file with a pool of workers
	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			parser.ProduceDependencyFile(filePathChan, resultChan)
		}()
	}
	done := make(chan []aggregator.FlatDependency, 1)
	go aggregator.CollectDependencies(resultChan, done)

	wg.Wait()
	close(resultChan)

	var flatRenderer aggregator.FlatRenderer
	var aggregateRenderer aggregator.AggregateRenderer

	switch cfg.Format {
	case "json":
		flatRenderer = &aggregator.JSONRenderer{}
		aggregateRenderer = &aggregator.JSONAggregateRenderer{}
	case "csv":
		flatRenderer = &aggregator.CSVRenderer{}
		aggregateRenderer = &aggregator.CSVAggregateRenderer{}
	case "md":
		flatRenderer = &aggregator.MarkdownRenderer{}
		aggregateRenderer = &aggregator.MarkdownAggregateRenderer{}
	default:
		log.Fatalf("unknown format: %s", cfg.Format)
	}

	// Render output
	flatDependencies := <-done

	if cfg.Aggregate {
		aggegateDependencies := aggregator.AggregateDependencies(flatDependencies)

		output, err := aggregateRenderer.Render(aggegateDependencies)
		if err != nil {
			log.Fatalf("failed to render dependencies after aggregation: %v", err)
		}

		fmt.Println(string(output))

	} else {
		output, err := flatRenderer.Render(flatDependencies)
		if err != nil {
			log.Fatalf("failed to render dependencies: %v", err)
		}

		fmt.Println(string(output))

	}

}
