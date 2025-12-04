package storage_test

import (
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	utilstorage "github.com/muhlba91/pulumi-shared-library/pkg/util/storage"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestWriteFileAndUpload(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		var wg sync.WaitGroup

		name := "out_basic.txt"
		content := "hello world"
		outputPath := t.TempDir()
		bucketID := "bucket-1"
		bucketPath := "path/in/bucket"
		expectedObjectName := filepath.Join(bucketPath, name)

		args := &utilstorage.WriteFileAndUploadArgs{
			Name:       name,
			Content:    pulumi.String(content),
			OutputPath: outputPath,
			BucketID:   bucketID,
			BucketPath: bucketPath,
		}

		out := utilstorage.WriteFileAndUpload(ctx, args)
		assert.NotNil(t, out)
		wg.Add(1)

		out.ApplyT(func(v any) error {
			defer wg.Done()
			assert.NotNil(t, v)

			bo, ok := v.(*storage.BucketObject)
			assert.True(t, ok)
			assert.NotNil(t, bo)

			bo.Bucket.ApplyT(func(b string) error {
				assert.Equal(t, bucketID, b)
				return nil
			})
			bo.Name.ApplyT(func(n string) error {
				assert.Equal(t, expectedObjectName, n)
				return nil
			})

			data, err := os.ReadFile(filepath.Join(outputPath, name))
			require.NoError(t, err)
			assert.Equal(t, content, string(data))
			return nil
		})

		wg.Wait()
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
