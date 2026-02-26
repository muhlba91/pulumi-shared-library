package rotation_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	rModel "github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/rotation"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestTrigger_Default(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &rModel.Options{}
		res, err := rotation.Trigger(ctx, "rotation", opts)
		require.NoError(t, err)
		require.NotNil(t, res)
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Len(t, counter.Resources["time:index/rotating:Rotating"], 1)
	assert.Contains(t, counter.Resources["time:index/rotating:Rotating"], "rotation-rotation")
}

func TestTrigger_WithOptions(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "new-name"

		opts := &rModel.Options{
			Name: &name,
		}
		res, err := rotation.Trigger(ctx, "rotation", opts)
		require.NoError(t, err)
		require.NotNil(t, res)
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Len(t, counter.Resources["time:index/rotating:Rotating"], 1)
	assert.Contains(t, counter.Resources["time:index/rotating:Rotating"], "rotation-new-name")
}

func TestTrigger_Nil(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		res, err := rotation.Trigger(ctx, "rotation", nil)
		require.NoError(t, err)
		require.Nil(t, res)
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Nil(t, counter.Resources["time:index/rotating:Rotating"])
}
