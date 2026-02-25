package primaryip_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libip "github.com/muhlba91/pulumi-shared-library/pkg/lib/hetzner/primaryip"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreatePrimaryIP(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &libip.CreateOptions{
			Name:       "myip",
			IPType:     "ipv4",
			Location:   "fsn1",
			AutoDelete: pulumi.Bool(true),
			Labels:     map[string]string{"env": "test"},
		}

		res, err := libip.Create(ctx, "primaryip", opts)
		require.NoError(t, err)
		require.NotNil(t, res)

		res.Name.ApplyT(func(n string) error {
			assert.Equal(t, "myip-ipv4-fsn1", n)
			return nil
		})
		res.Type.ApplyT(func(typ string) error {
			assert.Equal(t, opts.IPType, typ)
			return nil
		})
		res.Location.ApplyT(func(d string) error {
			assert.Equal(t, opts.Location, d)
			return nil
		})
		res.AutoDelete.ApplyT(func(b bool) error {
			assert.True(t, b)
			return nil
		})
		res.IpAddress.ApplyT(func(a string) error {
			assert.Equal(t, "mocked-ip-address-ipv4-fsn1", a)
			return nil
		})
		res.Labels.ApplyT(func(m map[string]string) error {
			assert.Equal(t, opts.Labels, m)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreatePrimaryIP_WithOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		dc := "fsn1-dc14"
		opts := &libip.CreateOptions{
			Name:       "protected",
			IPType:     "ipv6",
			Location:   "fsn1",
			Datacenter: &dc,
			AutoDelete: pulumi.Bool(false),
			Labels:     map[string]string{"team": "dev"},
			PulumiOptions: []pulumi.ResourceOption{
				pulumi.Protect(true),
			},
		}

		res, err := libip.Create(ctx, "primaryip", opts)
		require.NoError(t, err)
		require.NotNil(t, res)

		res.Name.ApplyT(func(n string) error {
			assert.Equal(t, "protected-ipv6-fsn1", n)
			return nil
		})
		res.Type.ApplyT(func(typ string) error {
			assert.Equal(t, opts.IPType, typ)
			return nil
		})
		res.Location.ApplyT(func(d string) error {
			assert.Equal(t, opts.Location, d)
			return nil
		})
		res.AutoDelete.ApplyT(func(b bool) error {
			assert.False(t, b)
			return nil
		})
		res.IpAddress.ApplyT(func(a string) error {
			assert.Equal(t, "mocked-ip-address-ipv6-fsn1", a)
			return nil
		})
		res.Labels.ApplyT(func(m map[string]string) error {
			assert.Equal(t, opts.Labels, m)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
