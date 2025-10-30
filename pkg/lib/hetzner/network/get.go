package network

import (
	"github.com/pulumi/pulumi-hcloud/sdk/go/hcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get gets or creates a Hetzner network.
// ctx: The Pulumi context.
// name: The name of the network.
func Get(ctx *pulumi.Context, name string) (*hcloud.LookupNetworkResult, error) {
	return hcloud.LookupNetwork(ctx, &hcloud.LookupNetworkArgs{
		Name: &name,
	})
}
