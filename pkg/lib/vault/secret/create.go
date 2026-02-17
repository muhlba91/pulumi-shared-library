package secret

import (
	"fmt"

	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/kv"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateOptions are the options for creating a Vault secret.
type CreateOptions struct {
	// Key is the secret key name.
	Key string
	// Value is the value to store, as JSON string input.
	Value pulumi.StringInput
	// Path is the KV mount path (e.g. "secret").
	Path string
	// PulumiOptions are optional resource options (e.g. provider).
	PulumiOptions []pulumi.ResourceOption
}

// Create stores a value in Vault KV v2 at the given mount/path with the given key.
// ctx: Pulumi context.
// opts: CreateOptions containing Key, Value, Path, and optional Pulumi options.
func Create(
	ctx *pulumi.Context,
	opts *CreateOptions,
) (*kv.SecretV2, error) {
	pulumiOpts := opts.PulumiOptions
	if pulumiOpts == nil {
		pulumiOpts = []pulumi.ResourceOption{}
	}
	pulumiOpts = append(pulumiOpts, pulumi.DeleteBeforeReplace(true))

	name := fmt.Sprintf("vault-secret-%s-%s", opts.Path, opts.Key)

	return kv.NewSecretV2(ctx, name, &kv.SecretV2Args{
		Mount:    pulumi.String(opts.Path),
		Name:     pulumi.String(opts.Key),
		DataJson: opts.Value,
	}, pulumiOpts...)
}
