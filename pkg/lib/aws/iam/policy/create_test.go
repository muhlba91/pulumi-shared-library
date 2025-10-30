package policy_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libpolicy "github.com/muhlba91/pulumi-shared-library/pkg/lib/aws/iam/policy"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreatePolicy(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "ci-policy"
		desc := pulumi.String("CI role policy")
		policyJSON := pulumi.String(`{"Version":"2012-10-17","Statement":[]}`)
		labels := map[string]string{"team": "ci"}

		res, err := libpolicy.Create(ctx, name, &libpolicy.CreateOptions{
			Name:        pulumi.String(name),
			Description: desc,
			Policy:      policyJSON,
			Labels:      labels,
		})
		require.NoError(t, err)
		require.NotNil(t, res)

		res.Name.ApplyT(func(n string) error {
			assert.Equal(t, name, n)
			return nil
		})
		res.Policy.ApplyT(func(p string) error {
			assert.JSONEq(t, `{"Version":"2012-10-17","Statement":[]}`, p)
			return nil
		})
		res.Tags.ApplyT(func(m map[string]string) error {
			assert.Equal(t, labels, m)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreatePolicy_WithOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "ci-policy-protected"
		desc := pulumi.String("Protected policy")
		policyJSON := pulumi.String(
			`{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":"s3:*","Resource":"*"}]}`,
		)
		labels := map[string]string{"env": "prod"}

		res, err := libpolicy.Create(ctx, name, &libpolicy.CreateOptions{
			Name:        pulumi.String(name),
			Description: desc,
			Policy:      policyJSON,
			Labels:      labels,
			PulumiOptions: []pulumi.ResourceOption{
				pulumi.Protect(true),
			},
		})
		require.NoError(t, err)
		require.NotNil(t, res)

		res.Name.ApplyT(func(n string) error {
			assert.Equal(t, name, n)
			return nil
		})
		res.Policy.ApplyT(func(p string) error {
			assert.JSONEq(
				t,
				`{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":"s3:*","Resource":"*"}]}`,
				p,
			)
			return nil
		})
		res.Tags.ApplyT(func(m map[string]string) error {
			assert.Equal(t, labels, m)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
