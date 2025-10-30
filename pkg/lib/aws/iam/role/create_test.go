package role_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	librole "github.com/muhlba91/pulumi-shared-library/pkg/lib/aws/iam/role"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateRole(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		resourceName := "res-basic"
		nameInput := pulumi.String("ci-role")
		desc := pulumi.String("CI role for tests")
		policy := pulumi.String(`{"Version":"2012-10-17","Statement":[]}`)
		labels := map[string]string{"team": "ci"}

		r, err := librole.Create(ctx, resourceName, &librole.CreateOptions{
			Name:             nameInput,
			Description:      desc,
			AssumeRolePolicy: policy,
			Labels:           labels,
		})
		require.NoError(t, err)
		require.NotNil(t, r)

		r.Name.ApplyT(func(n string) error {
			assert.Equal(t, "ci-role", n)
			return nil
		})
		r.Description.ApplyT(func(d *string) error {
			assert.Equal(t, "CI role for tests", *d)
			return nil
		})
		r.AssumeRolePolicy.ApplyT(func(p string) error {
			assert.JSONEq(t, `{"Version":"2012-10-17","Statement":[]}`, p)
			return nil
		})
		r.Tags.ApplyT(func(m map[string]string) error {
			assert.Equal(t, labels, m)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateRole_WithOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		resourceName := "res-protected"
		nameInput := pulumi.String("ci-role-protected")
		desc := pulumi.String("Protected CI role")
		policy := pulumi.String(
			`{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":"s3:*","Resource":"*"}]}`,
		)
		labels := map[string]string{"env": "prod"}

		r, err := librole.Create(ctx, resourceName, &librole.CreateOptions{
			Name:             nameInput,
			Description:      desc,
			AssumeRolePolicy: policy,
			Labels:           labels,
			PulumiOptions: []pulumi.ResourceOption{
				pulumi.Protect(true),
			},
		})
		require.NoError(t, err)
		require.NotNil(t, r)

		r.Name.ApplyT(func(n string) error {
			assert.Equal(t, "ci-role-protected", n)
			return nil
		})
		r.Description.ApplyT(func(d *string) error {
			assert.Equal(t, "Protected CI role", *d)
			return nil
		})
		r.AssumeRolePolicy.ApplyT(func(p string) error {
			assert.JSONEq(
				t,
				`{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":"s3:*","Resource":"*"}]}`,
				p,
			)
			return nil
		})
		r.Tags.ApplyT(func(m map[string]string) error {
			assert.Equal(t, labels, m)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
