package policy_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/iam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libpolicy "github.com/muhlba91/pulumi-shared-library/pkg/lib/scaleway/iam/policy"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreatePolicy(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "ci-policy"
		desc := pulumi.String("CI role policy")
		rules := []iam.PolicyRuleInput{
			iam.PolicyRuleArgs{},
		}
		labels := []string{"team:ci"}

		res, err := libpolicy.Create(ctx, name, &libpolicy.CreateOptions{
			Name:        pulumi.String(name),
			Description: desc,
			Rules:       rules,
			Labels:      labels,
			UserID:      pulumi.StringPtr(name),
		})
		require.NoError(t, err)
		require.NotNil(t, res)

		res.Name.ApplyT(func(n string) error {
			assert.Equal(t, name, n)
			return nil
		})
		res.UserId.ApplyT(func(n *string) error {
			assert.Equal(t, name, *n)
			return nil
		})
		res.ApplicationId.ApplyT(func(n *string) error {
			assert.Nil(t, n)
			return nil
		})
		res.Rules.ApplyT(func(p []iam.PolicyRule) error {
			assert.Len(t, p, 1)
			return nil
		})
		res.Tags.ApplyT(func(m []string) error {
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
		rules := []iam.PolicyRuleInput{
			iam.PolicyRuleArgs{},
		}
		labels := []string{"env:prod"}

		res, err := libpolicy.Create(ctx, name, &libpolicy.CreateOptions{
			Name:          pulumi.String(name),
			ApplicationID: pulumi.StringPtr(name),
			Description:   desc,
			Rules:         rules,
			Labels:        labels,
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
		res.UserId.ApplyT(func(n *string) error {
			assert.Nil(t, n)
			return nil
		})
		res.ApplicationId.ApplyT(func(n *string) error {
			assert.Equal(t, name, *n)
			return nil
		})
		res.Rules.ApplyT(func(p []iam.PolicyRule) error {
			assert.Len(t, p, 1)
			return nil
		})
		res.Tags.ApplyT(func(m []string) error {
			assert.Equal(t, labels, m)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
