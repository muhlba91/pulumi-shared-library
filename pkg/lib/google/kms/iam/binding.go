package iam

import (
	"fmt"

	gcpkms "github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

// KeyringBindingOptions represents the options for creating a KeyRing IAM binding.
type KeyringBindingOptions struct {
	// KeyRingID is the ID of the KeyRing to attach the IAM binding to.
	KeyRingID string
	// Member is the member to assign the role to (e.g., "user:<email>").
	Member string
	// Role is the role to assign to the IAM binding.
	Role string
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// CreateKeyringBinding defines a new IAM binding for a KeyRing.
// ctx: Pulumi context.
// opts: KeyringBindingOptions containing KeyRingID, Member, Role, and optional Pulumi options.
func CreateKeyringBinding(
	ctx *pulumi.Context,
	opts *KeyringBindingOptions,
) (*gcpkms.KeyRingIAMBinding, error) {
	name := fmt.Sprintf(
		"gcp-kms-iam-binding-%s-%s-%s",
		sanitize.Text(opts.KeyRingID),
		sanitize.Text(opts.Member),
		sanitize.Text(opts.Role),
	)

	return gcpkms.NewKeyRingIAMBinding(ctx, name, &gcpkms.KeyRingIAMBindingArgs{
		KeyRingId: pulumi.String(opts.KeyRingID),
		Role:      pulumi.String(opts.Role),
		Members:   pulumi.StringArray{pulumi.String(opts.Member)},
	}, opts.PulumiOptions...)
}
