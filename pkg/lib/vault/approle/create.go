package approle

import (
	"fmt"

	"github.com/pulumi/pulumi-vault/sdk/v7/go/vault/approle"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
)

const (
	// defaultBackend is the default mount path of the AppRole auth method.
	defaultBackend = "approle"
	// defaultTokenTTL is the default TTL of issued tokens in seconds (1 hour).
	defaultTokenTTL = 60 * 60
	// defaultTokenMaxTTL is the default maximum TTL of issued tokens in seconds (24 hours).
	defaultTokenMaxTTL = 24 * 60 * 60
	// defaultSecretIDTTL is the default TTL of secret IDs in seconds (0 means they do not expire).
	defaultSecretIDTTL = 0
)

// CreateOptions defines the options for creating a Vault AppRole role.
type CreateOptions struct {
	// Backend is the mount path of the AppRole auth method. Defaults to "approle".
	Backend *string
	// TokenPolicies are the policies attached to tokens issued by this role. Optional.
	TokenPolicies pulumi.StringArrayInput
	// TokenTTL is the TTL of issued tokens in seconds. Defaults to 1 hour.
	TokenTTL *int
	// TokenMaxTTL is the maximum TTL of issued tokens in seconds. Defaults to 24 hours.
	TokenMaxTTL *int
	// SecretIDTTL is the TTL of secret IDs in seconds. Defaults to 0 (secret IDs do not expire).
	SecretIDTTL *int
	// SecretIDBoundCIDRs restricts the source addresses that may log in with a secret ID. Optional.
	SecretIDBoundCIDRs pulumi.StringArrayInput
	// TokenBoundCIDRs restricts the source addresses that may use issued tokens. Optional.
	TokenBoundCIDRs pulumi.StringArrayInput
	// PulumiOptions are additional options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new Vault AppRole role.
// ctx: The Pulumi context.
// name: The logical name for the Pulumi resource (prefixed with "vault-approle-"), also used as the role name.
// opts: The options for creating the role. Must not be nil.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*approle.AuthBackendRole, error) {
	return approle.NewAuthBackendRole(ctx, fmt.Sprintf("vault-approle-%s", name), &approle.AuthBackendRoleArgs{
		Backend:            pulumi.String(defaults.GetOrDefault(opts.Backend, defaultBackend)),
		RoleName:           pulumi.String(name),
		TokenPolicies:      opts.TokenPolicies,
		TokenTtl:           pulumi.Int(defaults.GetOrDefault(opts.TokenTTL, defaultTokenTTL)),
		TokenMaxTtl:        pulumi.Int(defaults.GetOrDefault(opts.TokenMaxTTL, defaultTokenMaxTTL)),
		SecretIdTtl:        pulumi.Int(defaults.GetOrDefault(opts.SecretIDTTL, defaultSecretIDTTL)),
		SecretIdBoundCidrs: opts.SecretIDBoundCIDRs,
		TokenBoundCidrs:    opts.TokenBoundCIDRs,
	}, opts.PulumiOptions...)
}
