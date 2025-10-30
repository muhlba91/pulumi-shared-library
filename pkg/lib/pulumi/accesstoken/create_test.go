package accesstoken_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libtoken "github.com/muhlba91/pulumi-shared-library/pkg/lib/pulumi/accesstoken"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateAccessToken(t *testing.T) {
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
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
