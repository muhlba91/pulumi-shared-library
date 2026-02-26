package serviceaccount_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/google/iam/serviceaccount"
	"github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateKey(t *testing.T) {
	counter := mocks.NewCounter()

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
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Nil(t, counter.Resources["time:index/rotating:Rotating"])
	assert.Len(t, counter.Resources["gcp:serviceaccount/key:Key"], 1)
}

func TestCreateKey_WithOptionalArgs(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		serviceAccount := "my-service-account"

		opts := &serviceaccount.KeyOptions{
			ServiceAccount: pulumi.String(serviceAccount),
			Rotation:       &rotation.Options{},
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
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Len(t, counter.Resources["time:index/rotating:Rotating"], 1)
	assert.Len(t, counter.Resources["gcp:serviceaccount/key:Key"], 1)
}
