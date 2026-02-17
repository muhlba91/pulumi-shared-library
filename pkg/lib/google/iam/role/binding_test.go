package role_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	grole "github.com/muhlba91/pulumi-shared-library/pkg/lib/google/iam/role"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateBinding(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "binding-basic"
		serviceAccount := pulumi.String("my-service-account")
		role := pulumi.String("roles/iam.serviceAccountUser")
		members := []pulumi.StringInput{pulumi.String("user:alice@example.com")}

		opts := &grole.BindingOptions{
			ServiceAccount: serviceAccount,
			Role:           role,
			Members:        members,
		}

		binding, err := grole.CreateBinding(ctx, name, opts)
		require.NoError(t, err)
		require.NotNil(t, binding)

		binding.ServiceAccountId.ApplyT(func(s string) error {
			assert.Equal(t, "my-service-account", s)
			return nil
		})
		binding.Role.ApplyT(func(r string) error {
			assert.Equal(t, "roles/iam.serviceAccountUser", r)
			return nil
		})
		binding.Members.ApplyT(func(m []string) error {
			assert.Equal(t, []string{"user:alice@example.com"}, m)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateBinding_MultipleMembers(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "binding-multi"
		serviceAccount := pulumi.String("svc-account")
		role := pulumi.String("roles/iam.serviceAccountKeyAdmin")
		members := []pulumi.StringInput{
			pulumi.String("serviceAccount:svc+proj@example.com"),
			pulumi.String("group:devs@example.com"),
		}

		opts := &grole.BindingOptions{
			ServiceAccount: serviceAccount,
			Role:           role,
			Members:        members,
		}

		binding, err := grole.CreateBinding(ctx, name, opts)
		require.NoError(t, err)
		require.NotNil(t, binding)

		binding.ServiceAccountId.ApplyT(func(s string) error {
			assert.Equal(t, "svc-account", s)
			return nil
		})
		binding.Role.ApplyT(func(r string) error {
			assert.Equal(t, "roles/iam.serviceAccountKeyAdmin", r)
			return nil
		})
		binding.Members.ApplyT(func(m []string) error {
			assert.Equal(t, []string{"serviceAccount:svc+proj@example.com", "group:devs@example.com"}, m)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateBinding_WithOptionalArgs(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "binding-basic"
		serviceAccount := pulumi.String("my-service-account")
		role := pulumi.String("roles/iam.serviceAccountUser")
		members := []pulumi.StringInput{pulumi.String("user:alice@example.com")}

		opts := &grole.BindingOptions{
			ServiceAccount: serviceAccount,
			Role:           role,
			Members:        members,
			PulumiOptions:  []pulumi.ResourceOption{},
		}

		binding, err := grole.CreateBinding(ctx, name, opts)
		require.NoError(t, err)
		require.NotNil(t, binding)

		binding.ServiceAccountId.ApplyT(func(s string) error {
			assert.Equal(t, "my-service-account", s)
			return nil
		})
		binding.Role.ApplyT(func(r string) error {
			assert.Equal(t, "roles/iam.serviceAccountUser", r)
			return nil
		})
		binding.Members.ApplyT(func(m []string) error {
			assert.Equal(t, []string{"user:alice@example.com"}, m)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
