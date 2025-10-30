package bucket_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libbucket "github.com/muhlba91/pulumi-shared-library/pkg/lib/aws/s3/bucket"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateBucket(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := libbucket.CreateOptions{
			Name: "basic",
		}

		b, err := libbucket.Create(ctx, opts)
		require.NoError(t, err)
		require.NotNil(t, b)

		// Assert tags contain the expected key/value set in Create
		b.Tags.ApplyT(func(m map[string]string) error {
			assert.Empty(t, m)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateBucket_WithOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := libbucket.CreateOptions{
			Name: "with-opts",
			Labels: map[string]string{
				"tag": "label",
			},
			PulumiOptions: []pulumi.ResourceOption{
				pulumi.Protect(true),
			},
		}

		b, err := libbucket.Create(ctx, opts)
		require.NoError(t, err)
		require.NotNil(t, b)

		b.Tags.ApplyT(func(m map[string]string) error {
			assert.Equal(t, "label", m["tag"])
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
