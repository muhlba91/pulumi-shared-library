package secret

import (
	"fmt"

	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/kv"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// WriteArgs are the arguments for the Write function.
type WriteArgs struct {
	// Key is the secret key name.
	Key string
	// Value is the value to store, as JSON string input.
	Value pulumi.StringInput
	// Path is the KV mount path (e.g. "secret").
	Path string
	// PulumiOptions are optional resource options (e.g. provider).
	PulumiOptions []pulumi.ResourceOption
}

// Write stores a value in Vault KV v2 at the given mount/path with the given key.
// ctx: Pulumi context
// args: WriteArgs containing Key, Value, Path, and optional Pulumi options.
func Write(
	ctx *pulumi.Context,
	args *WriteArgs,
) (*kv.SecretV2, error) {
	name := fmt.Sprintf("vault-secret-%s-%s", args.Path, args.Key)

	return kv.NewSecretV2(ctx, name, &kv.SecretV2Args{
		Mount:    pulumi.String(args.Path),
		Name:     pulumi.String(args.Key),
		DataJson: args.Value,
	}, args.PulumiOptions...)
}
