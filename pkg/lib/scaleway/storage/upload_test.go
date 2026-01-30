package storage_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/scaleway/storage"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestUpload(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		args := &storage.UploadArgs{
			BucketID: "test-bucket",
			Key:      "test-key",
		}

		bo, err := storage.Upload(ctx, args)

		require.NoError(t, err)
		assert.NotNil(t, bo)

		pulumi.All(bo.URN(), bo.Key, bo.Content, bo.File).ApplyT(func(all []interface{}) error {
			urn := all[0].(pulumi.URN)
			key := all[1].(string)
			content := all[2].(*string)
			file := all[3].(*string)

			assert.Nil(t, file)
			assert.Nil(t, content)
			assert.Contains(t, string(urn), "scaleway-object-test-bucket-test-key")
			assert.Equal(t, "test-key", key)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestUpload_WithOptionalArgs(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		args := &storage.UploadArgs{
			BucketID:      "test-bucket",
			Key:           "test-key",
			PulumiOptions: []pulumi.ResourceOption{},
		}

		bo, err := storage.Upload(ctx, args)

		require.NoError(t, err)
		assert.NotNil(t, bo)

		pulumi.All(bo.URN(), bo.Key, bo.Content, bo.File).ApplyT(func(all []interface{}) error {
			urn := all[0].(pulumi.URN)
			key := all[1].(string)
			content := all[2].(*string)
			file := all[3].(*string)

			assert.Nil(t, file)
			assert.Nil(t, content)
			assert.Contains(t, string(urn), "scaleway-object-test-bucket-test-key")
			assert.Equal(t, "test-key", key)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestUploadWithContent(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		content := "test content"

		args := &storage.UploadArgs{
			BucketID: "test-bucket",
			Key:      "test-key",
			Content:  &content,
		}

		bo, err := storage.Upload(ctx, args)

		require.NoError(t, err)
		assert.NotNil(t, bo)

		pulumi.All(bo.URN(), bo.Key, bo.Content, bo.File).ApplyT(func(all []interface{}) error {
			urn := all[0].(pulumi.URN)
			key := all[1].(string)
			content := all[2].(*string)
			file := all[3].(*string)

			assert.Nil(t, file)
			assert.Equal(t, "test content", *content)
			assert.Contains(t, string(urn), "scaleway-object-test-bucket-test-key")
			assert.Equal(t, "test-key", key)
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
		args := &storage.UploadArgs{
			BucketID: "test-bucket",
			Key:      "test-key",
			File:     &path,
		}

		bo, err := storage.Upload(ctx, args)

		require.NoError(t, err)
		assert.NotNil(t, bo)

		pulumi.All(bo.URN(), bo.Key, bo.File).ApplyT(func(all []interface{}) error {
			urn := all[0].(pulumi.URN)
			key := all[1].(string)
			file := all[2].(*string)

			assert.NotNil(t, file)
			assert.Contains(t, string(urn), "scaleway-object-test-bucket-test-key")
			assert.Equal(t, "test-key", key)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
