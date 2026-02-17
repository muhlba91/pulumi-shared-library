package serviceaccount_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	utiliam "github.com/muhlba91/pulumi-shared-library/pkg/util/google/iam/serviceaccount"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateServiceAccountUser_NoRoles(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "sa-basic"
		project := "proj-basic"
		roles := []string{}

		opts := &utiliam.CreateOptions{
			Name:    name,
			Project: pulumi.String(project),
			Roles:   roles,
		}

		data, err := utiliam.CreateServiceAccountUser(ctx, opts)
		require.NoError(err)
		require.NotNil(data)
		require.NotNil(data.ServiceAccount)
		require.NotNil(data.Key)

		data.Key.ServiceAccountId.ApplyT(func(kid string) error {
			data.ServiceAccount.Name.ApplyT(func(saName string) error {
				assert.Equal(saName, kid)
				return nil
			})
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(err)
}

func TestCreateServiceAccountUser_WithRoles(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "sa-with-roles"
		project := "proj-roles"
		roles := []string{"roles/viewer", "roles/storage.objectAdmin"}

		opts := &utiliam.CreateOptions{
			Name:    name,
			Project: pulumi.String(project),
			Roles:   roles,
		}

		data, err := utiliam.CreateServiceAccountUser(ctx, opts)
		require.NoError(err)
		require.NotNil(data)
		require.NotNil(data.ServiceAccount)
		require.NotNil(data.Key)

		data.Key.ServiceAccountId.ApplyT(func(kid string) error {
			data.ServiceAccount.Name.ApplyT(func(saName string) error {
				assert.Equal(saName, kid)
				return nil
			})
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(err)
}
