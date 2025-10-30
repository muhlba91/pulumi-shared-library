package file_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/file"
)

func TestReadContents(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "in.txt")
	content := "hello read"

	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	got, err := file.ReadContents(path)
	require.NoError(t, err)
	assert.Equal(t, content, got)
}

func TestReadNotExist(t *testing.T) {
	path := filepath.Join(t.TempDir(), "does_not_exist.txt")
	got, err := file.ReadContents(path)
	require.Error(t, err)
	assert.Empty(t, got)
}
