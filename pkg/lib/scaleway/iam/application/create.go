package user

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/iam"
)

// CreateOptions holds options for creating an IAM application.
type CreateOptions struct {
	// Description is the description associated with the application.
	Description pulumi.StringPtrInput
	// Labels are key/value pairs to tag the application with.
	Labels []string
	// PulumiOptions are additional options to pass to Pulumi resource creation.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates an IAM application.
// ctx: The Pulumi context.
// name: The name of the IAM application.
// opts: Options for creating the IAM application.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*iam.Application, error) {
	return iam.NewApplication(ctx, fmt.Sprintf("scaleway-application-%s", name), &iam.ApplicationArgs{
		Name:        pulumi.String(name),
		Description: opts.Description,
		Tags:        pulumi.ToStringArray(opts.Labels),
	}, opts.PulumiOptions...)
}
