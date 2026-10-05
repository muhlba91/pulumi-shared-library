package network

import (
	"fmt"

	"github.com/KitStream/netbird-pulumi-provider/sdk/go/netbird"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateOptions defines the options for creating a NetBird network.
type CreateOptions struct {
	// Description is the description of the network. Optional.
	Description *string
	// PulumiOptions are additional options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new NetBird network.
// ctx: The Pulumi context.
// name: The logical name for the Pulumi resource (prefixed with "netbird-network-"), also used as the network name.
// opts: The options for creating the network. Must not be nil.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*netbird.Network, error) {
	resName := fmt.Sprintf("netbird-network-%s", name)

	args := &netbird.NetworkArgs{
		Name: pulumi.String(name),
	}
	if opts.Description != nil {
		args.Description = pulumi.String(*opts.Description)
	}

	return netbird.NewNetwork(ctx, resName, args, opts.PulumiOptions...)
}
