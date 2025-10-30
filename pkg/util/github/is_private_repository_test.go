package github_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	utilgithub "github.com/muhlba91/pulumi-shared-library/pkg/util/github"
)

func TestIsPrivateRepository(t *testing.T) {
	priv := "private"
	pub := "public"
	upper := "PRIVATE"

	tests := []struct {
		name     string
		input    *string
		expected bool
	}{
		{"private", &priv, true},
		{"public", &pub, false},
		{"nil", nil, false},
		{"uppercase_private", &upper, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := utilgithub.IsPrivateRepository(tc.input)
			require.NotNil(t, t) // keep require usage consistent; no-op guard
			assert.Equal(t, tc.expected, got)
		})
	}
}
