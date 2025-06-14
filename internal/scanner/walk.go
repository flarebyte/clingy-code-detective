package scanner

import (
	"log"
	"os"
	"path/filepath"
	"sync"
)

// WalkDirectories walks the directory trees starting at each root and sends
// the path of required files into filePathChan. It closes filePathChan when done.
func WalkDirectories(roots []string, includes []string, excludes []string, filePathChan chan<- string) {
	var wg sync.WaitGroup

	for _, root := range roots {
		wg.Add(1)
		go func(root string) {
			defer wg.Done()
			err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
				if err != nil {
					log.Printf("Error accessing path %q: %v\n", path, err)
					return nil
				}
				if IsFileExcluded(path, excludes) {
					if d.IsDir() {
						return filepath.SkipDir
					}
					return nil
				}
				if d.IsDir() {
					return nil
				}
				if IsFileRequired(d.Name(), includes) {
					filePathChan <- path
				}
				return nil
			})
			if err != nil {
				log.Printf("Walk error: %v\n", err)
			}
		}(root)
	}

	// Close the channel after all goroutines finish
	go func() {
		wg.Wait()
		close(filePathChan)
	}()
}
