package serviceaccount_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/google/iam/serviceaccount"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateServiceAccount_NoRoles(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "sa-no-roles"
		var roles []string
		project := "proj-no-roles"

		opts := &serviceaccount.CreateOptions{
			Name:    name,
			Roles:   roles,
			Project: pulumi.String(project),
		}

		sa, members, err := serviceaccount.CreateServiceAccount(ctx, opts)
		require.NoError(t, err)
		require.NotNil(t, sa)
		assert.Nil(t, members)
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateServiceAccount_NoRoles_WithOptionalArgs(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "sa-no-roles"
		var roles []string
		project := pulumi.String("proj-no-roles")

		opts := &serviceaccount.CreateOptions{
			Name:          name,
			Roles:         roles,
			Project:       project,
			PulumiOptions: []pulumi.ResourceOption{},
		}

		sa, members, err := serviceaccount.CreateServiceAccount(ctx, opts)
		require.NoError(t, err)
		require.NotNil(t, sa)
		assert.Nil(t, members)
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateServiceAccount_WithRoles(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "sa-with-roles"
		roles := []string{"roles/viewer", "roles/storage.objectAdmin"}
		project := pulumi.String("proj-123")

		opts := &serviceaccount.CreateOptions{
			Name:    name,
			Roles:   roles,
			Project: project,
		}

		sa, members, err := serviceaccount.CreateServiceAccount(ctx, opts)
		require.NoError(t, err)
		require.NotNil(t, sa)
		require.Len(t, members, len(roles))

		for i, mres := range members {
			expectedRole := roles[i]

			mres.Role.ApplyT(func(r string) error {
				assert.Equal(t, expectedRole, r)
				return nil
			})
			mres.Project.ApplyT(func(p string) error {
				assert.Equal(t, "proj-123", p)
				return nil
			})
			mres.Member.ApplyT(func(memberStr string) error {
				sa.Email.ApplyT(func(email string) error {
					assert.Equal(t, "serviceAccount:"+email, memberStr)
					return nil
				})
				return nil
			})
		}

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
