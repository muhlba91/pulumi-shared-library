package random_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/random"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreatePassword_Defaults(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		data, err := random.CreatePassword(ctx, "test", nil)
		require.NoError(err)
		require.NotNil(data)
		require.NotNil(data.Resource)

		// Assert the password value from the mocked resource
		data.Password.ApplyT(func(p string) error {
			assert.Equal("mocked-password-16", p)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(err)
}

func TestCreatePassword_CustomOptions(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	opts := &random.PasswordOptions{
		Length:  8,
		Special: false,
	}

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		data, err := random.CreatePassword(ctx, "custom", opts)
		require.NoError(err)
		require.NotNil(data)
		require.NotNil(data.Resource)

		// Assert the password value from the mocked resource
		data.Password.ApplyT(func(p string) error {
			assert.Equal("mocked-password-8", p)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(err)
}
