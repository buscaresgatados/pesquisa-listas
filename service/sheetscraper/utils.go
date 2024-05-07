package sheetscraper

import "regexp"

func removeSubstringInsensitive(input, substring string) string {
	// Create a regular expression that matches the substring case-insensitively
	pattern := regexp.MustCompile(`(?i)` + regexp.QuoteMeta(substring))
	// Replace all occurrences of the pattern with an empty string
	return pattern.ReplaceAllString(input, "")
}

func removeExtraSpaces(input string) string {
	// Compile a regex that matches one or more spaces
	spaceRegex := regexp.MustCompile(`\s+`)
	// Replace all sequences of spaces with a single space
	return spaceRegex.ReplaceAllString(input, " ")
}
