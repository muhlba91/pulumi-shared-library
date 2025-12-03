package subnet_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libsubnet "github.com/muhlba91/pulumi-shared-library/pkg/lib/hetzner/network/subnet"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateSubnet(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &libsubnet.CreateOptions{
			NetworkID:     pulumi.Int(42),
			Cidr:          "10.0.1.0/24",
			PulumiOptions: nil,
		}

		res, err := libsubnet.Create(ctx, "subnet", opts)
		require.NoError(t, err)
		require.NotNil(t, res)

		res.IpRange.ApplyT(func(ipr string) error {
			assert.Equal(t, "10.0.1.0/24", ipr)
			return nil
		})
		res.NetworkId.ApplyT(func(id int) error {
			assert.NotZero(t, id)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateSubnet_WithOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &libsubnet.CreateOptions{
			NetworkID: pulumi.Int(100),
			Cidr:      "10.1.2.0/24",
			PulumiOptions: []pulumi.ResourceOption{
				pulumi.Protect(true),
			},
		}

		res, err := libsubnet.Create(ctx, "subnet", opts)
		require.NoError(t, err)
		require.NotNil(t, res)

		res.IpRange.ApplyT(func(ipr string) error {
			assert.Equal(t, "10.1.2.0/24", ipr)
			return nil
		})
		res.NetworkId.ApplyT(func(id int) error {
			assert.NotZero(t, id)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
