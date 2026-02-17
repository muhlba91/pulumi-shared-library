package application

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/scaleway/iam/apikey"
	scwApp "github.com/muhlba91/pulumi-shared-library/pkg/lib/scaleway/iam/application"
	scwmodel "github.com/muhlba91/pulumi-shared-library/pkg/model/scaleway/iam/application"
)

// CreateOptions represents the options for creating a Scaleway IAM application.
type CreateOptions struct {
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
// opts: CreateOptions for creating the application.
func CreateApplication(
	ctx *pulumi.Context,
	opts *CreateOptions,
) (*scwmodel.Application, error) {
	application, errApp := scwApp.Create(ctx, opts.Name, &scwApp.CreateOptions{
		Description:   opts.Description,
		Labels:        opts.Labels,
		PulumiOptions: opts.PulumiOptions,
	})
	if errApp != nil {
		return nil, errApp
	}

	key, errKey := apikey.Create(ctx, opts.Name, &apikey.CreateOptions{
		ApplicationID:    application.ID(),
		DefaultProjectID: opts.DefaultProjectID,
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
