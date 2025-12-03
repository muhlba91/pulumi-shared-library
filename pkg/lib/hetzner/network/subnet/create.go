package subnet

import (
	"fmt"

	"github.com/pulumi/pulumi-hcloud/sdk/go/hcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CreateOptions struct {
	// NetworkID is the ID of the subnet.
	NetworkID pulumi.IntInput
	// Cidr is the CIDR block for the subnet.
	Cidr string
	// PulumiOptions are additional Pulumi resource options.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a Hetzner subnet.
// ctx: The Pulumi context.
// name: The name of the subnet.
// opts: The options for creating the subnet.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*hcloud.NetworkSubnet, error) {
	return hcloud.NewNetworkSubnet(
		ctx,
		fmt.Sprintf("hcloud-subnet-%s", name),
		&hcloud.NetworkSubnetArgs{
			NetworkId:   opts.NetworkID,
			Type:        pulumi.String("cloud"),
			NetworkZone: pulumi.String("eu-central"),
			IpRange:     pulumi.String(opts.Cidr),
		},
		opts.PulumiOptions...)
}
