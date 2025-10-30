package role_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	grole "github.com/muhlba91/pulumi-shared-library/pkg/lib/google/iam/role"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateCustomRole(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "custom-basic"
		args := &grole.CustomRoleArgs{
			ID:          pulumi.String("myCustomRole"),
			Title:       pulumi.String("My Custom Role"),
			Description: pulumi.String("A custom role for tests"),
			Permissions: []pulumi.StringInput{pulumi.String("resourcemanager.projects.get")},
			Project:     pulumi.String("proj-123"),
		}

		role, err := grole.CreateCustomRole(ctx, name, args)
		require.NoError(t, err)
		require.NotNil(t, role)

		role.RoleId.ApplyT(func(id string) error {
			assert.Equal(t, "myCustomRole", id)
			return nil
		})
		role.Title.ApplyT(func(title string) error {
			assert.Equal(t, "My Custom Role", title)
			return nil
		})
		role.Description.ApplyT(func(desc *string) error {
			assert.Equal(t, "A custom role for tests", *desc)
			return nil
		})
		role.Project.ApplyT(func(p string) error {
			assert.Equal(t, "proj-123", p)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateCustomRole_Permissions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "custom-perms"
		perms := []string{
			"resourcemanager.projects.get",
			"resourcemanager.projects.list",
		}
		args := &grole.CustomRoleArgs{
			ID:          pulumi.String("permRole"),
			Title:       pulumi.String("Perm Role"),
			Description: pulumi.String("Role with multiple permissions"),
			Permissions: []pulumi.StringInput{
				pulumi.String("resourcemanager.projects.get"),
				pulumi.String("resourcemanager.projects.list"),
			},
			Project: pulumi.String("proj-456"),
		}

		role, err := grole.CreateCustomRole(ctx, name, args)
		require.NoError(t, err)
		require.NotNil(t, role)

		role.Permissions.ApplyT(func(p []string) error {
			assert.Equal(t, perms, p)
			return nil
		})
		role.Project.ApplyT(func(proj string) error {
			assert.Equal(t, "proj-456", proj)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateCustomRole_WithOptionalArgs(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "custom-basic"
		args := &grole.CustomRoleArgs{
			ID:            pulumi.String("myCustomRole"),
			Title:         pulumi.String("Perm Role"),
			Project:       pulumi.String("proj-123"),
			PulumiOptions: []pulumi.ResourceOption{},
		}

		role, err := grole.CreateCustomRole(ctx, name, args)
		require.NoError(t, err)
		require.NotNil(t, role)

		role.RoleId.ApplyT(func(id string) error {
			assert.Equal(t, "myCustomRole", id)
			return nil
		})
		role.Title.ApplyT(func(title string) error {
			assert.Equal(t, "Perm Role", title)
			return nil
		})
		role.Project.ApplyT(func(p string) error {
			assert.Equal(t, "proj-123", p)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
