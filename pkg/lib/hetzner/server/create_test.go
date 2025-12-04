package server_test

import (
	"testing"

	"github.com/pulumi/pulumi-hcloud/sdk/go/hcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libprimaryip "github.com/muhlba91/pulumi-shared-library/pkg/lib/hetzner/primaryip"
	libserver "github.com/muhlba91/pulumi-shared-library/pkg/lib/hetzner/server"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateServer(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		// create primary ips to link to the server
		p4, err := libprimaryip.Create(ctx, "primaryip", &libprimaryip.CreateOptions{
			Name:       "p4",
			IPType:     "ipv4",
			Datacenter: "fsn1-dc14",
			AutoDelete: pulumi.Bool(true),
			Labels:     map[string]string{"env": "test"},
		})
		require.NoError(t, err)
		require.NotNil(t, p4)

		p6, err := libprimaryip.Create(ctx, "primaryip", &libprimaryip.CreateOptions{
			Name:       "p6",
			IPType:     "ipv6",
			Datacenter: "nbg1-dc3",
			AutoDelete: pulumi.Bool(true),
			Labels:     map[string]string{"env": "test"},
		})
		require.NoError(t, err)
		require.NotNil(t, p6)

		opts := &libserver.CreateOptions{
			Hostname:   pulumi.String("my-host"),
			ServerType: pulumi.String("cx11"),
			Image:      pulumi.String("debian-11"),
			SSHKeys: []pulumi.StringInput{
				pulumi.String("ssh-key-1"),
				pulumi.String("ssh-key-2"),
			},
			Location:           pulumi.String("fsn1"),
			NetworkID:          pulumi.Int(42),
			IPAddress:          pulumi.String("10.0.1.5"),
			PrimaryIPv4Address: p4,
			PrimaryIPv6Address: p6,
			Firewalls: []pulumi.IntInput{
				pulumi.Int(1),
				pulumi.Int(2),
			},
			Backups:       pulumi.Bool(false),
			Protection:    false,
			PublicSSH:     false,
			Labels:        map[string]string{"role": "app"},
			PulumiOptions: nil,
		}

		srv, err := libserver.Create(ctx, "srv-basic", opts)
		require.NoError(t, err)
		require.NotNil(t, srv)

		srv.Hostname.ApplyT(func(h string) error {
			assert.Equal(t, "my-host", h)
			return nil
		})
		srv.PrivateIPv4.ApplyT(func(ip string) error {
			assert.Equal(t, "10.0.1.5", ip)
			return nil
		})
		srv.PublicIPv4.ApplyT(func(ip string) error {
			assert.Equal(t, "mocked-ip-address-ipv4-fsn1-dc14", ip)
			return nil
		})
		srv.PublicIPv6.ApplyT(func(ip string) error {
			assert.Equal(t, "mocked-ip-address-ipv6-nbg1-dc3", ip)
			return nil
		})
		srv.SSHIPv4.ApplyT(func(ssh string) error {
			assert.Equal(t, "10.0.1.5", ssh)
			return nil
		})
		srv.Resource.(*hcloud.Server).PublicNets.ApplyT(func(pns []hcloud.ServerPublicNet) error {
			assert.Len(t, pns, 1)
			assert.True(t, *pns[0].Ipv4Enabled)
			assert.True(t, *pns[0].Ipv6Enabled)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateServer_PublicSSH(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		p4, err := libprimaryip.Create(ctx, "primaryip", &libprimaryip.CreateOptions{
			Name:       "p4b",
			IPType:     "ipv4",
			Datacenter: "fsn1-dc14",
			AutoDelete: pulumi.Bool(true),
			Labels:     map[string]string{"env": "prod"},
		})
		require.NoError(t, err)
		require.NotNil(t, p4)

		p6, err := libprimaryip.Create(ctx, "primaryip", &libprimaryip.CreateOptions{
			Name:       "p6b",
			IPType:     "ipv6",
			Datacenter: "nbg1-dc3",
			AutoDelete: pulumi.Bool(false),
			Labels:     map[string]string{"env": "prod"},
		})
		require.NoError(t, err)
		require.NotNil(t, p6)

		enableIPv6 := false
		opts := &libserver.CreateOptions{
			Hostname:           pulumi.String("public-host"),
			ServerType:         pulumi.String("cx21"),
			Image:              pulumi.String("ubuntu-20.04"),
			SSHKeys:            nil,
			Location:           pulumi.String("nbg1"),
			NetworkID:          pulumi.Int(100),
			IPAddress:          pulumi.String("10.1.1.10"),
			PrimaryIPv4Address: p4,
			PrimaryIPv6Address: p6,
			EnableIPv6:         &enableIPv6,
			Firewalls:          nil,
			Backups:            pulumi.Bool(false),
			Protection:         true,
			PublicSSH:          true,
			Labels:             map[string]string{"role": "db"},
			PulumiOptions:      []pulumi.ResourceOption{pulumi.Protect(true)},
		}

		srv, err := libserver.Create(ctx, "srv-public", opts)
		require.NoError(t, err)
		require.NotNil(t, srv)

		srv.SSHIPv4.ApplyT(func(ssh string) error {
			assert.Equal(t, "mocked-ip-address-ipv4-fsn1-dc14", ssh)
			return nil
		})
		srv.Resource.(*hcloud.Server).PublicNets.ApplyT(func(pns []hcloud.ServerPublicNet) error {
			assert.Len(t, pns, 1)
			assert.True(t, *pns[0].Ipv4Enabled)
			assert.False(t, *pns[0].Ipv6Enabled)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
