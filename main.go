package main

import (
	"fmt"
	"os"

	"github.com/flarebyte/clingy-code-detective/internal/cli"
)

func main() {
	cfg, err := cli.ParseArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	// Debug print to verify parsed config; replace with actual scanning logic.
	fmt.Printf("Paths: %v\n", cfg.Paths)
	fmt.Printf("Format: %s\n", cfg.Format)
	fmt.Printf("Aggregate: %v\n", cfg.Aggregate)
	fmt.Printf("Includes: %v\n", cfg.Includes)
}
