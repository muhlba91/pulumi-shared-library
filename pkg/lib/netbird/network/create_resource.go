package network

import (
	"fmt"

	"github.com/KitStream/netbird-pulumi-provider/sdk/go/netbird"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
)

// defaultEnabled is the default enabled setting for a network resource (true).
const defaultEnabled = true

// CreateResourceOptions defines the options for creating a NetBird network resource.
type CreateResourceOptions struct {
	// Address is the address of the resource: a host (e.g., "1.1.1.1"), a subnet (e.g., "192.168.178.0/24"),
	// or a domain (e.g., "example.com" or "*.example.com"). Required.
	Address string
	// Groups are the IDs of the groups containing the resource. Required.
	Groups pulumi.StringArrayInput
	// NetworkID is the ID of the network the resource belongs to. Required.
	NetworkID pulumi.StringInput
	// Description is the description of the network resource. Optional.
	Description *string
	// Enabled defines whether the network resource is active. Defaults to true.
	Enabled *bool
	// PulumiOptions are additional options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// CreateResource creates a new NetBird network resource.
// ctx: The Pulumi context.
// name: The logical name for the Pulumi resource (prefixed with "netbird-network-resource-"),
// also used as the network resource name.
// opts: The options for creating the network resource. Must not be nil.
func CreateResource(ctx *pulumi.Context, name string, opts *CreateResourceOptions) (*netbird.NetworkResource, error) {
	resName := fmt.Sprintf("netbird-network-resource-%s", name)

	args := &netbird.NetworkResourceArgs{
		Name:      pulumi.String(name),
		Address:   pulumi.String(opts.Address),
		Groups:    opts.Groups,
		NetworkId: opts.NetworkID,
		Enabled:   pulumi.Bool(defaults.GetOrDefault(opts.Enabled, defaultEnabled)),
	}
	if opts.Description != nil {
		args.Description = pulumi.String(*opts.Description)
	}

	return netbird.NewNetworkResource(ctx, resName, args, opts.PulumiOptions...)
}
