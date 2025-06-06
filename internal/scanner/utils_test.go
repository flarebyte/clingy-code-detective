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
			name:     "empty includes allows all - package.json",
			includes: []string{},
			filename: "package.json",
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
			name:     "ts alias allows package.json",
			includes: []string{"ts"},
			filename: "package.json",
			want:     true,
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
