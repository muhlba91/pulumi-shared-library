package rotation_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-time/sdk/go/time"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/rotation"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreate_DefaultDays(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "rot-default"

		res, err := rotation.Create(ctx, name, 0)
		require.NoError(t, err)
		require.NotNil(t, res)

		rt, ok := res.(*time.Rotating)
		require.True(t, ok)

		rt.RotationDays.ApplyT(func(d *int) error {
			assert.Equal(t, 30, *d)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreate_CustomDays(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "rot-custom"
		days := 7

		res, err := rotation.Create(ctx, name, days)
		require.NoError(t, err)
		require.NotNil(t, res)

		rt, ok := res.(*time.Rotating)
		require.True(t, ok)

		rt.RotationDays.ApplyT(func(d *int) error {
			assert.Equal(t, days, *d)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
