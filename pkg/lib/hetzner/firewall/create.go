package firewall

import (
	"fmt"

	"github.com/pulumi/pulumi-hcloud/sdk/go/hcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Rule represents a Hetzner firewall rule.
type Rule struct {
	// Direction is the direction of the rule (in or out).
	Direction string
	// Protocol is the protocol of the rule (e.g., tcp, udp).
	Protocol string
	// Port is the port or port range of the rule.
	Port string
	// SourceIps are the source IPs or CIDR blocks for the rule.
	SourceIps []pulumi.StringInput
}

// CreateOptions represents the options for creating a Hetzner firewall.
type CreateOptions struct {
	// Name is the name of the firewall.
	Name string
	// Rules are the firewall rules to apply.
	Rules []Rule
	// Labels are the labels to assign to the firewall.
	Labels map[string]string
	// PulumiOptions are the options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a Hetzner firewall with the given options.
// ctx: Pulumi context.
// name: The name of the firewall.
// opts: The options for creating the firewall.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*hcloud.Firewall, error) {
	rules := hcloud.FirewallRuleArray{}
	for _, r := range opts.Rules {
		var src pulumi.StringArray
		if len(r.SourceIps) == 0 {
			src = pulumi.StringArray{
				pulumi.String("0.0.0.0/0"),
				pulumi.String("::/0"),
			}
		} else {
			src = pulumi.StringArray(r.SourceIps)
		}

		rules = append(rules, hcloud.FirewallRuleArgs{
			Direction: pulumi.String(r.Direction),
			Protocol:  pulumi.String(r.Protocol),
			Port:      pulumi.String(r.Port),
			SourceIps: src,
		})
	}

	return hcloud.NewFirewall(ctx, fmt.Sprintf("hcloud-firewall-%s", name), &hcloud.FirewallArgs{
		Name:   pulumi.String(opts.Name),
		Rules:  rules,
		Labels: pulumi.ToStringMap(opts.Labels),
	}, opts.PulumiOptions...)
}
