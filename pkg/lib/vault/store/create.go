package store

import (
	"fmt"

	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateOptions are the options for creating a Vault store.
type CreateOptions struct {
	// Path is the key path to store the value at.
	Path pulumi.StringInput
	// Description is the description of the vault store.
	Description pulumi.StringInput
	// NamePrefix is an optional prefix for the resource name.
	NamePrefix *string
	// PulumiOptions are optional resource options (e.g. provider).
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new Vault KV v2 store at the specified path.
// ctx: Pulumi context
// name: Name of the vault store resource
// opts: CreateOptions containing the path and description.
func Create(
	ctx *pulumi.Context,
	name string,
	opts *CreateOptions,
) (*vault.Mount, error) {
	prefix := "store"
	if opts.NamePrefix != nil {
		prefix = *opts.NamePrefix
	}

	return vault.NewMount(ctx, fmt.Sprintf("vault-%s-%s", prefix, name), &vault.MountArgs{
		Path: opts.Path,
		Type: pulumi.String("kv"),
		Options: pulumi.ToStringMap(map[string]string{
			"version": "2",
		}),
		Description: opts.Description,
	}, opts.PulumiOptions...)
}
