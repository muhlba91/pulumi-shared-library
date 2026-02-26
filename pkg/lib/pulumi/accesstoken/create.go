package accesstoken

import (
	"fmt"

	"github.com/pulumi/pulumi-pulumiservice/sdk/go/pulumiservice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	rModel "github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/rotation"
)

// CreateOptions defines the options for creating a Pulumi Access Token.
type CreateOptions struct {
	// Description is an optional description for the access token.
	Description pulumi.StringInput
	// Rotation defines the rotation options for the resource.
	Rotation *rModel.Options
}

// Create creates a new Pulumi Access Token with the given name and options.
// ctx: The Pulumi context.
// name: The name to use for the access token resource.
// opts: Options for creating the access token.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*pulumiservice.AccessToken, error) {
	resName := fmt.Sprintf("pulumi-access-token-%s", name)

	pulumiOpts := []pulumi.ResourceOption{}
	if trigger, _ := rotation.Trigger(ctx, resName, opts.Rotation); trigger != nil {
		pulumiOpts = append(pulumiOpts, pulumi.ReplacementTrigger(trigger))
	}

	return pulumiservice.NewAccessToken(
		ctx,
		resName,
		&pulumiservice.AccessTokenArgs{
			Description: opts.Description,
		},
		pulumiOpts...,
	)
}
