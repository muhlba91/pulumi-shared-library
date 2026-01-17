package primaryip

import (
	"fmt"

	"github.com/pulumi/pulumi-hcloud/sdk/go/hcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateOptions defines the options for creating a Hetzner Primary IP.
type CreateOptions struct {
	// Name is the base name for the Primary IP.
	Name string
	// IPType is the type of the Primary IP.
	IPType string
	// Datacenter is the datacenter where the Primary IP will be created.
	Datacenter string
	// AutoDelete indicates whether the Primary IP should be automatically deleted when the assignee is deleted.
	AutoDelete pulumi.BoolInput
	// Labels are the labels to assign to the Primary IP.
	Labels map[string]string
	// PulumiOptions are the options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a Hetzner Primary IP with the given options.
// ctx: Pulumi context for resource creation.
// opts: Options for creating the Primary IP.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*hcloud.PrimaryIp, error) {
	return hcloud.NewPrimaryIp(
		ctx,
		fmt.Sprintf("hcloud-primary-ip-%s-%s-%s", name, opts.IPType, opts.Datacenter),
		&hcloud.PrimaryIpArgs{
			Name:         pulumi.String(fmt.Sprintf("%s-%s-%s", opts.Name, opts.IPType, opts.Datacenter)),
			AssigneeType: pulumi.String("server"),
			Type:         pulumi.String(opts.IPType),
			Location:     pulumi.String(opts.Datacenter),
			AutoDelete:   opts.AutoDelete,
			Labels:       pulumi.ToStringMap(opts.Labels),
		},
		opts.PulumiOptions...)
}
