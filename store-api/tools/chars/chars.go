package chars

import (
	"regexp"
	"strings"
)

func ToLatin(str string) string {
	str = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(str, "")
	str = regexp.MustCompile("\\s+").ReplaceAllString(str, " ")
	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, " ", "-")
	return str
}
