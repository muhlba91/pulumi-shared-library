package accesstoken_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libtoken "github.com/muhlba91/pulumi-shared-library/pkg/lib/pulumi/accesstoken"
	"github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateAccessToken(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		desc := pulumi.String("integration token")
		tok, err := libtoken.Create(ctx, "with-desc", &libtoken.CreateOptions{
			Description: desc,
		})
		require.NoError(t, err)
		require.NotNil(t, tok)

		tok.Value.ApplyT(func(v string) error {
			assert.NotEmpty(t, v)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Nil(t, counter.Resources["time:index/rotating:Rotating"])
	assert.Len(t, counter.Resources["pulumiservice:index:AccessToken"], 1)
}

func TestCreateAccessToken_WithRotation(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		desc := pulumi.String("integration token")
		tok, err := libtoken.Create(ctx, "with-desc", &libtoken.CreateOptions{
			Description: desc,
			Rotation:    &rotation.Options{},
		})
		require.NoError(t, err)
		require.NotNil(t, tok)

		tok.Value.ApplyT(func(v string) error {
			assert.NotEmpty(t, v)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Len(t, counter.Resources["time:index/rotating:Rotating"], 1)
	assert.Len(t, counter.Resources["pulumiservice:index:AccessToken"], 1)
}
