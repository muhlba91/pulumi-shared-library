package network_test

import (
	"testing"

	"github.com/pulumi/pulumi-hcloud/sdk/go/hcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libnet "github.com/muhlba91/pulumi-shared-library/pkg/lib/hetzner/network"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateNetwork(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &libnet.CreateOptions{
			Name: "net-basic",
			Cidr: pulumi.String("10.0.0.0/16"),
			Labels: map[string]string{
				"env": "test",
			},
		}

		res, err := libnet.Create(ctx, opts)
		require.NoError(t, err)
		require.NotNil(t, res)

		res.Name.ApplyT(func(n string) error {
			assert.Equal(t, "net-basic", n)
			return nil
		})
		res.IpRange.ApplyT(func(cidr string) error {
			assert.Equal(t, "10.0.0.0/16", cidr)
			return nil
		})
		res.ExposeRoutesToVswitch.ApplyT(func(b *bool) error {
			assert.True(t, *b)
			return nil
		})
		res.Labels.ApplyT(func(m map[string]string) error {
			assert.Equal(t, map[string]string{"env": "test"}, m)
			return nil
		})
		// also assert underlying type shape via Rules of SDK if needed (no-op here)
		_ = hcloud.Network{}
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateNetwork_WithOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &libnet.CreateOptions{
			Name: "net-protected",
			Cidr: pulumi.String("192.168.0.0/24"),
			Labels: map[string]string{
				"team": "dev",
			},
			PulumiOptions: []pulumi.ResourceOption{
				pulumi.Protect(true),
			},
		}

		res, err := libnet.Create(ctx, opts)
		require.NoError(t, err)
		require.NotNil(t, res)

		res.Name.ApplyT(func(n string) error {
			assert.Equal(t, "net-protected", n)
			return nil
		})
		res.IpRange.ApplyT(func(cidr string) error {
			assert.Equal(t, "192.168.0.0/24", cidr)
			return nil
		})
		res.ExposeRoutesToVswitch.ApplyT(func(b *bool) error {
			assert.True(t, *b)
			return nil
		})
		res.Labels.ApplyT(func(m map[string]string) error {
			assert.Equal(t, map[string]string{"team": "dev"}, m)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
