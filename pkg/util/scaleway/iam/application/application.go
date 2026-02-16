package application

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/scaleway/iam/apikey"
	scwApp "github.com/muhlba91/pulumi-shared-library/pkg/lib/scaleway/iam/application"
	scwmodel "github.com/muhlba91/pulumi-shared-library/pkg/model/scaleway/iam/application"
)

// CreateApplicationArgs defines the input arguments for CreateApplication function.
type CreateApplicationArgs struct {
	// Name is the name of the application.
	Name string
	// Description is a brief description of the application.
	Description pulumi.StringPtrInput
	// DefaultProjectID is the default project ID for the API key.
	DefaultProjectID pulumi.StringPtrInput
	// Labels are key/value pairs to tag the application with.
	Labels []string
	// PulumiOptions are additional options to pass to Pulumi resource creation.
	PulumiOptions []pulumi.ResourceOption
}

// CreateApplication creates a new application and key.
// ctx: Pulumi context.
// args: Arguments for creating the application.
func CreateApplication(
	ctx *pulumi.Context,
	args *CreateApplicationArgs,
) (*scwmodel.Application, error) {
	application, errApp := scwApp.Create(ctx, args.Name, &scwApp.CreateOptions{
		Description:   args.Description,
		Labels:        args.Labels,
		PulumiOptions: args.PulumiOptions,
	})
	if errApp != nil {
		return nil, errApp
	}

	key, errKey := apikey.Create(ctx, args.Name, &apikey.CreateOptions{
		ApplicationID:    application.ID(),
		DefaultProjectID: args.DefaultProjectID,
		PulumiOptions: []pulumi.ResourceOption{
			pulumi.DependsOn([]pulumi.Resource{application}),
		},
	})
	if errKey != nil {
		return nil, errKey
	}

	return &scwmodel.Application{
		Application: application,
		Key:         key,
	}, nil
}
