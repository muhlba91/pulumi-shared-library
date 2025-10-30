package accesskey_test

import (
	"testing"

	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libaccesskey "github.com/muhlba91/pulumi-shared-library/pkg/lib/aws/iam/accesskey"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateAccessKey(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		userName := "test-user"

		user, err := iam.NewUser(ctx, "user", &iam.UserArgs{
			Name: pulumi.String(userName),
		})
		require.NoError(t, err)
		require.NotNil(t, user)

		ak, err := libaccesskey.Create(ctx, &libaccesskey.CreateOptions{
			UserName: userName,
			User:     user,
		})
		require.NoError(t, err)
		require.NotNil(t, ak)

		ak.User.ApplyT(func(u string) error {
			assert.Equal(t, userName, u)
			return nil
		})
		ak.Secret.ApplyT(func(s string) error {
			assert.NotEmpty(t, s)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateAccessKey_WithOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		userName := "protected-user"

		user, err := iam.NewUser(ctx, "user2", &iam.UserArgs{
			Name: pulumi.String(userName),
		})
		require.NoError(t, err)
		require.NotNil(t, user)

		ak, err := libaccesskey.Create(ctx, &libaccesskey.CreateOptions{
			UserName: userName,
			User:     user,
			PulumiOptions: []pulumi.ResourceOption{
				pulumi.Protect(true),
			},
		})
		require.NoError(t, err)
		require.NotNil(t, ak)

		ak.User.ApplyT(func(u string) error {
			assert.Equal(t, userName, u)
			return nil
		})
		ak.Secret.ApplyT(func(s string) error {
			assert.NotEmpty(t, s)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
