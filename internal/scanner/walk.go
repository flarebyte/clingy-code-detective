package scanner

import (
	"log"
	"os"
	"path/filepath"
)

// WalkDirectories walks the directory tree starting at root and sends
// the path of required files into filePathChan. It closes filePathChan when done.
func WalkDirectories(root string, includes []string, excludes []string, filePathChan chan<- string) {
	defer close(filePathChan)

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			// Log error and continue walking
			log.Printf("Error accessing path %q: %v\n", path, err)
			return nil
		}

		// Check excludes first — skip entire subtree if directory is excluded
		if IsFileExcluded(path, excludes) {
			if d.IsDir() {
				// Skip entire directory
				return filepath.SkipDir
			}
			// Skip this file
			return nil
		}

		// Skip directories (already handled above)
		if d.IsDir() {
			return nil
		}

		// Use IsFileRequired to check if this file should be included
		if IsFileRequired(d.Name(), includes) {
			filePathChan <- path
		}

		return nil
	})

	if err != nil {
		log.Printf("Walk error: %v\n", err)
	}
}
