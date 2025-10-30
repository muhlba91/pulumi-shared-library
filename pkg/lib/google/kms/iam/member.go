package iam

import (
	"fmt"

	gcpkms "github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

// MemberArgs represents the arguments for creating a CryptoKey IAM member.
type MemberArgs struct {
	// CryptoKeyID is the ID of the CryptoKey to attach the IAM member to.
	CryptoKeyID string
	// Member is the member to assign the role to (e.g., "user:<email>").
	Member string
	// Role is the role to assign to the IAM member.
	Role string
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// CreateMember defines a new IAM member for a CryptoKey.
// ctx: Pulumi context.
// args: MemberArgs containing CryptoKeyID, Member, Role, and optional Pulumi options.
func CreateMember(
	ctx *pulumi.Context,
	args *MemberArgs,
) (*gcpkms.CryptoKeyIAMMember, error) {
	name := fmt.Sprintf(
		"gcp-kms-iam-member-%s-%s-%s",
		sanitize.Text(args.Member),
		sanitize.Text(args.CryptoKeyID),
		sanitize.Text(args.Role),
	)

	return gcpkms.NewCryptoKeyIAMMember(ctx, name, &gcpkms.CryptoKeyIAMMemberArgs{
		CryptoKeyId: pulumi.String(args.CryptoKeyID),
		Role:        pulumi.String(args.Role),
		Member:      pulumi.String(args.Member),
	}, args.PulumiOptions...)
}
