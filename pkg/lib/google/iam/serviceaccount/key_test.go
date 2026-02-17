package serviceaccount_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/google/iam/serviceaccount"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateKey(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		serviceAccount := "my-service-account"

		opts := &serviceaccount.KeyOptions{
			ServiceAccount: pulumi.String(serviceAccount),
		}

		key, err := serviceaccount.CreateKey(ctx, "test", opts)
		require.NoError(t, err)
		require.NotNil(t, key)

		key.ServiceAccountId.ApplyT(func(id string) error {
			assert.Equal(t, serviceAccount, id)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateKey_WithOptionalArgs(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		serviceAccount := "my-service-account"

		opts := &serviceaccount.KeyOptions{
			ServiceAccount: pulumi.String(serviceAccount),
			PulumiOptions:  []pulumi.ResourceOption{},
		}

		key, err := serviceaccount.CreateKey(ctx, "test", opts)
		require.NoError(t, err)
		require.NotNil(t, key)

		key.ServiceAccountId.ApplyT(func(id string) error {
			assert.Equal(t, serviceAccount, id)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
