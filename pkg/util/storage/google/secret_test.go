package google_test

import (
	"errors"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/storage"
	utilstorage "github.com/muhlba91/pulumi-shared-library/pkg/util/storage/google"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

const resourceName = "gcs-object-b-p-k"

func runUpload(t *testing.T, rec *mocks.Recorder, content pulumi.StringInput, outputPath string, wait bool) error {
	t.Helper()

	return pulumi.RunErr(func(ctx *pulumi.Context) error {
		out := utilstorage.WriteFileAndUpload(ctx, &storage.WriteFileAndUploadOptions{
			Name:        "k",
			Content:     content,
			OutputPath:  outputPath,
			BucketID:    "b",
			BucketPath:  "p",
			Permissions: []os.FileMode{0o600},
		})
		if wait {
			var wg sync.WaitGroup
			wg.Add(1)
			out.ApplyT(func(any) int { wg.Done(); return 0 })
			wg.Wait()
		}

		return nil
	}, pulumi.WithMocks("project", "stack", rec))
}

func TestWriteFileAndUploadSecretInput(t *testing.T) {
	rec := mocks.NewRecorder()
	dir := t.TempDir()

	// not awaiting the returned output: the object must be registered anyway.
	err := runUpload(t, rec, pulumi.ToSecret(pulumi.String("x")).(pulumi.StringOutput), dir, false)
	require.NoError(t, err)

	in, ok := rec.Inputs(resourceName)
	require.True(t, ok, "object must be registered with the unchanged resource name")
	assert.True(t, in["content"].IsSecret())

	b, err := os.ReadFile(filepath.Join(dir, "k"))
	require.NoError(t, err)
	assert.Equal(t, "x", string(b))
}

func TestWriteFileAndUploadPlainInputIsSecret(t *testing.T) {
	rec := mocks.NewRecorder()

	require.NoError(t, runUpload(t, rec, pulumi.String("x"), t.TempDir(), true))

	in, ok := rec.Inputs(resourceName)
	require.True(t, ok)
	assert.True(t, in["content"].IsSecret())
}

func TestWriteFileAndUploadRejectedContentFails(t *testing.T) {
	bad := pulumi.String("x").ToStringOutput().
		ApplyT(func(string) (string, error) { return "", errors.New("boom") }).(pulumi.StringOutput)

	err := runUpload(t, mocks.NewRecorder(), bad, t.TempDir(), false)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "boom")
}

func TestWriteFileAndUploadFailedWriteFails(t *testing.T) {
	// the parent "directory" is a regular file, so the write cannot succeed.
	blocker := filepath.Join(t.TempDir(), "file")
	require.NoError(t, os.WriteFile(blocker, []byte("x"), 0o600))

	err := runUpload(t, mocks.NewRecorder(), pulumi.String("x"), filepath.Join(blocker, "sub"), false)
	require.Error(t, err)
}

func TestWriteFileAndUploadCreatesDirAndKeepsBytes(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "a", "b")
	content := "line1\nline2 ü\n\n"

	require.NoError(t, runUpload(t, mocks.NewRecorder(), pulumi.String(content), dir, true))

	b, err := os.ReadFile(filepath.Join(dir, "k"))
	require.NoError(t, err)
	assert.Equal(t, content, string(b))
}

func TestWriteFileAndUploadPreviewRegistersObject(t *testing.T) {
	rec := mocks.NewRecorder()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		utilstorage.WriteFileAndUpload(ctx, &storage.WriteFileAndUploadOptions{
			Name: "k", Content: pulumi.String("x"), OutputPath: t.TempDir(), BucketID: "b", BucketPath: "p",
		})

		return nil
	}, pulumi.WithMocks("project", "stack", rec), func(info *pulumi.RunInfo) { info.DryRun = true })
	require.NoError(t, err)

	in, ok := rec.Inputs(resourceName)
	require.True(t, ok, "object must be registered during preview")
	assert.True(t, in["content"].IsSecret())
}
