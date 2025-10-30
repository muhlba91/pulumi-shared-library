package serviceaccount

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/google/iam/serviceaccount"
	gmodel "github.com/muhlba91/pulumi-shared-library/pkg/model/google/iam/serviceaccount"
)

// CreateServiceAccountUserArgs defines the input arguments for CreateServiceAccountUser function.
type CreateServiceAccountUserArgs struct {
	// Name is the name of the service account.
	Name string
	// Project is the GCP project ID where the service account will be created.
	Project pulumi.StringInput
	// Roles is a list of roles to assign to the service account.
	Roles []string
}

// CreateServiceAccountUser creates a new service account and key.
// ctx: Pulumi context.
// args: CreateServiceAccountUserArgs containing the name, project, and roles.
func CreateServiceAccountUser(
	ctx *pulumi.Context,
	args *CreateServiceAccountUserArgs,
) (*gmodel.User, error) {
	serviceAccount, _, errSa := serviceaccount.CreateServiceAccount(ctx, &serviceaccount.Args{
		Name:    args.Name,
		Roles:   args.Roles,
		Project: args.Project,
	})
	if errSa != nil {
		return nil, errSa
	}

	key, errKey := serviceaccount.CreateKey(ctx, args.Name, &serviceaccount.KeyArgs{
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
