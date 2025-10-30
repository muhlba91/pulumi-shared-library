package sshkey

import (
	"fmt"

	"github.com/pulumi/pulumi-hcloud/sdk/go/hcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateOptions defines the options for creating a Hetzner Cloud SSH key.
type CreateOptions struct {
	// Name is the base name for the SSH key.
	Name string
	// PublicKey is the public key content.
	PublicKey pulumi.StringInput
	// Labels are the labels to assign to the SSH key.
	Labels map[string]string
	// PulumiOptions are the options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a Hetzner Cloud SSH key with the given name, public key, and labels.
// ctx: The Pulumi context.
// opts: The options for creating the SSH key.
func Create(ctx *pulumi.Context, opts *CreateOptions) (*hcloud.SshKey, error) {
	return hcloud.NewSshKey(ctx, fmt.Sprintf("hcloud-ssh-%s", opts.Name), &hcloud.SshKeyArgs{
		Name:      pulumi.String(opts.Name),
		PublicKey: opts.PublicKey,
		Labels:    pulumi.ToStringMap(opts.Labels),
	}, opts.PulumiOptions...)
}
