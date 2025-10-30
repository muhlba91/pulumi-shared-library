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
			Name:       "myp",
			IPType:     "ipv4",
			Datacenter: "fsn1-dc14",
			AutoDelete: pulumi.Bool(true),
			Labels:     map[string]string{"env": "test"},
		}

		res, err := libip.Create(ctx, opts)
		require.NoError(t, err)
		require.NotNil(t, res)

		// Name should contain the provided base name
		res.Name.ApplyT(func(n string) error {
			assert.Contains(t, n, opts.Name)
			return nil
		})
		res.Type.ApplyT(func(typ string) error {
			assert.Equal(t, opts.IPType, typ)
			return nil
		})
		res.Datacenter.ApplyT(func(d string) error {
			assert.Equal(t, opts.Datacenter, d)
			return nil
		})
		res.AutoDelete.ApplyT(func(b bool) error {
			assert.True(t, b)
			return nil
		})
		res.IpAddress.ApplyT(func(a string) error {
			assert.Equal(t, "mocked-ip-address-ipv4-fsn1-dc14", a)
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
		opts := &libip.CreateOptions{
			Name:       "protected",
			IPType:     "ipv6",
			Datacenter: "nbg1-dc3",
			AutoDelete: pulumi.Bool(false),
			Labels:     map[string]string{"team": "dev"},
			PulumiOptions: []pulumi.ResourceOption{
				pulumi.Protect(true),
			},
		}

		res, err := libip.Create(ctx, opts)
		require.NoError(t, err)
		require.NotNil(t, res)

		res.Type.ApplyT(func(typ string) error {
			assert.Equal(t, opts.IPType, typ)
			return nil
		})
		res.Datacenter.ApplyT(func(d string) error {
			assert.Equal(t, opts.Datacenter, d)
			return nil
		})
		res.AutoDelete.ApplyT(func(b bool) error {
			assert.False(t, b)
			return nil
		})
		res.IpAddress.ApplyT(func(a string) error {
			assert.Equal(t, "mocked-ip-address-ipv6-nbg1-dc3", a)
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
