package user_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libuser "github.com/muhlba91/pulumi-shared-library/pkg/lib/scaleway/iam/user"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateUser(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "alice"
		email := "alice@email.com"
		labels := []string{"env:test"}

		u, err := libuser.Create(ctx, name, &libuser.CreateOptions{
			Email:  pulumi.String(email),
			Labels: labels,
		})
		require.NoError(t, err)
		require.NotNil(t, u)

		u.Username.ApplyT(func(n string) error {
			assert.Equal(t, name, n)
			return nil
		})
		u.Email.ApplyT(func(n string) error {
			assert.Equal(t, email, n)
			return nil
		})
		u.Tags.ApplyT(func(m []string) error {
			assert.Equal(t, labels, m)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
