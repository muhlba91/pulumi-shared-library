package network

import (
	"fmt"

	"github.com/pulumi/pulumi-hcloud/sdk/go/hcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CreateOptions struct {
	// Name is the name of the network.
	Name string
	// Cidr is the CIDR block for the network.
	Cidr pulumi.StringInput
	// Labels are the labels to apply to the network.
	Labels map[string]string
	// PulumiOptions are additional Pulumi resource options.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a Hetzner network.
// ctx: The Pulumi context.
// opts: The options for creating the network.
func Create(ctx *pulumi.Context, opts *CreateOptions) (*hcloud.Network, error) {
	return hcloud.NewNetwork(ctx, fmt.Sprintf("hcloud-network-%s", opts.Name), &hcloud.NetworkArgs{
		Name:                  pulumi.String(opts.Name),
		IpRange:               opts.Cidr,
		ExposeRoutesToVswitch: pulumi.Bool(true),
		Labels:                pulumi.ToStringMap(opts.Labels),
	}, opts.PulumiOptions...)
}
