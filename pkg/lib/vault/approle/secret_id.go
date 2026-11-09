package approle

import (
	"fmt"

	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/approle"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	rModel "github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/rotation"
)

// SecretIDOptions defines the options for creating a secret ID of a Vault AppRole role.
type SecretIDOptions struct {
	// RoleName is the name of the AppRole role, which may be managed elsewhere. Required.
	RoleName pulumi.StringInput
	// Backend is the mount path of the AppRole auth method. Defaults to "approle".
	Backend *string
	// Metadata is a JSON object of metadata attached to tokens issued with this secret ID. Optional.
	Metadata pulumi.StringPtrInput
	// Rotation defines the rotation options for the resource. Optional.
	Rotation *rModel.Options
	// PulumiOptions are additional options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// CreateSecretID creates a new secret ID for a Vault AppRole role.
// The role ID is available from approle.GetAuthBackendRoleId of the Vault provider.
// ctx: The Pulumi context.
// name: The logical name for the Pulumi resource (prefixed with "vault-approle-secret-id-").
// opts: The options for creating the secret ID. Must not be nil.
func CreateSecretID(
	ctx *pulumi.Context,
	name string,
	opts *SecretIDOptions,
) (*approle.AuthBackendRoleSecretId, error) {
	resName := fmt.Sprintf("vault-approle-secret-id-%s", name)

	pulumiOpts := append([]pulumi.ResourceOption{}, opts.PulumiOptions...)
	if trigger, _ := rotation.Trigger(ctx, resName, opts.Rotation); trigger != nil {
		pulumiOpts = append(pulumiOpts, pulumi.ReplacementTrigger(trigger))
	}

	return approle.NewAuthBackendRoleSecretId(ctx, resName, &approle.AuthBackendRoleSecretIdArgs{
		Backend:  pulumi.String(defaults.GetOrDefault(opts.Backend, defaultBackend)),
		RoleName: opts.RoleName,
		Metadata: opts.Metadata,
	}, pulumiOpts...)
}
