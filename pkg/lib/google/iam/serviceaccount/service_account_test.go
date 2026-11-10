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
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &serviceaccount.CreateOptions{
			Name:    "sa-no-roles",
			Project: pulumi.String("proj-no-roles"),
		}

		sa, members, err := serviceaccount.CreateServiceAccount(ctx, opts)
		require.NoError(t, err)
		require.NotNil(t, sa)
		assert.Nil(t, members)

		pulumi.All(sa.AccountId, sa.DisplayName, sa.Description, sa.CreateIgnoreAlreadyExists).
			ApplyT(func(all []any) error {
				assert.Equal(t, "sa-no-roles", all[0])
				assert.Equal(t, "sa-no-roles", *all[1].(*string))
				assert.Nil(t, all[2].(*string))
				assert.Nil(t, all[3].(*bool))
				return nil
			})
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Equal(t, []string{"gcp-sa-sa-no-roles"}, counter.Resources["gcp:serviceaccount/account:Account"])
	assert.Empty(t, counter.Resources["gcp:projects/iAMMember:IAMMember"])
}

func TestCreateServiceAccount_WithOptionalArgs(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		displayName := "Service Account"
		description := "a description"
		ignoreExisting := true

		opts := &serviceaccount.CreateOptions{
			Name:                      "sa-optional",
			DisplayName:               &displayName,
			Description:               &description,
			Project:                   pulumi.String("proj-optional"),
			CreateIgnoreAlreadyExists: &ignoreExisting,
			PulumiOptions:             []pulumi.ResourceOption{},
		}

		sa, members, err := serviceaccount.CreateServiceAccount(ctx, opts)
		require.NoError(t, err)
		require.NotNil(t, sa)
		assert.Nil(t, members)

		pulumi.All(sa.DisplayName, sa.Description, sa.CreateIgnoreAlreadyExists).ApplyT(func(all []any) error {
			assert.Equal(t, displayName, *all[0].(*string))
			assert.Equal(t, description, *all[1].(*string))
			assert.True(t, *all[2].(*bool))
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}

func TestCreateServiceAccount_WithRoles(t *testing.T) {
	counter := mocks.NewCounter()
	roles := []string{"roles/viewer", "roles/storage.objectAdmin"}

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &serviceaccount.CreateOptions{
			Name:    "sa-with-roles",
			Roles:   roles,
			Project: pulumi.String("proj-123"),
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
			pulumi.All(mres.Member, sa.Email).ApplyT(func(all []any) error {
				assert.Equal(t, "serviceAccount:"+all[1].(string), all[0])
				return nil
			})
		}

		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Len(t, counter.Resources["gcp:serviceaccount/account:Account"], 1)
	assert.Len(t, counter.Resources["gcp:projects/iAMMember:IAMMember"], len(roles))
}
