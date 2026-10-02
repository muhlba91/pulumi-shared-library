package subnet

import (
	"fmt"

	"github.com/pulumi/pulumi-hcloud/sdk/go/hcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateOptions defines the options for creating a Hetzner subnet.
type CreateOptions struct {
	// NetworkID is the ID of the network the subnet belongs to.
	NetworkID pulumi.IntInput
	// Cidr is the CIDR block for the subnet.
	Cidr string
	// PulumiOptions are additional Pulumi resource options.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a Hetzner subnet of the type cloud in the network zone eu-central.
// ctx: The Pulumi context.
// name: The logical name for the Pulumi resource (prefixed with "hcloud-subnet-").
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
