package utils

import (
	"os"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func GetServiceAccountJSON(filePath string) []byte {
	fileContent, _ := os.ReadFile(filePath)
	return fileContent
}

func RemoveSubstringInsensitive(input, substring string) string {
	// Create a regular expression that matches the substring case-insensitively
	pattern := regexp.MustCompile(`(?i)` + regexp.QuoteMeta(substring))
	// Replace all occurrences of the pattern with an empty string
	return pattern.ReplaceAllString(input, "")
}

func RemoveExtraSpaces(input string) string {
	// Compile a regex that matches one or more spaces
	spaceRegex := regexp.MustCompile(`\s+`)
	// Replace all sequences of spaces with a single space
	return spaceRegex.ReplaceAllString(input, " ")
}

func RemoveAccents(s string) string {
	isMn := func(r rune) bool {
		return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
	}
	t := norm.NFD.String(s)
	return strings.Map(func(r rune) rune {
		if isMn(r) {
			return -1
		}
		return r
	}, t)
}
