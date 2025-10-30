package policy

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateOptions defines the options for creating an IAM Policy.
type CreateOptions struct {
	// Name is the name of the IAM Policy.
	Name pulumi.StringInput
	// Description is the description of the IAM Policy.
	Description pulumi.StringInput
	// Policy is the JSON policy document.
	Policy pulumi.StringInput
	// Labels are the tags to apply to the IAM Policy.
	Labels map[string]string
	// PulumiOptions are additional options for the resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new IAM Policy with the given options.
// ctx: The Pulumi context.
// name: The Pulumi resource name.
// opts: The options for creating the IAM Policy.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*iam.Policy, error) {
	return iam.NewPolicy(ctx, fmt.Sprintf("aws-iam-role-ci-policy-%s", name), &iam.PolicyArgs{
		Name:        opts.Name,
		Description: opts.Description,
		Policy:      opts.Policy,
		Tags:        pulumi.ToStringMap(opts.Labels),
	}, opts.PulumiOptions...)
}
