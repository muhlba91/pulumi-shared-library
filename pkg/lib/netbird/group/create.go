package group

import (
	"fmt"

	"github.com/KitStream/netbird-pulumi-provider/sdk/go/netbird"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateOptions defines the options for creating a NetBird group.
type CreateOptions struct {
	// PulumiOptions are additional options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new NetBird group.
// ctx: The Pulumi context.
// name: The logical name for the Pulumi resource (prefixed with "netbird-group-"), also used as the group name.
// opts: The options for creating the group. Must not be nil.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*netbird.Group, error) {
	resName := fmt.Sprintf("netbird-group-%s", name)

	return netbird.NewGroup(ctx, resName, &netbird.GroupArgs{
		Name: pulumi.String(name),
	}, opts.PulumiOptions...)
}
