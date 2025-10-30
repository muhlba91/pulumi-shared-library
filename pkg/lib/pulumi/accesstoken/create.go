package accesstoken

import (
	"fmt"

	"github.com/pulumi/pulumi-pulumiservice/sdk/go/pulumiservice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateOptions defines the options for creating a Pulumi Access Token.
type CreateOptions struct {
	// Description is an optional description for the access token.
	Description pulumi.StringInput
}

// Create creates a new Pulumi Access Token with the given name and options.
// ctx: The Pulumi context.
// name: The name to use for the access token resource.
// opts: Options for creating the access token.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*pulumiservice.AccessToken, error) {
	return pulumiservice.NewAccessToken(
		ctx,
		fmt.Sprintf("pulumi-access-token-%s", name),
		&pulumiservice.AccessTokenArgs{
			Description: opts.Description,
		},
	)
}
