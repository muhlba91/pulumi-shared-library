package role

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v10/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// BindingOptions represents the options for creating a service account IAM Binding.
type BindingOptions struct {
	// ServiceAccount is the Service Account ID to create the IAM Binding for.
	ServiceAccount pulumi.StringInput
	// Role is the role to assign to the IAM Binding.
	Role pulumi.StringInput
	// Members are the members the role is granted to on the Service Account.
	Members []pulumi.StringInput
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// CreateBinding creates an IAM Binding of a role for a Service Account.
// ctx: Pulumi context.
// name: Name for the IAM Binding resource.
// opts: BindingOptions containing service account, role, members, and optional Pulumi options.
func CreateBinding(
	ctx *pulumi.Context,
	name string,
	opts *BindingOptions,
) (*serviceaccount.IAMBinding, error) {
	return serviceaccount.NewIAMBinding(
		ctx,
		fmt.Sprintf("gcp-iam-identity-member-%s", name),
		&serviceaccount.IAMBindingArgs{
			ServiceAccountId: opts.ServiceAccount,
			Role:             opts.Role,
			Members:          pulumi.StringArray(opts.Members),
		},
		opts.PulumiOptions...)
}
