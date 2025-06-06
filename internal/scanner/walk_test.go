package scanner

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

// helper to create a file
func createFile(t *testing.T, dir, name string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	err := os.WriteFile(path, []byte("test content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file %s: %v", path, err)
	}
	return path
}

func TestWalkDirectories(t *testing.T) {
	tmpDir := t.TempDir()

	// Create supported files
	supported := []string{
		"package.json",
		"pubspec.yaml",
		"go.mod",
		"requirements.txt",
	}

	var expectedPaths []string
	for _, fname := range supported {
		path := createFile(t, tmpDir, fname)
		expectedPaths = append(expectedPaths, path)
	}

	// Create some unsupported files
	_ = createFile(t, tmpDir, "README.md")
	_ = createFile(t, tmpDir, "main.go")

	// Run WalkDirectories
	filePathChan := make(chan string)
	var foundPaths []string

	go WalkDirectories(tmpDir, filePathChan)

	for path := range filePathChan {
		foundPaths = append(foundPaths, path)
	}

	// Sort to make order irrelevant
	slices.Sort(expectedPaths)
	slices.Sort(foundPaths)

	if len(expectedPaths) != len(foundPaths) {
		t.Fatalf("Expected %d supported files, found %d", len(expectedPaths), len(foundPaths))
	}

	for i := range expectedPaths {
		if expectedPaths[i] != foundPaths[i] {
			t.Errorf("Expected %q, got %q", expectedPaths[i], foundPaths[i])
		}
	}
}
