package scanner

import "testing"

func TestIsFileRequired(t *testing.T) {
	tests := []struct {
		name     string
		includes []string
		filename string
		want     bool
	}{
		{
			name:     "empty includes allows all - go.mod",
			includes: []string{},
			filename: "go.mod",
			want:     true,
		},
		{
			name:     "python only allows requirements.txt",
			includes: []string{"python"},
			filename: "requirements.txt",
			want:     true,
		},
		{
			name:     "python excludes go.mod",
			includes: []string{"python"},
			filename: "go.mod",
			want:     false,
		},
		{
			name:     "node allows package.json",
			includes: []string{"node"},
			filename: "package.json",
			want:     true,
		},
		{
			name:     "js alias allows package.json",
			includes: []string{"js"},
			filename: "package.json",
			want:     true,
		},
		{
			name:     "multiple includes - python and node - python file",
			includes: []string{"python", "node"},
			filename: "requirements.txt",
			want:     true,
		},
		{
			name:     "multiple includes - python and node - node file",
			includes: []string{"python", "node"},
			filename: "package.json",
			want:     true,
		},
		{
			name:     "multiple includes - python and node - go.mod excluded",
			includes: []string{"python", "node"},
			filename: "go.mod",
			want:     false,
		},
		{
			name:     "duplicates in includes - python python js - requirements.txt",
			includes: []string{"python", "python", "js"},
			filename: "requirements.txt",
			want:     true,
		},
		{
			name:     "duplicates in includes - python python js - package.json",
			includes: []string{"python", "python", "js"},
			filename: "package.json",
			want:     true,
		},
		{
			name:     "duplicates in includes - python python js - go.mod excluded",
			includes: []string{"python", "python", "js"},
			filename: "go.mod",
			want:     false,
		},
		{
			name:     "dart allows pubspec.yaml",
			includes: []string{"dart"},
			filename: "pubspec.yaml",
			want:     true,
		},
		{
			name:     "unsupported file is excluded",
			includes: []string{"go"},
			filename: "not_supported.txt",
			want:     false,
		},
		{
			name:     "unknown category ignored",
			includes: []string{"unknown"},
			filename: "go.mod",
			want:     false,
		},
		{
			name:     "case insensitive alias - JS",
			includes: []string{"JS"},
			filename: "package.json",
			want:     true,
		},
		{
			name:     "case insensitive filename",
			includes: []string{"python"},
			filename: "REQUIREMENTS.TXT",
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsFileRequired(tt.filename, tt.includes)
			if got != tt.want {
				t.Errorf("IsFileRequired(%q, %v) = %v; want %v", tt.filename, tt.includes, got, tt.want)
			}
		})
	}
}
