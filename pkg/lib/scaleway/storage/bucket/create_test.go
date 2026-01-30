package bucket_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libbucket "github.com/muhlba91/pulumi-shared-library/pkg/lib/scaleway/storage/bucket"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateBucket_Basic(t *testing.T) {
	req := require.New(t)
	assert := assert.New(t)

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &libbucket.CreateOptions{
			Location: pulumi.String("EU"),
			Labels:   map[string]string{"env": "test"},
		}

		b, err := libbucket.Create(ctx, "mybucket", opts)
		req.NoError(err)
		req.NotNil(b)

		b.Region.ApplyT(func(loc *string) error {
			assert.Equal("EU", *loc)
			return nil
		})

		b.Tags.ApplyT(func(m map[string]string) error {
			assert.Equal(map[string]string{"env": "test"}, m)
			return nil
		})

		b.ID().ApplyT(func(id string) error {
			assert.NotEmpty(id)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	req.NoError(err)
}

func TestCreateBucket_WithOptions(t *testing.T) {
	req := require.New(t)
	assert := assert.New(t)

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &libbucket.CreateOptions{
			Location:      pulumi.String("US"),
			Labels:        map[string]string{"owner": "ci"},
			PulumiOptions: []pulumi.ResourceOption{pulumi.Protect(true)},
		}

		b, err := libbucket.Create(ctx, "withopts", opts)
		req.NoError(err)
		req.NotNil(b)

		b.Region.ApplyT(func(loc *string) error {
			assert.Equal("US", *loc)
			return nil
		})
		b.Tags.ApplyT(func(m map[string]string) error {
			assert.Equal(map[string]string{"owner": "ci"}, m)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	req.NoError(err)
}
