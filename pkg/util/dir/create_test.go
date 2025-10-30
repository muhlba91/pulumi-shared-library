package dir_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/dir"
)

func TestCreate(t *testing.T) {
	cases := []struct {
		name  string
		setup func(t *testing.T) string // returns path to operate on
	}{
		{
			name: "non-existent path",
			setup: func(t *testing.T) string {
				return filepath.Join(t.TempDir(), "newdir")
			},
		},
		{
			name: "existing empty dir",
			setup: func(t *testing.T) string {
				p := filepath.Join(t.TempDir(), "emptydir")
				require.NoError(t, os.MkdirAll(p, 0o755))
				return p
			},
		},
		{
			name: "existing non-empty dir",
			setup: func(t *testing.T) string {
				p := filepath.Join(t.TempDir(), "nonempty")
				require.NoError(t, os.MkdirAll(p, 0o755))
				require.NoError(t, os.WriteFile(filepath.Join(p, "file.txt"), []byte("data"), 0o644))
				return p
			},
		},
		{
			name: "existing file at path",
			setup: func(t *testing.T) string {
				p := filepath.Join(t.TempDir(), "somefile")
				require.NoError(t, os.WriteFile(p, []byte("x"), 0o644))
				return p
			},
		},
		{
			name: "nested path parent missing",
			setup: func(t *testing.T) string {
				return filepath.Join(t.TempDir(), "a", "b", "c")
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			path := tc.setup(t)

			require.NoError(t, dir.Create(path))

			info, err := os.Stat(path)
			require.NoError(t, err)
			assert.True(t, info.IsDir())

			entries, err := os.ReadDir(path)
			require.NoError(t, err)
			assert.Empty(t, entries)

			leftover := filepath.Join(path, "file.txt")
			_, err = os.Stat(leftover)
			assert.True(t, os.IsNotExist(err))
		})
	}
}
