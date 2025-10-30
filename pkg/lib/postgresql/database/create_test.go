package database_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libdb "github.com/muhlba91/pulumi-shared-library/pkg/lib/postgresql/database"
	libuser "github.com/muhlba91/pulumi-shared-library/pkg/lib/postgresql/user"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreate_PostgresDatabase(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		username := "dbowner"
		ud, err := libuser.Create(ctx, &libuser.CreateOptions{
			Username: username,
		})
		require.NoError(t, err)
		require.NotNil(t, ud)

		db, err := libdb.Create(ctx, &libdb.CreateOptions{
			Name:  "mydb",
			Owner: *ud,
		})
		require.NoError(t, err)
		require.NotNil(t, db)

		db.Name.ApplyT(func(n string) error {
			assert.Equal(t, "mydb", n)
			return nil
		})

		// db.Owner should equal the role name of the created user
		db.Owner.ApplyT(func(ownerName string) error {
			ud.User.Name.ApplyT(func(exp string) error {
				assert.Equal(t, exp, ownerName)
				return nil
			})
			return nil
		})

		// ensure the generated password is non-empty
		ud.Password.ApplyT(func(p string) error {
			assert.NotEmpty(t, p)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreate_PostgresDatabase_WithOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		username := "dbowner2"
		ud, err := libuser.Create(ctx, &libuser.CreateOptions{
			Username: username,
		})
		require.NoError(t, err)
		require.NotNil(t, ud)

		db, err := libdb.Create(ctx, &libdb.CreateOptions{
			Name:  "mydb2",
			Owner: *ud,
			PulumiOptions: []pulumi.ResourceOption{
				pulumi.Protect(true),
			},
		})
		require.NoError(t, err)
		require.NotNil(t, db)

		db.Name.ApplyT(func(n string) error {
			assert.Equal(t, "mydb2", n)
			return nil
		})

		db.Owner.ApplyT(func(ownerName string) error {
			ud.User.Name.ApplyT(func(exp string) error {
				assert.Equal(t, exp, ownerName)
				return nil
			})
			return nil
		})

		ud.Password.ApplyT(func(p string) error {
			assert.NotEmpty(t, p)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
