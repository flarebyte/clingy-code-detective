package scanner

import (
	"log"
	"os"
	"path/filepath"
)

// WalkDirectories walks the directory tree starting at root and sends
// the path of required files into filePathChan. It closes filePathChan when done.
func WalkDirectories(root string, includes []string, filePathChan chan<- string) {
	defer close(filePathChan)

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			// Log error and continue walking
			log.Printf("Error accessing path %q: %v\n", path, err)
			return nil
		}

		// Skip directories
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
