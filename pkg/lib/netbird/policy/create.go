package policy

import (
	"fmt"

	"github.com/KitStream/netbird-pulumi-provider/sdk/go/netbird"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	pModel "github.com/muhlba91/pulumi-shared-library/pkg/model/netbird/policy"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
)

const (
	// defaultEnabled is the default enabled setting for a policy and its rules (true).
	defaultEnabled = true
	// defaultAction is the default action for a rule.
	defaultAction = "accept"
	// defaultBidirectional is the default bidirectional setting for a rule (true).
	defaultBidirectional = true
	// defaultProtocol is the default protocol for a rule.
	defaultProtocol = "all"
	// protocolTCP and protocolUDP are the only protocols that support limiting a rule to ports.
	protocolTCP = "tcp"
	protocolUDP = "udp"
	// maxRules is the maximum number of rules per policy supported by the NetBird provider.
	maxRules = 1
)

// CreateOptions defines the options for creating a NetBird policy.
type CreateOptions struct {
	// Description is the description of the policy. Optional.
	Description *string
	// Enabled defines whether the policy is active. Defaults to true.
	Enabled *bool
	// Rules are the rules of the policy. Optional.
	// The NetBird provider currently supports a single rule per policy.
	Rules []pModel.Rule
	// PulumiOptions are additional options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new NetBird policy.
// Returns an error if more rules are given than the NetBird provider supports.
// ctx: The Pulumi context.
// name: The logical name for the Pulumi resource (prefixed with "netbird-policy-"), also used as the policy name.
// opts: The options for creating the policy. Must not be nil.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*netbird.Policy, error) {
	resName := fmt.Sprintf("netbird-policy-%s", name)

	if len(opts.Rules) > maxRules {
		return nil, fmt.Errorf("a policy supports at most %d rule(s), got %d", maxRules, len(opts.Rules))
	}

	args := &netbird.PolicyArgs{
		Name:    pulumi.String(name),
		Enabled: pulumi.Bool(defaults.GetOrDefault(opts.Enabled, defaultEnabled)),
	}
	if opts.Description != nil {
		args.Description = pulumi.String(*opts.Description)
	}
	if len(opts.Rules) > 0 {
		rule, err := toRuleArgs(opts.Rules[0])
		if err != nil {
			return nil, err
		}
		args.Rule = rule
	}

	return netbird.NewPolicy(ctx, resName, args, opts.PulumiOptions...)
}

// toRuleArgs converts a rule to the input format of the NetBird provider, applying defaults.
// Returns an error if ports are set for a protocol other than "tcp" or "udp".
// rule: The rule to convert.
func toRuleArgs(rule pModel.Rule) (netbird.PolicyRulePtrInput, error) {
	protocol := defaults.GetOrDefault(rule.Protocol, defaultProtocol)
	if len(rule.Ports) > 0 && protocol != protocolTCP && protocol != protocolUDP {
		return nil, fmt.Errorf("rule %q: ports require protocol %q or %q, got %q",
			rule.Name, protocolTCP, protocolUDP, protocol)
	}

	args := netbird.PolicyRuleArgs{
		Name:          pulumi.String(rule.Name),
		Action:        pulumi.String(defaults.GetOrDefault(rule.Action, defaultAction)),
		Enabled:       pulumi.Bool(defaults.GetOrDefault(rule.Enabled, defaultEnabled)),
		Bidirectional: pulumi.Bool(defaults.GetOrDefault(rule.Bidirectional, defaultBidirectional)),
		Protocol:      pulumi.String(protocol),
		Sources:       rule.Sources,
		Destinations:  rule.Destinations,
	}
	if len(rule.Ports) > 0 {
		args.Ports = pulumi.ToStringArray(rule.Ports)
	}
	return &args, nil
}
