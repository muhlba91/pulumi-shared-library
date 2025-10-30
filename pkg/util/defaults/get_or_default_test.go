package defaults_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
)

func TestGetOrDefault(t *testing.T) {
	type args[T any] struct {
		v   *T
		def T
	}
	tests := []struct {
		name string
		args args[int]
		want int
	}{
		{"nil pointer", args[int]{nil, 42}, 42},
		{"non-nil pointer", args[int]{new(int), 42}, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, defaults.GetOrDefault(tc.args.v, tc.args.def))
		})
	}
}
