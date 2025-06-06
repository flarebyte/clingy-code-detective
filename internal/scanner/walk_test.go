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

// helper to create a subdirectory
func createDir(t *testing.T, parent, name string) string {
	t.Helper()
	path := filepath.Join(parent, name)
	err := os.Mkdir(path, 0755)
	if err != nil {
		t.Fatalf("Failed to create directory %s: %v", path, err)
	}
	return path
}

func TestWalkDirectories(t *testing.T) {
	tmpDir := t.TempDir()

	// Create supported files in root
	supported := map[string]string{
		"package.json":     createFile(t, tmpDir, "package.json"),
		"pubspec.yaml":     createFile(t, tmpDir, "pubspec.yaml"),
		"go.mod":           createFile(t, tmpDir, "go.mod"),
		"requirements.txt": createFile(t, tmpDir, "requirements.txt"),
	}

	// Create unsupported files
	_ = createFile(t, tmpDir, "README.md")
	_ = createFile(t, tmpDir, "main.go")

	// Create subdir with supported files
	nodeModulesDir := createDir(t, tmpDir, "node_modules")
	nodeModulesFile := createFile(t, nodeModulesDir, "package.json")

	vendorDir := createDir(t, tmpDir, "vendor")
	vendorFile := createFile(t, vendorDir, "go.mod")

	tests := []struct {
		name          string
		includes      []string
		excludes      []string
		expectedPaths []string
	}{
		{
			name:     "empty includes - all supported files, no excludes",
			includes: []string{},
			excludes: []string{},
			expectedPaths: []string{
				supported["package.json"],
				supported["pubspec.yaml"],
				supported["go.mod"],
				supported["requirements.txt"],
				nodeModulesFile,
				vendorFile,
			},
		},
		{
			name:          "includes python only, no excludes",
			includes:      []string{"python"},
			excludes:      []string{},
			expectedPaths: []string{supported["requirements.txt"]},
		},
		{
			name:          "includes node only, exclude node_modules dir",
			includes:      []string{"node"},
			excludes:      []string{"/node_modules/"},
			expectedPaths: []string{supported["package.json"]}, // nodeModulesFile should be skipped
		},
		{
			name:          "includes go only, exclude vendor dir",
			includes:      []string{"go"},
			excludes:      []string{"/vendor/"},
			expectedPaths: []string{supported["go.mod"]}, // vendorFile should be skipped
		},
		{
			name:          "includes node and go, exclude node_modules and vendor",
			includes:      []string{"node", "go"},
			excludes:      []string{"/node_modules/", "/vendor/"},
			expectedPaths: []string{supported["package.json"], supported["go.mod"]},
		},
		{
			name:     "includes empty (all), exclude vendor only",
			includes: []string{},
			excludes: []string{"/vendor/"},
			expectedPaths: []string{
				supported["package.json"],
				supported["pubspec.yaml"],
				supported["go.mod"],
				supported["requirements.txt"],
				nodeModulesFile, // node_modules is not excluded here
			},
		},
		{
			name:     "includes empty (all), exclude specific file path",
			includes: []string{},
			excludes: []string{vendorFile}, // exclude specific file path
			expectedPaths: []string{
				supported["package.json"],
				supported["pubspec.yaml"],
				supported["go.mod"],
				supported["requirements.txt"],
				nodeModulesFile,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePathChan := make(chan string)
			var foundPaths []string

			go WalkDirectories(tmpDir, tt.includes, tt.excludes, filePathChan)

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
