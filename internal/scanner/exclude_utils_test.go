package scanner

import "testing"

func TestIsFileExcluded(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		excludes []string
		want     bool
	}{
		{
			name:     "no excludes, path should not be excluded",
			filepath: "/home/user/project/main.go",
			excludes: []string{},
			want:     false,
		},
		{
			name:     "path matches /node_modules/",
			filepath: "/home/user/project/node_modules/foo.js",
			excludes: []string{"/node_modules/"},
			want:     true,
		},
		{
			name:     "path matches /vendor/",
			filepath: "/home/user/project/vendor/bar.go",
			excludes: []string{"/vendor/"},
			want:     true,
		},
		{
			name:     "path does not match any exclude",
			filepath: "/home/user/project/src/utils/helper.go",
			excludes: []string{"/node_modules/", "/vendor/"},
			want:     false,
		},
		{
			name:     "multiple excludes, match on second one",
			filepath: "/home/user/project/build/tmp/output.o",
			excludes: []string{"/node_modules/", "/build/tmp/"},
			want:     true,
		},
		{
			name:     "substring match works anywhere in path",
			filepath: "/home/user/project/tmp-build/obj.o",
			excludes: []string{"tmp-build"},
			want:     true,
		},
		{
			name:     "case sensitive match - no match",
			filepath: "/home/user/project/Node_Modules/foo.js",
			excludes: []string{"/node_modules/"},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsFileExcluded(tt.filepath, tt.excludes)
			if got != tt.want {
				t.Errorf("IsFileExcluded(%q, %v) = %v; want %v", tt.filepath, tt.excludes, got, tt.want)
			}
		})
	}
}
