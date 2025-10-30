package tls_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libkey "github.com/muhlba91/pulumi-shared-library/pkg/lib/tls"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateSSHKey_DefaultBits(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		k, err := libkey.CreateSSHKey(ctx, "default", 0)
		require.NoError(t, err)
		require.NotNil(t, k)

		k.PrivateKeyPem.ApplyT(func(pem string) error {
			assert.Equal(t, "mocked-private-key-RSA-4096", pem)
			return nil
		})
		k.PublicKeyPem.ApplyT(func(pub string) error {
			assert.Equal(t, "mocked-public-key-RSA-4096", pub)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateSSHKey_CustomBits(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		k, err := libkey.CreateSSHKey(ctx, "custom", 2048)
		require.NoError(t, err)
		require.NotNil(t, k)

		k.PrivateKeyPem.ApplyT(func(pem string) error {
			assert.Equal(t, "mocked-private-key-RSA-2048", pem)
			return nil
		})
		k.PublicKeyPem.ApplyT(func(pub string) error {
			assert.Equal(t, "mocked-public-key-RSA-2048", pub)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
