package scanner

import (
	"log"
	"os"
	"path/filepath"
)

var supportedFiles = map[string]struct{}{
	"package.json":     {},
	"pubspec.yaml":     {},
	"go.mod":           {},
	"requirements.txt": {},
}

// WalkDirectories walks the directory tree starting at root and sends
// the path of supported files into filePathChan. It closes filePathChan when done.
func WalkDirectories(root string, filePathChan chan<- string) {
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

		// Check if file is supported
		if _, ok := supportedFiles[d.Name()]; ok {
			filePathChan <- path
		}

		return nil
	})

	if err != nil {
		log.Printf("Walk error: %v\n", err)
	}
}
