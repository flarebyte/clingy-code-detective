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
	supported := map[string]string{
		"package.json":     createFile(t, tmpDir, "package.json"),
		"pubspec.yaml":     createFile(t, tmpDir, "pubspec.yaml"),
		"go.mod":           createFile(t, tmpDir, "go.mod"),
		"requirements.txt": createFile(t, tmpDir, "requirements.txt"),
	}

	// Create unsupported files
	_ = createFile(t, tmpDir, "README.md")
	_ = createFile(t, tmpDir, "main.go")

	tests := []struct {
		name          string
		includes      []string
		expectedPaths []string
	}{
		{
			name:          "empty includes - all supported files",
			includes:      []string{},
			expectedPaths: slices.Clone(mapsValues(supported)),
		},
		{
			name:          "includes python only",
			includes:      []string{"python"},
			expectedPaths: []string{supported["requirements.txt"]},
		},
		{
			name:          "includes node only",
			includes:      []string{"node"},
			expectedPaths: []string{supported["package.json"]},
		},
		{
			name:          "includes js alias",
			includes:      []string{"js"},
			expectedPaths: []string{supported["package.json"]},
		},
		{
			name:          "includes dart and go",
			includes:      []string{"dart", "go"},
			expectedPaths: []string{supported["pubspec.yaml"], supported["go.mod"]},
		},
		{
			name:          "includes python python js duplicates",
			includes:      []string{"python", "python", "js"},
			expectedPaths: []string{supported["requirements.txt"], supported["package.json"]},
		},
		{
			name:          "includes unknown category",
			includes:      []string{"unknown"},
			expectedPaths: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePathChan := make(chan string)
			var foundPaths []string

			go WalkDirectories(tmpDir, tt.includes, filePathChan)

			for path := range filePathChan {
				foundPaths = append(foundPaths, path)
			}

			// Sort to make order irrelevant
			slices.Sort(foundPaths)
			slices.Sort(tt.expectedPaths)

			if len(foundPaths) != len(tt.expectedPaths) {
				t.Fatalf("Expected %d files, found %d.\nExpected: %v\nFound: %v",
					len(tt.expectedPaths), len(foundPaths), tt.expectedPaths, foundPaths)
			}

			for i := range tt.expectedPaths {
				if foundPaths[i] != tt.expectedPaths[i] {
					t.Errorf("Expected %q, got %q", tt.expectedPaths[i], foundPaths[i])
				}
			}
		})
	}
}

// mapsValues extracts the values from a map[string]string.
func mapsValues(m map[string]string) []string {
	values := make([]string, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
