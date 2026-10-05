package peer

import (
	"github.com/KitStream/netbird-pulumi-provider/sdk/go/netbird"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GetOptions defines the options for looking up a NetBird peer.
// At least one of ID, IP, or Name must be set.
type GetOptions struct {
	// ID is the ID of the peer. Optional.
	ID *string
	// IP is the NetBird IP address of the peer. Optional.
	IP *string
	// Name is the name of the peer. Optional.
	Name *string
	// PulumiOptions are additional options to pass to the Pulumi invoke.
	PulumiOptions []pulumi.InvokeOption
}

// Get looks up an existing NetBird peer.
// ctx: The Pulumi context.
// opts: The options for looking up the peer. Must not be nil.
func Get(ctx *pulumi.Context, opts *GetOptions) (*netbird.LookupPeerResult, error) {
	return netbird.LookupPeer(ctx, &netbird.LookupPeerArgs{
		Id:   opts.ID,
		Ip:   opts.IP,
		Name: opts.Name,
	}, opts.PulumiOptions...)
}
