package serviceaccount

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/google/iam/serviceaccount"
	gmodel "github.com/muhlba91/pulumi-shared-library/pkg/model/google/iam/serviceaccount"
)

// CreateOptions represents the options for creating a Google service account user.
type CreateOptions struct {
	// Name is the name of the service account.
	Name string
	// Project is the GCP project ID where the service account will be created.
	Project pulumi.StringInput
	// Roles is a list of roles to assign to the service account.
	Roles []string
}

// CreateServiceAccountUser creates a new service account and key.
// ctx: Pulumi context.
// opts: CreateOptions for creating the service account user.
func CreateServiceAccountUser(
	ctx *pulumi.Context,
	opts *CreateOptions,
) (*gmodel.User, error) {
	serviceAccount, _, errSa := serviceaccount.CreateServiceAccount(ctx, &serviceaccount.CreateOptions{
		Name:    opts.Name,
		Roles:   opts.Roles,
		Project: opts.Project,
	})
	if errSa != nil {
		return nil, errSa
	}

	key, errKey := serviceaccount.CreateKey(ctx, opts.Name, &serviceaccount.KeyOptions{
		ServiceAccount: serviceAccount.Name,
		PulumiOptions: []pulumi.ResourceOption{
			pulumi.DependsOn([]pulumi.Resource{serviceAccount}),
		},
	})
	if errKey != nil {
		return nil, errKey
	}

	return &gmodel.User{
		ServiceAccount: serviceAccount,
		Key:            key,
	}, nil
}
