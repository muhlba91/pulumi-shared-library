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

func TestCreateBytes_Defaults(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		data, err := random.CreateBytes(ctx, "test", nil)
		require.NoError(t, err)
		require.NotNil(t, data)
		require.NotNil(t, data.Resource)

		data.Hex.ApplyT(func(p string) error {
			assert.Equal(t, "mocked-bytes-hex-16", p)
			return nil
		})
		data.Base64Std.ApplyT(func(p string) error {
			assert.Equal(t, "mocked-bytes-b64std-16", p)
			return nil
		})
		data.Base64URL.ApplyT(func(p string) error {
			assert.Equal(t, "mocked-bytes-b64url-16", p)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Len(t, counter.Resources["random:index/randomId:RandomId"], 1)
}

func TestCreateBytes_CustomOptions(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		data, err := random.CreateBytes(ctx, "custom", &random.BytesOptions{Length: 8})
		require.NoError(t, err)
		require.NotNil(t, data)
		require.NotNil(t, data.Resource)

		data.Hex.ApplyT(func(p string) error {
			assert.Equal(t, "mocked-bytes-hex-8", p)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Nil(t, counter.Resources["time:index/rotating:Rotating"])
	assert.Len(t, counter.Resources["random:index/randomId:RandomId"], 1)
}

func TestCreateBytes_WithRotation(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		data, err := random.CreateBytes(ctx, "rotating", &random.BytesOptions{
			Length:   8,
			Rotation: &rotation.Options{},
		})
		require.NoError(t, err)
		require.NotNil(t, data)
		require.NotNil(t, data.Resource)
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Len(t, counter.Resources["time:index/rotating:Rotating"], 1)
	assert.Len(t, counter.Resources["random:index/randomId:RandomId"], 1)
}

func TestCreateBytes_EmptyOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		data, err := random.CreateBytes(ctx, "empty", &random.BytesOptions{})
		require.NoError(t, err)

		data.Hex.ApplyT(func(p string) error {
			assert.Equal(t, "mocked-bytes-hex-16", p)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}
