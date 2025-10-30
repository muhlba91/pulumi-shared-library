package role_test

import (
	"testing"

	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	librole "github.com/muhlba91/pulumi-shared-library/pkg/lib/aws/iam/role"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreatePolicyAttachment(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		// create two roles to attach the policy to
		r1, err := iam.NewRole(ctx, "role-a", &iam.RoleArgs{
			Name:             pulumi.String("role-a"),
			AssumeRolePolicy: pulumi.String(`{"Version":"2012-10-17","Statement":[]}`),
		})
		require.NoError(t, err)
		require.NotNil(t, r1)

		r2, err := iam.NewRole(ctx, "role-b", &iam.RoleArgs{
			Name:             pulumi.String("role-b"),
			AssumeRolePolicy: pulumi.String(`{"Version":"2012-10-17","Statement":[]}`),
		})
		require.NoError(t, err)
		require.NotNil(t, r2)

		policyArn := pulumi.String("arn:aws:iam::123456789012:policy/test-policy")

		pa, err := librole.CreatePolicyAttachment(ctx, "basic", &librole.CreatePolicyAttachmentOptions{
			Roles:     []pulumi.StringInput{r1.Name, r2.Name},
			PolicyArn: policyArn,
		})
		require.NoError(t, err)
		require.NotNil(t, pa)

		pa.PolicyArn.ApplyT(func(a string) error {
			assert.Equal(t, "arn:aws:iam::123456789012:policy/test-policy", a)
			return nil
		})
		pa.Roles.ApplyT(func(rs []string) error {
			assert.Equal(t, []string{"role-a", "role-b"}, rs)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreatePolicyAttachment_WithOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		r, err := iam.NewRole(ctx, "role-protected", &iam.RoleArgs{
			Name:             pulumi.String("role-protected"),
			AssumeRolePolicy: pulumi.String(`{"Version":"2012-10-17","Statement":[]}`),
		})
		require.NoError(t, err)
		require.NotNil(t, r)

		policyArn := pulumi.String("arn:aws:iam::123456789012:policy/prod-policy")

		pa, err := librole.CreatePolicyAttachment(ctx, "withopts", &librole.CreatePolicyAttachmentOptions{
			Roles:     []pulumi.StringInput{r.Name},
			PolicyArn: policyArn,
			PulumiOptions: []pulumi.ResourceOption{
				pulumi.Protect(true),
			},
		})
		require.NoError(t, err)
		require.NotNil(t, pa)

		pa.PolicyArn.ApplyT(func(a string) error {
			assert.Equal(t, "arn:aws:iam::123456789012:policy/prod-policy", a)
			return nil
		})
		pa.Roles.ApplyT(func(rs []string) error {
			assert.Equal(t, []string{"role-protected"}, rs)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
