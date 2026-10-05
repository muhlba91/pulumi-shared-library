package network_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/netbird/network"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

const resourceTok = "netbird:index/networkResource:NetworkResource"

func TestCreateResource_Defaults(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		r, err := network.CreateResource(ctx, "test", &network.CreateResourceOptions{
			Address:   "10.0.0.0/24",
			Groups:    pulumi.StringArray{pulumi.String("g1")},
			NetworkID: pulumi.String("net1"),
		})
		require.NoError(t, err)
		require.NotNil(t, r)

		r.Name.ApplyT(func(v string) error { assert.Equal(t, "test", v); return nil })
		r.Address.ApplyT(func(v string) error { assert.Equal(t, "10.0.0.0/24", v); return nil })
		r.NetworkId.ApplyT(func(v string) error { assert.Equal(t, "net1", v); return nil })
		r.Groups.ApplyT(func(v []string) error { assert.Equal(t, []string{"g1"}, v); return nil })
		r.Enabled.ApplyT(func(v bool) error { assert.True(t, v); return nil })
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Equal(t, []string{"netbird-network-resource-test"}, counter.Resources[resourceTok])
}

func TestCreateResource_CustomOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		r, err := network.CreateResource(ctx, "custom", &network.CreateResourceOptions{
			Address:     "example.com",
			Groups:      pulumi.StringArray{pulumi.String("g1")},
			NetworkID:   pulumi.String("net1"),
			Description: ptr("desc"),
			Enabled:     ptr(false),
		})
		require.NoError(t, err)

		r.Description.ApplyT(func(v string) error { assert.Equal(t, "desc", v); return nil })
		r.Enabled.ApplyT(func(v bool) error { assert.False(t, v); return nil })
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}
