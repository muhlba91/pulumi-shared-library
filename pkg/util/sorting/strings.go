package sorting

import (
	"sort"
	"strings"
)

// SortStrings sorts the provided []string in a case-insensitive way.
func SortStrings(slice []string) {
	sort.Slice(slice, func(i, j int) bool {
		return strings.ToLower(slice[i]) < strings.ToLower(slice[j])
	})
}
