package sorting_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/sorting"
)

func TestSortStrings(t *testing.T) {
	cases := []struct {
		name  string
		input []string
	}{
		{"nil slice", nil},
		{"empty slice", []string{}},
		{"single element", []string{"OnlyOne"}},
		{"mixed case", []string{"banana", "Apple", "cherry", "apple", "Banana"}},
		{"duplicates", []string{"b", "B", "a", "A", "b"}},
		{"unicode", []string{"Äpfel", "apfel", "Ápple", "apple"}},
		{"already sorted (different cases)", []string{"Apple", "banana", "Cherry"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			in := append([]string(nil), tc.input...)

			require.NotPanics(t, func() { sorting.SortStrings(in) })

			for i := 1; i < len(in); i++ {
				assert.LessOrEqual(t, strings.ToLower(in[i-1]), strings.ToLower(in[i]))
			}
		})
	}
}
