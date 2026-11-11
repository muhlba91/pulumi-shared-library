package file_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/file"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestWriteContents(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.txt")
	content := "hello world"

	got, err := file.WriteContents(path, content)
	require.NoError(t, err)
	assert.Equal(t, content, got)

	b, err := os.ReadFile(path)
	require.NoError(t, err)
	assert.Equal(t, content, string(b))
}

func TestWriteContentsPermissions(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.txt")
	content := "hello world"

	got, err := file.WriteContents(path, content, 0o600)
	require.NoError(t, err)
	assert.Equal(t, content, got)

	info, err := os.Stat(path)
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0o600), info.Mode().Perm())
}

func TestWriteContentsError(t *testing.T) {
	dir := t.TempDir()
	path := dir

	got, err := file.WriteContents(dir, "should fail")
	require.Error(t, err)
	assert.Equal(t, "should fail", got)

	info, statErr := os.Stat(path)
	require.NoError(t, statErr)
	assert.True(t, info.IsDir())
}

func TestWritePulumi(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "pulumi_out.txt")
	content := "pulumi content"

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		out := file.WritePulumi(path, pulumi.String(content))
		ctx.Export("written", out)
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)

	b, err := os.ReadFile(path)
	require.NoError(t, err)
	assert.Equal(t, content, string(b))
}

func TestWritePulumiPermissions(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "pulumi_out.txt")
	content := "pulumi content"

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		out := file.WritePulumi(path, pulumi.String(content), 0o600)
		ctx.Export("written", out)
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)

	b, err := os.ReadFile(path)
	require.NoError(t, err)
	assert.Equal(t, content, string(b))

	info, err := os.Stat(path)
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0o600), info.Mode().Perm())
}

func TestWritePulumiError(t *testing.T) {
	dir := t.TempDir()
	path := dir

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		out := file.WritePulumi(path, pulumi.String("should fail"))
		ctx.Export("written", out)
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)

	info, statErr := os.Stat(path)
	require.NoError(t, statErr)
	assert.True(t, info.IsDir())
}

func TestWritePulumiStrict(t *testing.T) {
	path := filepath.Join(t.TempDir(), "nested", "dir", "out.txt")

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		ctx.Export("written", file.WritePulumiStrict(path, pulumi.String("a\n"), 0o600))
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)

	b, err := os.ReadFile(path)
	require.NoError(t, err)
	assert.Equal(t, "a\n", string(b))
	info, err := os.Stat(path)
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0o600), info.Mode().Perm())
}

func TestWritePulumiStrictFailsRun(t *testing.T) {
	dir := t.TempDir() // a directory cannot be written as a file

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		ctx.Export("written", file.WritePulumiStrict(dir, pulumi.String("x")))
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.Error(t, err)
}
