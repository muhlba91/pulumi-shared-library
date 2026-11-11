package storage_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/scaleway/storage"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestUploadContentInputIsSecret(t *testing.T) {
	for name, input := range map[string]pulumi.StringPtrInput{
		"plain":  pulumi.StringPtr("x"),
		"secret": pulumi.ToSecret(pulumi.StringPtr("x")).(pulumi.StringPtrOutput),
	} {
		t.Run(name, func(t *testing.T) {
			rec := mocks.NewRecorder()
			err := pulumi.RunErr(func(ctx *pulumi.Context) error {
				_, err := storage.Upload(ctx, &storage.UploadOptions{
					BucketID: "test-bucket", Key: "test-key", ContentInput: input,
				})

				return err
			}, pulumi.WithMocks("project", "stack", rec))
			require.NoError(t, err)

			in, ok := rec.Inputs("scaleway-object-test-bucket-test-key")
			require.True(t, ok)
			assert.True(t, in["content"].IsSecret())
		})
	}
}

func TestUploadContentIsNotForcedSecret(t *testing.T) {
	rec := mocks.NewRecorder()
	content := "x"
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		_, err := storage.Upload(ctx, &storage.UploadOptions{
			BucketID: "test-bucket", Key: "test-key", Content: &content,
		})

		return err
	}, pulumi.WithMocks("project", "stack", rec))
	require.NoError(t, err)

	in, ok := rec.Inputs("scaleway-object-test-bucket-test-key")
	require.True(t, ok)
	assert.False(t, in["content"].IsSecret())
}
