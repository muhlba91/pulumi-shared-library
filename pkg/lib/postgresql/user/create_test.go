package user_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libuser "github.com/muhlba91/pulumi-shared-library/pkg/lib/postgresql/user"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreate_PostgresUser_Basic(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		username := "pguser"

		ud, err := libuser.Create(ctx, &libuser.CreateOptions{
			Username: username,
		})
		require.NoError(t, err)
		require.NotNil(t, ud)
		require.NotNil(t, ud.Password)
		require.NotNil(t, ud.User)

		ud.Password.ApplyT(func(pw string) error {
			assert.NotEmpty(t, pw)
			return nil
		})
		ud.User.Name.ApplyT(func(n string) error {
			assert.Equal(t, username, n)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreate_PostgresUser_WithOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		username := "pguser2"

		ud, err := libuser.Create(ctx, &libuser.CreateOptions{
			Username: username,
			PulumiOptions: []pulumi.ResourceOption{
				pulumi.Protect(true),
			},
		})
		require.NoError(t, err)
		require.NotNil(t, ud)
		require.NotNil(t, ud.Password)
		require.NotNil(t, ud.User)

		ud.Password.ApplyT(func(pw string) error {
			assert.NotEmpty(t, pw)
			return nil
		})
		ud.User.Name.ApplyT(func(n string) error {
			assert.Equal(t, username, n)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
