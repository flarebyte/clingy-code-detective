package aggregator

import "strings"

func EscapeMarkdown(s string) string {
	return strings.ReplaceAll(s, "|", "\\|")
}
