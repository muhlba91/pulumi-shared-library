package iam

import (
	"fmt"

	gcpkms "github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

// MemberOptions represents the options for creating a CryptoKey IAM member.
type MemberOptions struct {
	// CryptoKeyID is the ID of the CryptoKey to attach the IAM member to.
	CryptoKeyID string
	// Member is the member to assign the role to (e.g., "user:<email>").
	Member string
	// Role is the role to assign to the IAM member.
	Role string
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// KeyringMemberOptions represents the options for creating a KeyRing IAM member.
type KeyringMemberOptions struct {
	// KeyRingID is the ID of the KeyRing to attach the IAM member to.
	KeyRingID string
	// Member is the member to assign the role to (e.g., "user:<email>").
	Member string
	// Role is the role to assign to the IAM member.
	Role string
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// CreateMember defines a new IAM member for a CryptoKey.
// ctx: Pulumi context.
// opts: MemberOptions containing CryptoKeyID, Member, Role, and optional Pulumi options.
func CreateMember(
	ctx *pulumi.Context,
	opts *MemberOptions,
) (*gcpkms.CryptoKeyIAMMember, error) {
	name := fmt.Sprintf(
		"gcp-kms-iam-cryptokey-member-%s-%s-%s",
		sanitize.Text(opts.CryptoKeyID),
		sanitize.Text(opts.Member),
		sanitize.Text(opts.Role),
	)

	return gcpkms.NewCryptoKeyIAMMember(ctx, name, &gcpkms.CryptoKeyIAMMemberArgs{
		CryptoKeyId: pulumi.String(opts.CryptoKeyID),
		Role:        pulumi.String(opts.Role),
		Member:      pulumi.String(opts.Member),
	}, opts.PulumiOptions...)
}

// CreateKeyringMember defines a new IAM member for a KeyRing.
// ctx: Pulumi context.
// opts: KeyringMemberOptions containing KeyRingID, Member, Role, and optional Pulumi options.
func CreateKeyringMember(
	ctx *pulumi.Context,
	opts *KeyringMemberOptions,
) (*gcpkms.KeyRingIAMMember, error) {
	name := fmt.Sprintf(
		"gcp-kms-iam-member-%s-%s-%s",
		sanitize.Text(opts.KeyRingID),
		sanitize.Text(opts.Member),
		sanitize.Text(opts.Role),
	)

	return gcpkms.NewKeyRingIAMMember(ctx, name, &gcpkms.KeyRingIAMMemberArgs{
		KeyRingId: pulumi.String(opts.KeyRingID),
		Role:      pulumi.String(opts.Role),
		Member:    pulumi.String(opts.Member),
	}, opts.PulumiOptions...)
}
