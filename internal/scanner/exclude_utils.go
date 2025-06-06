package scanner

import "strings"

// IsFileExcluded returns true if filepath contains any of the exclude substrings.
func IsFileExcluded(filepath string, excludes []string) bool {
	for _, ex := range excludes {
		if strings.Contains(filepath, ex) {
			return true
		}
	}
	return false
}
