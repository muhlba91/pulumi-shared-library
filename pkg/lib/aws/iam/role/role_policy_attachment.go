package role

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateRolePolicyAttachmentOptions defines options for creating an IAM Role Policy Attachment of a single role.
type CreateRolePolicyAttachmentOptions struct {
	// Role is the name of the IAM Role to attach the policy to.
	Role pulumi.StringInput
	// PolicyArn is the ARN of the IAM Policy to attach.
	PolicyArn pulumi.StringInput
	// PulumiOptions are additional options for the resource.
	PulumiOptions []pulumi.ResourceOption
}

// CreateRolePolicyAttachment attaches a policy to a single IAM Role without affecting other entities of the policy.
// ctx: Pulumi context.
// name: The logical name for the Pulumi resource (prefixed with "aws-iam-role-policy-attachment-").
// opts: Options for creating the role policy attachment.
func CreateRolePolicyAttachment(
	ctx *pulumi.Context,
	name string,
	opts *CreateRolePolicyAttachmentOptions,
) (*iam.RolePolicyAttachment, error) {
	return iam.NewRolePolicyAttachment(
		ctx,
		fmt.Sprintf("aws-iam-role-policy-attachment-%s", name),
		&iam.RolePolicyAttachmentArgs{
			Role:      opts.Role,
			PolicyArn: opts.PolicyArn,
		},
		opts.PulumiOptions...)
}
