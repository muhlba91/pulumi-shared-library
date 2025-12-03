package iam

import (
	"fmt"

	gcpkms "github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

// KeyringBindingArgs represents the arguments for creating a KeyRing IAM binding.
type KeyringBindingArgs struct {
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
// args: KeyringBindingArgs containing KeyRingID, Member, Role, and optional Pulumi options.
func CreateKeyringBinding(
	ctx *pulumi.Context,
	args *KeyringBindingArgs,
) (*gcpkms.KeyRingIAMBinding, error) {
	name := fmt.Sprintf(
		"gcp-kms-iam-binding-%s-%s-%s",
		sanitize.Text(args.KeyRingID),
		sanitize.Text(args.Member),
		sanitize.Text(args.Role),
	)

	return gcpkms.NewKeyRingIAMBinding(ctx, name, &gcpkms.KeyRingIAMBindingArgs{
		KeyRingId: pulumi.String(args.KeyRingID),
		Role:      pulumi.String(args.Role),
		Members:   pulumi.StringArray{pulumi.String(args.Member)},
	}, args.PulumiOptions...)
}
