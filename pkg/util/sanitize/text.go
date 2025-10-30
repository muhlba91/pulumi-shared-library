package sanitize

import "regexp"

var nonAlnumRegex = regexp.MustCompile(`[^a-zA-Z0-9]`)

// Text replaces all non-alphanumeric characters with '-'.
// text: input string to sanitize.
func Text(text string) string {
	return nonAlnumRegex.ReplaceAllString(text, "-")
}
