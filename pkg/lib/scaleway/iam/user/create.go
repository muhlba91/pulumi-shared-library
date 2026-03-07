package user

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/iam"
)

// CreateOptions holds options for creating an IAM user.
type CreateOptions struct {
	// Email is the email address associated with the user.
	Email pulumi.StringInput
	// Labels are key/value pairs to tag the user with.
	Labels []string
	// PulumiOptions are additional options to pass to Pulumi resource creation.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates an IAM user.
// ctx: The Pulumi context.
// name: The name of the IAM user.
// opts: Options for creating the IAM user.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*iam.User, error) {
	return iam.NewUser(ctx, fmt.Sprintf("scaleway-user-%s", name), &iam.UserArgs{
		Username: pulumi.String(name),
		Email:    opts.Email,
		Tags:     pulumi.ToStringArray(opts.Labels),
	}, opts.PulumiOptions...)
}
