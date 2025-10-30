package record_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libip "github.com/muhlba91/pulumi-shared-library/pkg/lib/hetzner/primaryip"
	librdns "github.com/muhlba91/pulumi-shared-library/pkg/lib/hetzner/reversedns/record"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateRdns(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		primaryOpts := &libip.CreateOptions{
			Name:       "myp",
			IPType:     "ipv4",
			Datacenter: "fsn1-dc14",
			AutoDelete: pulumi.Bool(true),
			Labels:     map[string]string{"env": "test"},
		}
		primary, err := libip.Create(ctx, primaryOpts)
		require.NoError(t, err)
		require.NotNil(t, primary)

		rdnsOpts := &librdns.CreateOptions{
			DNSName:    "example.com",
			PrimaryIP:  primary,
			IPType:     "ipv4",
			Datacenter: "fsn1-dc14",
		}
		rdns, err := librdns.Create(ctx, rdnsOpts)
		require.NoError(t, err)
		require.NotNil(t, rdns)

		rdns.IpAddress.ApplyT(func(ip string) error {
			assert.Equal(t, "mocked-ip-address-ipv4-fsn1-dc14", ip)
			return nil
		})
		rdns.DnsPtr.ApplyT(func(d string) error {
			assert.Equal(t, "example.com", d)
			return nil
		})
		rdns.PrimaryIpId.ApplyT(func(id *int) error {
			assert.NotZero(t, *id)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateRdns_WithOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		primaryOpts := &libip.CreateOptions{
			Name:       "protected",
			IPType:     "ipv6",
			Datacenter: "nbg1-dc3",
			AutoDelete: pulumi.Bool(false),
			Labels:     map[string]string{"team": "dev"},
			PulumiOptions: []pulumi.ResourceOption{
				pulumi.Protect(true),
			},
		}
		primary, err := libip.Create(ctx, primaryOpts)
		require.NoError(t, err)
		require.NotNil(t, primary)

		rdnsOpts := &librdns.CreateOptions{
			DNSName:    "srv.example.net",
			PrimaryIP:  primary,
			IPType:     "ipv6",
			Datacenter: "nbg1-dc3",
			PulumiOptions: []pulumi.ResourceOption{
				pulumi.Protect(true),
			},
		}
		rdns, err := librdns.Create(ctx, rdnsOpts)
		require.NoError(t, err)
		require.NotNil(t, rdns)

		rdns.IpAddress.ApplyT(func(ip string) error {
			assert.Equal(t, "mocked-ip-address-ipv6-nbg1-dc3", ip)
			return nil
		})
		rdns.DnsPtr.ApplyT(func(d string) error {
			assert.Equal(t, "srv.example.net", d)
			return nil
		})
		rdns.PrimaryIpId.ApplyT(func(id *int) error {
			assert.NotZero(t, *id)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
