package role

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreatePolicyAttachmentOptions defines options for creating an IAM Role Policy Attachment.
type CreatePolicyAttachmentOptions struct {
	// Roles is a list of IAM Role names to attach the policy to.
	Roles []pulumi.StringInput
	// PolicyArn is the ARN of the IAM Policy to attach.
	PolicyArn pulumi.StringInput
	// PulumiOptions are additional options for the resource.
	PulumiOptions []pulumi.ResourceOption
}

// CreatePolicyAttachment creates an IAM Role Policy Attachment.
// ctx: Pulumi context
// name: Name of the policy attachment
// opts: Options for creating the policy attachment
func CreatePolicyAttachment(
	ctx *pulumi.Context,
	name string,
	opts *CreatePolicyAttachmentOptions,
) (*iam.PolicyAttachment, error) {
	roles := pulumi.Array{}
	for _, r := range opts.Roles {
		roles = append(roles, r)
	}

	return iam.NewPolicyAttachment(
		ctx,
		fmt.Sprintf("aws-iam-role-policy-attachment-%s", name),
		&iam.PolicyAttachmentArgs{
			PolicyArn: opts.PolicyArn,
			Roles:     roles,
		},
		opts.PulumiOptions...)
}
