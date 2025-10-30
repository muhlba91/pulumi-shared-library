package random_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/random"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateString_Defaults(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		data, err := random.CreateString(ctx, "test", nil)
		require.NoError(err)
		require.NotNil(data)
		require.NotNil(data.Resource)

		// Assert the string value from the mocked resource
		data.Text.ApplyT(func(p string) error {
			assert.Equal("mocked-string-16", p)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(err)
}

func TestCreateString_CustomOptions(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	opts := &random.StringOptions{
		Length:  8,
		Special: false,
	}

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		data, err := random.CreateString(ctx, "custom", opts)
		require.NoError(err)
		require.NotNil(data)
		require.NotNil(data.Resource)

		// Assert the string value from the mocked resource
		data.Text.ApplyT(func(p string) error {
			assert.Equal("mocked-string-8", p)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(err)
}
