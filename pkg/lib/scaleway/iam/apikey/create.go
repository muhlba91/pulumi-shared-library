package apikey

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/iam"

	rModel "github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/rotation"
)

// CreateOptions defines the options for creating an IAM API key.
type CreateOptions struct {
	// Description is the description of the IAM API key.
	Description pulumi.StringInput
	// UserID is the ID of the user for whom the API key is created.
	UserID pulumi.StringPtrInput
	// ApplicationID is the ID of the application associated with the API key.
	ApplicationID pulumi.StringPtrInput
	// DefaultProjectID is the default project ID for the API key.
	DefaultProjectID pulumi.StringPtrInput
	// Rotation defines the rotation options for the resource.
	Rotation *rModel.Options
	// PulumiOptions are additional options to pass to the resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new IAM API key for the specified user.
// ctx: The Pulumi context.
// name: The name of the API key resource.
// opts: The options for creating the API key.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*iam.ApiKey, error) {
	resName := fmt.Sprintf("scaleway-api-key-%s", name)

	pulumiOpts := append([]pulumi.ResourceOption{}, opts.PulumiOptions...)
	if trigger, _ := rotation.Trigger(ctx, resName, opts.Rotation); trigger != nil {
		pulumiOpts = append(pulumiOpts, pulumi.ReplacementTrigger(trigger))
	}

	return iam.NewApiKey(
		ctx,
		resName,
		&iam.ApiKeyArgs{
			Description:      opts.Description,
			UserId:           opts.UserID,
			ApplicationId:    opts.ApplicationID,
			DefaultProjectId: opts.DefaultProjectID,
		},
		pulumiOpts...,
	)
}
