package rotation_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/rotation"
	rModel "github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreate_DefaultDays(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "rot-default"

		opts := &rModel.Options{
			Name: &name,
		}
		res, err := rotation.Create(ctx, opts)
		require.NoError(t, err)
		require.NotNil(t, res)

		res.RotationDays.ApplyT(func(d *int) error {
			assert.Equal(t, 30, *d)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}

func TestCreate_CustomDays(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "rot-custom"
		days := 7

		opts := &rModel.Options{
			Name: &name,
			Days: days,
		}
		res, err := rotation.Create(ctx, opts)
		require.NoError(t, err)
		require.NotNil(t, res)

		res.RotationDays.ApplyT(func(d *int) error {
			assert.Equal(t, days, *d)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}
