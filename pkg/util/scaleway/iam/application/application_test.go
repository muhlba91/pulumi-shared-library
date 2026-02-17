package application_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	utiliam "github.com/muhlba91/pulumi-shared-library/pkg/util/scaleway/iam/application"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateApplication(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "sa-basic"
		project := "proj-basic"
		labels := []string{}

		opts := &utiliam.CreateOptions{
			Name:             name,
			Description:      pulumi.StringPtr(name),
			DefaultProjectID: pulumi.String(project),
			Labels:           labels,
		}

		data, err := utiliam.CreateApplication(ctx, opts)
		require.NoError(err)
		require.NotNil(data)
		require.NotNil(data.Application)
		require.NotNil(data.Key)

		data.Key.ApplicationId.ApplyT(func(kid *string) error {
			assert.NotEmpty(kid)
			data.Application.ID().ToStringOutput().ApplyT(func(saName string) error {
				assert.Equal(saName, *kid)
				return nil
			})
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(err)
}
