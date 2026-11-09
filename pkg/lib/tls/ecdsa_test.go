package tls_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libkey "github.com/muhlba91/pulumi-shared-library/pkg/lib/tls"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateECDSAKey_DefaultCurve(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		k, err := libkey.CreateECDSAKey(ctx, "default", "")
		require.NoError(t, err)
		require.NotNil(t, k)

		k.PrivateKeyPem.ApplyT(func(pem string) error {
			assert.Equal(t, "mocked-private-key-ECDSA-P256", pem)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Equal(t, []string{"ecdsa-key-default"}, counter.Resources["tls:index/privateKey:PrivateKey"])
}

func TestCreateECDSAKey_CustomCurve(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		k, err := libkey.CreateECDSAKey(ctx, "custom", "P521")
		require.NoError(t, err)

		k.PublicKeyPem.ApplyT(func(pub string) error {
			assert.Equal(t, "mocked-public-key-ECDSA-P521", pub)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}
