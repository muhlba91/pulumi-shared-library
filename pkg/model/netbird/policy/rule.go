package policy

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

// Rule defines a NetBird policy rule.
type Rule struct {
	// Name is the name of the rule. Required.
	Name string
	// Action is the effect of the rule ("accept" or "drop"). Defaults to "accept".
	Action *string
	// Enabled defines whether the rule is active. Defaults to true.
	Enabled *bool
	// Bidirectional defines whether the rule applies in both directions. Defaults to true.
	Bidirectional *bool
	// Protocol is the protocol of the rule ("tcp", "udp", "icmp", "all", or "netbird-ssh"). Defaults to "all".
	// The value is not validated by the library, but by the NetBird provider.
	Protocol *string
	// Ports are the individual destination ports (e.g., "443") the rule is limited to. Optional.
	// If empty, all ports are allowed. Only allowed if the protocol is "tcp" or "udp".
	Ports []string
	// Sources are the IDs of the source groups. Optional.
	Sources pulumi.StringArrayInput
	// Destinations are the IDs of the destination groups. Optional.
	Destinations pulumi.StringArrayInput
}
