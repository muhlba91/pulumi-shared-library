package sanitize

import "regexp"

var nonAlnumRegex = regexp.MustCompile(`[^a-zA-Z0-9]`)

// Text replaces each non-alphanumeric character (anything except a-z, A-Z, and 0-9) with '-'.
// text: input string to sanitize.
func Text(text string) string {
	return nonAlnumRegex.ReplaceAllString(text, "-")
}
