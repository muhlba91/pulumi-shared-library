package user_test

import (
	"fmt"
	"testing"

	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libuser "github.com/muhlba91/pulumi-shared-library/pkg/lib/aws/iam/user"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateUser(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "alice"
		labels := map[string]string{"env": "test"}

		u, err := libuser.Create(ctx, name, &libuser.CreateOptions{
			Labels: labels,
		})
		require.NoError(t, err)
		require.NotNil(t, u)

		u.Name.ApplyT(func(n string) error {
			assert.Equal(t, fmt.Sprintf("aws-user-%s", name), n)
			return nil
		})
		u.Tags.ApplyT(func(m map[string]string) error {
			assert.Equal(t, labels, m)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateUser_WithPolicies(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "bob"

		p1, err := iam.NewPolicy(ctx, "policy1", &iam.PolicyArgs{
			Policy: pulumi.String(`{"Version":"2012-10-17","Statement":[]}`),
		})
		require.NoError(t, err)
		p2, err := iam.NewPolicy(ctx, "policy2", &iam.PolicyArgs{
			Policy: pulumi.String(`{"Version":"2012-10-17","Statement":[]}`),
		})
		require.NoError(t, err)

		u, err := libuser.Create(ctx, name, &libuser.CreateOptions{
			Policies: []*iam.Policy{p1, p2},
		})
		require.NoError(t, err)
		require.NotNil(t, u)

		u.Name.ApplyT(func(n string) error {
			assert.Equal(t, fmt.Sprintf("aws-user-%s", name), n)
			return nil
		})

		p1.Arn.ApplyT(func(a string) error {
			assert.NotEmpty(t, a)
			return nil
		})
		p2.Arn.ApplyT(func(a string) error {
			assert.NotEmpty(t, a)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
