//nolint:revive // package name is fine as is
package user

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

// CreateOptions holds options for creating an IAM user.
type CreateOptions struct {
	// Policies is a list of managed IAM policies to attach to the user.
	Policies []*iam.Policy
	// Labels are key/value pairs to tag the user with.
	Labels map[string]string
	// PulumiOptions are additional options to pass to Pulumi resource creation.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates an IAM user and attaches the provided managed policies.
// ctx: The Pulumi context.
// name: The name of the IAM user.
// opts: Options for creating the IAM user.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*iam.User, error) {
	user, err := iam.NewUser(ctx, fmt.Sprintf("aws-user-%s", name), &iam.UserArgs{
		Tags: pulumi.ToStringMap(opts.Labels),
	}, opts.PulumiOptions...)
	if err != nil {
		return nil, err
	}

	for _, p := range opts.Policies {
		p.Arn.ApplyT(func(arn string) (*iam.UserPolicyAttachment, error) {
			optsWithDepends := append(
				make([]pulumi.ResourceOption, 0, len(opts.PulumiOptions)+1),
				opts.PulumiOptions...)
			optsWithDepends = append(optsWithDepends, pulumi.DependsOn([]pulumi.Resource{user, p}))
			return iam.NewUserPolicyAttachment(
				ctx,
				fmt.Sprintf("aws-user-policy-%s-%s", name, sanitize.Text(arn)),
				&iam.UserPolicyAttachmentArgs{
					User:      user.Name,
					PolicyArn: pulumi.String(arn),
				},
				optsWithDepends...)
		})
	}

	return user, nil
}
