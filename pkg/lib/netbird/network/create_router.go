package network

import (
	"errors"
	"fmt"

	"github.com/KitStream/netbird-pulumi-provider/sdk/go/netbird"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
)

const (
	// defaultMetric is the default route metric for a network router (9999, the lowest priority).
	defaultMetric = 9999
	// defaultMasquerade is the default masquerade setting for a network router (false).
	defaultMasquerade = false
)

// CreateRouterOptions defines the options for creating a NetBird network router.
type CreateRouterOptions struct {
	// NetworkID is the ID of the network the router belongs to. Required.
	NetworkID pulumi.StringInput
	// Peer is the ID of the peer acting as router. Optional.
	// Must not be set together with PeerGroups.
	Peer pulumi.StringInput
	// PeerGroups are the IDs of the peer groups acting as routers. Optional.
	// Must not be set together with Peer.
	PeerGroups []pulumi.StringInput
	// Metric is the route metric. The lowest number has the highest priority. Defaults to 9999.
	Metric *int
	// Masquerade defines whether peers masquerade traffic to the route's prefix. Defaults to false.
	Masquerade *bool
	// PulumiOptions are additional options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// CreateRouter creates a new NetBird network router.
// Returns an error if both a peer and peer groups are given.
// ctx: The Pulumi context.
// name: The logical name for the Pulumi resource (prefixed with "netbird-network-router-").
// opts: The options for creating the network router. Must not be nil.
func CreateRouter(ctx *pulumi.Context, name string, opts *CreateRouterOptions) (*netbird.NetworkRouter, error) {
	resName := fmt.Sprintf("netbird-network-router-%s", name)

	if opts.Peer != nil && len(opts.PeerGroups) > 0 {
		return nil, errors.New("a network router supports either a peer or peer groups, not both")
	}

	args := &netbird.NetworkRouterArgs{
		NetworkId:  opts.NetworkID,
		Metric:     pulumi.Int(defaults.GetOrDefault(opts.Metric, defaultMetric)),
		Masquerade: pulumi.Bool(defaults.GetOrDefault(opts.Masquerade, defaultMasquerade)),
	}
	if opts.Peer != nil {
		args.Peer = opts.Peer.ToStringOutput().ToStringPtrOutput()
	}
	if len(opts.PeerGroups) > 0 {
		peerGroups := make(pulumi.StringArray, 0, len(opts.PeerGroups))
		for _, g := range opts.PeerGroups {
			peerGroups = append(peerGroups, g)
		}
		args.PeerGroups = peerGroups
	}

	return netbird.NewNetworkRouter(ctx, resName, args, opts.PulumiOptions...)
}
