package policy

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/iam"
)

// CreateOptions defines the options for creating an IAM Policy.
type CreateOptions struct {
	// Name is the name of the IAM Policy.
	Name pulumi.StringInput
	// Description is the description of the IAM Policy.
	Description pulumi.StringInput
	// Rules are the rules of the IAM Policy.
	Rules []iam.PolicyRuleInput
	// UserID is the ID of the user to attach the policy to.
	UserID pulumi.StringPtrInput
	// ApplicationID is the ID of the application to attach the policy to.
	ApplicationID pulumi.StringPtrInput
	// Labels are the tags to apply to the IAM Policy.
	Labels []string
	// PulumiOptions are additional options for the resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new IAM Policy with the given options.
// ctx: The Pulumi context.
// name: The Pulumi resource name.
// opts: The options for creating the IAM Policy.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*iam.Policy, error) {
	return iam.NewPolicy(ctx, fmt.Sprintf("scaleway-policy-%s", name), &iam.PolicyArgs{
		Name:          opts.Name,
		Description:   opts.Description,
		UserId:        opts.UserID,
		ApplicationId: opts.ApplicationID,
		Rules:         iam.PolicyRuleArray(opts.Rules),
		Tags:          pulumi.ToStringArray(opts.Labels),
	}, opts.PulumiOptions...)
}
