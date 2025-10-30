package hetzner

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

// Server represents a Hetzner Server with commonly used outputs.
type Server struct {
	// Resource is the underlying Pulumi resource.
	Resource pulumi.Resource
	// Hostname is the server's hostname.
	Hostname pulumi.StringOutput
	// PrivateIPv4 is the server's private IPv4 address.
	PrivateIPv4 pulumi.StringOutput
	// PublicIPv4 is the server's public IPv4 address.
	PublicIPv4 pulumi.StringOutput
	// PublicIPv6 is the server's public IPv6 address.
	PublicIPv6 pulumi.StringOutput
	// SSHIPv4 is the server's SSH IPv4 address.
	SSHIPv4 pulumi.StringOutput
}
