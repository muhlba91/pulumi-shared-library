package tls_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libkey "github.com/muhlba91/pulumi-shared-library/pkg/lib/tls"
	rModel "github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreate_RSADefaultBits(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		k, err := libkey.Create(ctx, "default", &libkey.CreateOptions{Algorithm: libkey.AlgorithmRSA})
		require.NoError(t, err)
		require.NotNil(t, k)

		k.PublicKeyPem.ApplyT(func(pub string) error {
			assert.Equal(t, "mocked-public-key-RSA-4096", pub)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Equal(t, []string{"rsa-key-default"}, counter.Resources["tls:index/privateKey:PrivateKey"])
}

func TestCreate_RSACustomBits(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		bits := 2048
		k, err := libkey.Create(ctx, "custom", &libkey.CreateOptions{
			Algorithm: libkey.AlgorithmRSA,
			RSABits:   &bits,
		})
		require.NoError(t, err)

		k.PrivateKeyPem.ApplyT(func(pem string) error {
			assert.Equal(t, "mocked-private-key-RSA-2048", pem)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}

func TestCreate_ED25519(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		k, err := libkey.Create(ctx, "github", &libkey.CreateOptions{Algorithm: libkey.AlgorithmED25519})
		require.NoError(t, err)

		k.PrivateKeyOpenssh.ApplyT(func(key string) error {
			assert.Equal(t, "mocked-private-key-openssh-ED25519", key)
			return nil
		})
		k.PublicKeyOpenssh.ApplyT(func(pub string) error {
			assert.Equal(t, "mocked-public-key-openssh-ED25519", pub)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Equal(t, []string{"ed25519-key-github"}, counter.Resources["tls:index/privateKey:PrivateKey"])
}

func TestCreate_ECDSA(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		def, err := libkey.Create(ctx, "default", &libkey.CreateOptions{Algorithm: libkey.AlgorithmECDSA})
		require.NoError(t, err)
		def.PublicKeyPem.ApplyT(func(pub string) error {
			assert.Equal(t, "mocked-public-key-ECDSA-P256", pub)
			return nil
		})

		curve := "P384"
		custom, err := libkey.Create(ctx, "custom", &libkey.CreateOptions{
			Algorithm:  libkey.AlgorithmECDSA,
			ECDSACurve: &curve,
		})
		require.NoError(t, err)
		custom.PublicKeyPem.ApplyT(func(pub string) error {
			assert.Equal(t, "mocked-public-key-ECDSA-P384", pub)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}

func TestCreate_Rotation(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		_, err := libkey.Create(ctx, "rotated", &libkey.CreateOptions{
			Algorithm: libkey.AlgorithmED25519,
			Rotation:  &rModel.Options{Days: 365},
		})
		require.NoError(t, err)
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Len(t, counter.Resources["time:index/rotating:Rotating"], 1)
}

func TestCreate_MissingAlgorithm(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		k, err := libkey.Create(ctx, "missing", &libkey.CreateOptions{})
		require.Error(t, err)
		assert.Nil(t, k)
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}

func TestCreate_UnsupportedAlgorithm(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		k, err := libkey.Create(ctx, "invalid", &libkey.CreateOptions{Algorithm: "DSA"})
		require.Error(t, err)
		assert.Nil(t, k)
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}
