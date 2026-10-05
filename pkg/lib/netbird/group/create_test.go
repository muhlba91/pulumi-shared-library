package group_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/netbird/group"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreate(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		g, err := group.Create(ctx, "test", &group.CreateOptions{})
		require.NoError(t, err)
		require.NotNil(t, g)

		g.Name.ApplyT(func(v string) error { assert.Equal(t, "test", v); return nil })
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Equal(t, []string{"netbird-group-test"}, counter.Resources["netbird:index/group:Group"])
}
