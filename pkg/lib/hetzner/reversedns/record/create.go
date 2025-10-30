package record

import (
	"fmt"

	"github.com/pulumi/pulumi-hcloud/sdk/go/hcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/pulumi/convert"
)

// CreateOptions defines the options for creating a Hetzner Reverse DNS Record.
type CreateOptions struct {
	// DNSName is the DNS name to set for the RDNS record.
	DNSName string
	// PrimaryIP is the primary IP to associate with the RDNS record.
	PrimaryIP *hcloud.PrimaryIp
	// IPType is the type of IP address (e.g., "ipv4" or "ipv6").
	IPType string
	// Datacenter is the datacenter where the IP address is located.
	Datacenter string
	// PulumiOptions are the options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a Hetzner Reverse DNS Record with the given options.
// ctx: The Pulumi context.
// opts: The options for creating the RDNS record.
func Create(ctx *pulumi.Context, opts *CreateOptions) (*hcloud.Rdns, error) {
	return hcloud.NewRdns(
		ctx,
		fmt.Sprintf("hcloud-rdns-%s-%s-%s", opts.IPType, opts.Datacenter, opts.DNSName),
		&hcloud.RdnsArgs{
			PrimaryIpId: convert.IDToInt(opts.PrimaryIP.ID()),
			IpAddress:   opts.PrimaryIP.IpAddress,
			DnsPtr:      pulumi.String(opts.DNSName),
		},
		opts.PulumiOptions...)
}
