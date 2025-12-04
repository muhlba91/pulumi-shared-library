package server

import (
	"fmt"

	"github.com/pulumi/pulumi-hcloud/sdk/go/hcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	hModel "github.com/muhlba91/pulumi-shared-library/pkg/model/hetzner"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/pulumi/convert"
)

// CreateOptions are the options for creating a Hetzner server.
type CreateOptions struct {
	// Hostname is the hostname of the server.
	Hostname pulumi.StringInput
	// ServerType is the type of the server.
	ServerType pulumi.StringInput
	// Image is the image of the server.
	Image pulumi.StringInput
	// SSHKeys are the SSH keys to add to the server.
	SSHKeys []pulumi.StringInput
	// Location is the location of the server.
	Location pulumi.StringInput
	// NetworkID is the ID of the network to attach the server to.
	NetworkID pulumi.IntInput
	// IPAddress is the IP address to assign to the server.
	IPAddress pulumi.StringInput
	// PrimaryIPv4Address is the primary IPv4 address for the server's public network.
	PrimaryIPv4Address *hcloud.PrimaryIp
	// PrimaryIPv6Address is the primary IPv6 address for the server's public network.
	PrimaryIPv6Address *hcloud.PrimaryIp
	// EnableIPv6 indicates whether IPv6 should be enabled for the server.
	EnableIPv6 *bool
	// Firewalls are the firewalls to add to the server.
	Firewalls []pulumi.IntInput
	// Backups indicates whether backups should be enabled for the server.
	Backups pulumi.BoolInput
	// Protection indicates whether delete and rebuild protection should be enabled for the server.
	Protection bool
	// PublicSSH indicates whether the server should have a public SSH access.
	PublicSSH bool
	// Labels are the labels to assign to the server.
	Labels map[string]string
	// PulumiOptions are the options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a Hetzner server.
// ctx: Pulumi context
// name: the name of the server
// opts: CreateOptions for the server
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*hModel.Server, error) {
	optsWithProtection := append([]pulumi.ResourceOption{}, opts.PulumiOptions...)
	optsWithProtection = append(optsWithProtection, pulumi.Protect(opts.Protection))

	sshKeys := pulumi.StringArray{}
	for _, key := range opts.SSHKeys {
		sshKeys = append(sshKeys, key)
	}

	firewalls := pulumi.IntArray{}
	for _, fw := range opts.Firewalls {
		firewalls = append(firewalls, fw)
	}

	s, err := hcloud.NewServer(ctx, fmt.Sprintf("hcloud-server-%s", name), &hcloud.ServerArgs{
		Name:       opts.Hostname,
		ServerType: opts.ServerType,
		Image:      opts.Image,
		SshKeys:    sshKeys,
		Location:   opts.Location,
		Networks: &hcloud.ServerNetworkTypeArray{
			hcloud.ServerNetworkTypeArgs{
				NetworkId: opts.NetworkID,
				Ip:        opts.IPAddress,
			},
		},
		PublicNets: &hcloud.ServerPublicNetArray{
			hcloud.ServerPublicNetArgs{
				Ipv4Enabled: pulumi.Bool(true),
				Ipv4:        convert.IDToInt(opts.PrimaryIPv4Address.ID()),
				Ipv6Enabled: pulumi.Bool(defaults.GetOrDefault(opts.EnableIPv6, true)),
				Ipv6:        convert.IDToInt(opts.PrimaryIPv6Address.ID()),
			},
		},
		FirewallIds:       firewalls,
		Backups:           opts.Backups,
		DeleteProtection:  pulumi.Bool(opts.Protection),
		RebuildProtection: pulumi.Bool(opts.Protection),
		KeepDisk:          pulumi.Bool(opts.Protection),
		Labels:            pulumi.ToStringMap(opts.Labels),
	}, optsWithProtection...)
	if err != nil {
		return nil, err
	}

	ssh := opts.IPAddress.ToStringOutput()
	if opts.PublicSSH {
		ssh = opts.PrimaryIPv4Address.IpAddress
	}
	return &hModel.Server{
		Resource:    s,
		Hostname:    s.Name,
		PrivateIPv4: opts.IPAddress.ToStringOutput(),
		PublicIPv4:  opts.PrimaryIPv4Address.IpAddress,
		PublicIPv6:  opts.PrimaryIPv6Address.IpAddress,
		SSHIPv4:     ssh,
	}, nil
}
