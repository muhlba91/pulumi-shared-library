package policy

import (
	"fmt"

	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateOptions are the options for creating a Vault policy.
type CreateOptions struct {
	// Name is the name of the policy.
	Name string
	// Policy is the policy document for the vault.
	Policy pulumi.StringInput
	// PulumiOptions are optional resource options (e.g. provider).
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new Vault policy with the specified name and policy document.
// ctx: Pulumi context
// opts: CreateOptions containing the name and policy document.
func Create(
	ctx *pulumi.Context,
	opts *CreateOptions,
) (*vault.Policy, error) {
	return vault.NewPolicy(ctx,
		fmt.Sprintf("vault-policy-%s", opts.Name),
		&vault.PolicyArgs{
			Name:   pulumi.String(opts.Name),
			Policy: opts.Policy,
		},
		opts.PulumiOptions...)
}
