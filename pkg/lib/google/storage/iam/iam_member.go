package iam

import (
	gcpStorage "github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

// MemberArgs represents the arguments for creating a GCS bucket IAM member.
type MemberArgs struct {
	// BucketID is the ID of the GCS bucket.
	BucketID string
	// Member is the member ID to create the IAM member for.
	Member string
	// Role is the role to assign to the IAM member.
	Role string
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// CreateIAMMember defines a new IAM member for a GCS bucket.
// ctx: The Pulumi context.
// args: MemberArgs containing member, role, bucket ID, and optional Pulumi options.
func CreateIAMMember(
	ctx *pulumi.Context,
	args *MemberArgs,
) (*gcpStorage.BucketIAMMember, error) {
	name := "gcp-gcs-iam-member-" + sanitize.Text(
		args.BucketID,
	) + "-" + sanitize.Text(
		args.Member,
	) + "-" + sanitize.Text(
		args.Role,
	)

	iamMember, err := gcpStorage.NewBucketIAMMember(ctx, name, &gcpStorage.BucketIAMMemberArgs{
		Bucket: pulumi.String(args.BucketID),
		Role:   pulumi.String(args.Role),
		Member: pulumi.String(args.Member),
	}, args.PulumiOptions...)
	if err != nil {
		return nil, err
	}
	return iamMember, nil
}
