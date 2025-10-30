package secret_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/vault/secret"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestWrite_Defaults(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		args := &secret.WriteArgs{
			Key:   "mykey",
			Value: pulumi.String("my-value"),
			Path:  "secret",
		}

		secret, err := secret.Write(ctx, args)

		require.NoError(err)
		require.NotNil(secret)
		require.NotNil(secret.ID())

		secret.DataJson.ApplyT(func(s *string) error {
			assert.Equal("my-value", *s)
			return nil
		})

		secret.Mount.ApplyT(func(m string) error {
			assert.Equal("secret", m)
			return nil
		})
		secret.Name.ApplyT(func(n string) error {
			assert.Equal("mykey", n)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(err)
}

func TestWrite_Defaults_WithOptionalArgs(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		args := &secret.WriteArgs{
			Key:           "mykey",
			Value:         pulumi.String("my-value"),
			Path:          "secret",
			PulumiOptions: []pulumi.ResourceOption{},
		}

		secret, err := secret.Write(ctx, args)

		require.NoError(err)
		require.NotNil(secret)
		require.NotNil(secret.ID())

		secret.DataJson.ApplyT(func(s *string) error {
			assert.Equal("my-value", *s)
			return nil
		})

		secret.Mount.ApplyT(func(m string) error {
			assert.Equal("secret", m)
			return nil
		})
		secret.Name.ApplyT(func(n string) error {
			assert.Equal("mykey", n)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(err)
}
