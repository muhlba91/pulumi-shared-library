package serviceaccount

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/google/iam/serviceaccount"
	gmodel "github.com/muhlba91/pulumi-shared-library/pkg/model/google/iam/serviceaccount"
	rModel "github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
)

// CreateOptions represents the options for creating a Google service account user.
type CreateOptions struct {
	// Name is the name of the service account.
	Name string
	// DisplayName is the display name of the service account. Optional, defaults to Name.
	DisplayName *string
	// Description is the description of the service account. Optional.
	Description *string
	// Project is the GCP project ID where the service account will be created.
	Project pulumi.StringInput
	// Roles is a list of roles to assign to the service account.
	Roles []string
	// Rotation defines the rotation options for the key. Optional.
	Rotation *rModel.Options
	// CreateIgnoreAlreadyExists adopts an existing account with the same ID instead of failing (e.g. to migrate an
	// account created with google-native). Optional.
	CreateIgnoreAlreadyExists *bool
	// PulumiOptions are additional Pulumi resource options for the service account, its IAM members and its key.
	// Optional.
	PulumiOptions []pulumi.ResourceOption
}

// CreateServiceAccountUser creates a new service account, its project roles, and a key.
// ctx: Pulumi context.
// opts: CreateOptions for creating the service account user.
func CreateServiceAccountUser(
	ctx *pulumi.Context,
	opts *CreateOptions,
) (*gmodel.User, error) {
	serviceAccount, _, errSa := serviceaccount.CreateServiceAccount(ctx, &serviceaccount.CreateOptions{
		Name:                      opts.Name,
		DisplayName:               opts.DisplayName,
		Description:               opts.Description,
		Roles:                     opts.Roles,
		Project:                   opts.Project,
		CreateIgnoreAlreadyExists: opts.CreateIgnoreAlreadyExists,
		PulumiOptions:             opts.PulumiOptions,
	})
	if errSa != nil {
		return nil, errSa
	}

	keyOpts := append([]pulumi.ResourceOption{}, opts.PulumiOptions...)
	keyOpts = append(keyOpts, pulumi.DependsOn([]pulumi.Resource{serviceAccount}))
	key, errKey := serviceaccount.CreateKey(ctx, opts.Name, &serviceaccount.KeyOptions{
		ServiceAccount: serviceAccount.Name,
		Rotation:       opts.Rotation,
		PulumiOptions:  keyOpts,
	})
	if errKey != nil {
		return nil, errKey
	}

	return &gmodel.User{
		ServiceAccount: serviceAccount,
		Key:            key,
	}, nil
}
