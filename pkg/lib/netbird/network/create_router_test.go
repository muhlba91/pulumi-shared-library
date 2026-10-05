package network_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/netbird/network"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

const routerTok = "netbird:index/networkRouter:NetworkRouter"

func TestCreateRouter_Defaults(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		r, err := network.CreateRouter(ctx, "test", &network.CreateRouterOptions{
			NetworkID:  pulumi.String("net1"),
			PeerGroups: []pulumi.StringInput{pulumi.String("g1"), pulumi.String("g2")},
		})
		require.NoError(t, err)
		require.NotNil(t, r)

		r.NetworkId.ApplyT(func(v string) error { assert.Equal(t, "net1", v); return nil })
		r.PeerGroups.ApplyT(func(v []string) error { assert.Equal(t, []string{"g1", "g2"}, v); return nil })
		r.Metric.ApplyT(func(v int) error { assert.Equal(t, 9999, v); return nil })
		r.Masquerade.ApplyT(func(v bool) error { assert.False(t, v); return nil })
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Equal(t, []string{"netbird-network-router-test"}, counter.Resources[routerTok])
}

func TestCreateRouter_CustomOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		r, err := network.CreateRouter(ctx, "custom", &network.CreateRouterOptions{
			NetworkID:  pulumi.String("net1"),
			PeerGroups: []pulumi.StringInput{pulumi.String("g1")},
			Metric:     ptr(100),
			Masquerade: ptr(true),
		})
		require.NoError(t, err)

		r.Metric.ApplyT(func(v int) error { assert.Equal(t, 100, v); return nil })
		r.Masquerade.ApplyT(func(v bool) error { assert.True(t, v); return nil })
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}

func TestCreateRouter_Peer(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		r, err := network.CreateRouter(ctx, "peer", &network.CreateRouterOptions{
			NetworkID: pulumi.String("net1"),
			Peer:      pulumi.String("p1"),
		})
		require.NoError(t, err)

		r.Peer.ApplyT(func(v *string) error { assert.Equal(t, "p1", *v); return nil })
		r.PeerGroups.ApplyT(func(v []string) error { assert.Empty(t, v); return nil })
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}

func TestCreateRouter_PeerAndPeerGroups(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		r, err := network.CreateRouter(ctx, "both", &network.CreateRouterOptions{
			NetworkID:  pulumi.String("net1"),
			Peer:       pulumi.String("p1"),
			PeerGroups: []pulumi.StringInput{pulumi.String("g1")},
		})
		require.Error(t, err)
		assert.Nil(t, r)
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}
