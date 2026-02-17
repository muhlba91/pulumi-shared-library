package role_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/google/iam/role"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateMember_MultipleRoles(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "testname"
		member := "user:alice@example.com"
		roles := []string{"roles/viewer", "roles/storage.objectAdmin"}
		project := pulumi.String("proj-123")

		created, err := role.CreateMember(ctx, name, &role.MemberOptions{
			Member:  pulumi.String(member),
			Roles:   roles,
			Project: project,
		})
		require.NoError(t, err)
		require.Len(t, created, len(roles))

		for i, res := range created {
			expectedRole := roles[i]

			res.Role.ApplyT(func(r string) error {
				assert.Equal(t, expectedRole, r)
				return nil
			})
			res.Member.ApplyT(func(m string) error {
				assert.Equal(t, member, m)
				return nil
			})
			res.Project.ApplyT(func(p string) error {
				assert.Equal(t, "proj-123", p)
				return nil
			})
		}
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateMember_MultipleRoles_WithOptionalArgs(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "testname"
		member := "user:alice@example.com"
		roles := []string{"roles/viewer", "roles/storage.objectAdmin"}
		project := pulumi.String("proj-123")

		created, err := role.CreateMember(ctx, name, &role.MemberOptions{
			Member:        pulumi.String(member),
			Roles:         roles,
			Project:       project,
			PulumiOptions: []pulumi.ResourceOption{},
		})
		require.NoError(t, err)
		require.Len(t, created, len(roles))

		for i, res := range created {
			expectedRole := roles[i]

			res.Role.ApplyT(func(r string) error {
				assert.Equal(t, expectedRole, r)
				return nil
			})
			res.Member.ApplyT(func(m string) error {
				assert.Equal(t, member, m)
				return nil
			})
			res.Project.ApplyT(func(p string) error {
				assert.Equal(t, "proj-123", p)
				return nil
			})
		}
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
