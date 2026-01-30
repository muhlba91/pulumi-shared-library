package accesskey

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateOptions defines the options for creating an IAM access key.
type CreateOptions struct {
	// UserName is the name of the IAM user to create the access key for.
	UserName string
	// User is the IAM user resource to associate the access key with.
	User pulumi.Resource
	// PulumiOptions are additional options to pass to the resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new IAM access key for the specified user.
// ctx: The Pulumi context.
// opts: The options for creating the access key.
func Create(ctx *pulumi.Context, opts *CreateOptions) (*iam.AccessKey, error) {
	optsWithDepends := append([]pulumi.ResourceOption{}, opts.PulumiOptions...)
	optsWithDepends = append(optsWithDepends, pulumi.DependsOn([]pulumi.Resource{opts.User}))

	return iam.NewAccessKey(
		ctx,
		fmt.Sprintf("aws-access-key-%s", opts.UserName),
		&iam.AccessKeyArgs{
			User: pulumi.String(opts.UserName),
		},
		optsWithDepends...,
	)
}
