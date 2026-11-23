package role_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	librole "github.com/muhlba91/pulumi-shared-library/pkg/lib/aws/iam/role"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateRolePolicyAttachment(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		pa, err := librole.CreateRolePolicyAttachment(ctx, "basic", &librole.CreateRolePolicyAttachmentOptions{
			Role:      pulumi.String("role-a"),
			PolicyArn: pulumi.String("arn:aws:iam::123456789012:policy/test-policy"),
		})
		require.NoError(t, err)
		require.NotNil(t, pa)

		pa.Role.ApplyT(func(r string) error {
			assert.Equal(t, "role-a", r)
			return nil
		})
		pa.PolicyArn.ApplyT(func(a string) error {
			assert.Equal(t, "arn:aws:iam::123456789012:policy/test-policy", a)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}
