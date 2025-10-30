package location_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/hetzner/location"
)

func TestToDatacenter(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"fsn1", "fsn1", "fsn1-dc14"},
		{"nbg1", "nbg1", "nbg1-dc3"},
		{"unknown", "does-not-exist", "fsn1-dc14"},
		{"empty", "", "fsn1-dc14"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := location.ToDatacenter(&tc.input)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestToDatacenter_Nil(t *testing.T) {
	got := location.ToDatacenter(nil)
	assert.Equal(t, "fsn1-dc14", got)
}
