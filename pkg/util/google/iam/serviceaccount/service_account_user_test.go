package serviceaccount_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	rModel "github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
	utiliam "github.com/muhlba91/pulumi-shared-library/pkg/util/google/iam/serviceaccount"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateServiceAccountUser_NoRoles(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &utiliam.CreateOptions{
			Name:    "sa-basic",
			Project: pulumi.String("proj-basic"),
			Roles:   []string{},
		}

		data, err := utiliam.CreateServiceAccountUser(ctx, opts)
		require.NoError(err)
		require.NotNil(data)
		require.NotNil(data.ServiceAccount)
		require.NotNil(data.Key)

		pulumi.All(data.Key.ServiceAccountId, data.ServiceAccount.Name).ApplyT(func(all []any) error {
			assert.Equal(all[1], all[0])
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(err)

	assert.Equal([]string{"gcp-sa-sa-basic"}, counter.Resources["gcp:serviceaccount/account:Account"])
	assert.Equal([]string{"gcp-sa-key-sa-basic"}, counter.Resources["gcp:serviceaccount/key:Key"])
	assert.Empty(counter.Resources["time:index/rotating:Rotating"])
}

func TestCreateServiceAccountUser_WithRolesAndOptionalArgs(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	counter := mocks.NewCounter()
	roles := []string{"roles/viewer", "roles/storage.objectAdmin"}

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		displayName := "Service Account"
		description := "a description"
		ignoreExisting := true

		opts := &utiliam.CreateOptions{
			Name:                      "sa-with-roles",
			DisplayName:               &displayName,
			Description:               &description,
			Project:                   pulumi.String("proj-roles"),
			Roles:                     roles,
			Rotation:                  &rModel.Options{Days: 365},
			CreateIgnoreAlreadyExists: &ignoreExisting,
			PulumiOptions:             []pulumi.ResourceOption{},
		}

		data, err := utiliam.CreateServiceAccountUser(ctx, opts)
		require.NoError(err)
		require.NotNil(data)
		require.NotNil(data.ServiceAccount)
		require.NotNil(data.Key)

		pulumi.All(data.ServiceAccount.DisplayName, data.ServiceAccount.Description).ApplyT(func(all []any) error {
			assert.Equal(displayName, *all[0].(*string))
			assert.Equal(description, *all[1].(*string))
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(err)

	assert.Len(counter.Resources["gcp:serviceaccount/account:Account"], 1)
	assert.Len(counter.Resources["gcp:projects/iAMMember:IAMMember"], len(roles))
	assert.Len(counter.Resources["gcp:serviceaccount/key:Key"], 1)
	assert.Len(counter.Resources["time:index/rotating:Rotating"], 1)
}
