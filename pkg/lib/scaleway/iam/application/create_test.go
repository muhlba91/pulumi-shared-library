package application_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libapplication "github.com/muhlba91/pulumi-shared-library/pkg/lib/scaleway/iam/application"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateUser(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "alice"
		email := "alice@email.com"
		labels := []string{"env:test"}

		u, err := libapplication.Create(ctx, name, &libapplication.CreateOptions{
			Description: pulumi.StringPtr(email),
			Labels:      labels,
		})
		require.NoError(t, err)
		require.NotNil(t, u)

		u.Name.ApplyT(func(n string) error {
			assert.Equal(t, name, n)
			return nil
		})
		u.Description.ApplyT(func(n *string) error {
			assert.Equal(t, email, *n)
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
