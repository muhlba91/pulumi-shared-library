package user_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	utiliam "github.com/muhlba91/pulumi-shared-library/pkg/util/scaleway/iam/user"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateServiceAccountUser_NoRoles(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "sa-basic"
		project := "proj-basic"
		labels := []string{}

		args := &utiliam.CreateUserArgs{
			Name:             name,
			Email:            pulumi.String(name),
			DefaultProjectID: pulumi.String(project),
			Labels:           labels,
		}

		data, err := utiliam.CreateUser(ctx, args)
		require.NoError(err)
		require.NotNil(data)
		require.NotNil(data.User)
		require.NotNil(data.Key)

		data.Key.UserId.ApplyT(func(kid *string) error {
			data.User.ID().ToStringOutput().ApplyT(func(saName string) error {
				assert.Equal(saName, *kid)
				return nil
			})
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(err)
}
