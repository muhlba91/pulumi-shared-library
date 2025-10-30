package sanitize_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

func TestSanitize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"alphanumeric", "abc123", "abc123"},
		{"spaces", "a b", "a-b"},
		{"underscore", "a_b", "a-b"},
		{"punctuation", "A.B/C", "A-B-C"},
		{"unicode", "éü", "--"},
		{"empty", "", ""},
		{"multipleSpaces", "multiple  spaces", "multiple--spaces"},
		{"alreadyDashes", "a--b", "a--b"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := sanitize.Text(tc.input)
			assert.Equal(t, tc.expected, got)
		})
	}
}
