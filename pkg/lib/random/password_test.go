package random_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/random"
	"github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreatePassword_Defaults(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		data, err := random.CreatePassword(ctx, "test", nil)
		require.NoError(t, err)
		require.NotNil(t, data)
		require.NotNil(t, data.Resource)

		// Assert the password value from the mocked resource
		data.Password.ApplyT(func(p string) error {
			assert.Equal(t, "mocked-password-16", p)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Nil(t, counter.Resources["time:index/rotating:Rotating"])
	assert.Len(t, counter.Resources["random:index/randomPassword:RandomPassword"], 1)
}

func TestCreatePassword_DefaultOptions(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		data, err := random.CreatePassword(ctx, "test", &random.PasswordOptions{})
		require.NoError(t, err)
		require.NotNil(t, data)
		require.NotNil(t, data.Resource)

		// Assert the password value from the mocked resource
		data.Password.ApplyT(func(p string) error {
			assert.Equal(t, "mocked-password-16", p)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Nil(t, counter.Resources["time:index/rotating:Rotating"])
	assert.Len(t, counter.Resources["random:index/randomPassword:RandomPassword"], 1)
}

func TestCreatePassword_CustomOptions(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &random.PasswordOptions{
			Length:   8,
			Special:  false,
			Rotation: &rotation.Options{},
		}
		data, err := random.CreatePassword(ctx, "custom", opts)
		require.NoError(t, err)
		require.NotNil(t, data)
		require.NotNil(t, data.Resource)

		// Assert the password value from the mocked resource
		data.Password.ApplyT(func(p string) error {
			assert.Equal(t, "mocked-password-8", p)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Len(t, counter.Resources["time:index/rotating:Rotating"], 1)
	assert.Len(t, counter.Resources["random:index/randomPassword:RandomPassword"], 1)
}
