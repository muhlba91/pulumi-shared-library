package storage_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/google/storage"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestUpload(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &storage.UploadOptions{
			BucketID: "test-bucket",
			Key:      "test-key",
		}

		bo, err := storage.Upload(ctx, opts)

		require.NoError(t, err)
		assert.NotNil(t, bo)

		pulumi.All(bo.URN(), bo.Name, bo.Content, bo.Source).ApplyT(func(all []interface{}) error {
			urn := all[0].(pulumi.URN)
			name := all[1].(string)
			content := all[2].(string)
			source := all[3]

			assert.Nil(t, source)
			assert.Empty(t, content)
			assert.Contains(t, string(urn), "gcs-object-test-bucket-test-key")
			assert.Equal(t, "test-key", name)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestUpload_WithOptionalArgs(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &storage.UploadOptions{
			BucketID:      "test-bucket",
			Key:           "test-key",
			PulumiOptions: []pulumi.ResourceOption{},
		}

		bo, err := storage.Upload(ctx, opts)

		require.NoError(t, err)
		assert.NotNil(t, bo)

		pulumi.All(bo.URN(), bo.Name, bo.Content, bo.Source).ApplyT(func(all []interface{}) error {
			urn := all[0].(pulumi.URN)
			name := all[1].(string)
			content := all[2].(string)
			source := all[3]

			assert.Nil(t, source)
			assert.Empty(t, content)
			assert.Contains(t, string(urn), "gcs-object-test-bucket-test-key")
			assert.Equal(t, "test-key", name)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestUploadWithContent(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		content := "test content"

		opts := &storage.UploadOptions{
			BucketID: "test-bucket",
			Key:      "test-key",
			Content:  &content,
		}

		bo, err := storage.Upload(ctx, opts)

		require.NoError(t, err)
		assert.NotNil(t, bo)

		pulumi.All(bo.URN(), bo.Name, bo.Content, bo.Source).ApplyT(func(all []interface{}) error {
			urn := all[0].(pulumi.URN)
			name := all[1].(string)
			content := all[2].(string)
			source := all[3]

			assert.Nil(t, source)
			assert.Equal(t, "test content", content)
			assert.Contains(t, string(urn), "gcs-object-test-bucket-test-key")
			assert.Equal(t, "test-key", name)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestUploadWithFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "in.txt")
	content := "hello read"

	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &storage.UploadOptions{
			BucketID: "test-bucket",
			Key:      "test-key",
			File:     &path,
		}

		bo, err := storage.Upload(ctx, opts)

		require.NoError(t, err)
		assert.NotNil(t, bo)

		pulumi.All(bo.URN(), bo.Name, bo.Source).ApplyT(func(all []interface{}) error {
			urn := all[0].(pulumi.URN)
			name := all[1].(string)
			source := all[2]

			assert.NotNil(t, source)
			assert.Contains(t, string(urn), "gcs-object-test-bucket-test-key")
			assert.Equal(t, "test-key", name)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
