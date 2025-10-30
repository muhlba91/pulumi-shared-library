package role

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateOptions defines the input parameters for creating an IAM Role.
type CreateOptions struct {
	// Name of the IAM Role
	Name pulumi.StringInput
	// Description of the IAM Role
	Description pulumi.StringInput
	// AssumeRolePolicy is the policy that grants an entity permission to assume the role.
	AssumeRolePolicy pulumi.StringInput
	// Labels (Tags) to apply to the IAM Role
	Labels map[string]string
	// Additional Pulumi Resource Options
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new IAM Role with the specified options.
// ctx: Pulumi context
// name: A unique name for the IAM Role resource
// opts: CreateOptions containing configuration for the IAM Role
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*iam.Role, error) {
	return iam.NewRole(ctx, fmt.Sprintf("aws-iam-role-ci-%s", name), &iam.RoleArgs{
		Name:             opts.Name,
		Description:      opts.Description,
		AssumeRolePolicy: opts.AssumeRolePolicy,
		Tags:             pulumi.ToStringMap(opts.Labels),
	}, opts.PulumiOptions...)
}
