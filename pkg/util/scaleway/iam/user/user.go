//nolint:revive // package name is fine as is
package user

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/scaleway/iam/apikey"
	"github.com/muhlba91/pulumi-shared-library/pkg/lib/scaleway/iam/user"
	scwmodel "github.com/muhlba91/pulumi-shared-library/pkg/model/scaleway/iam/user"
)

// CreateUserArgs defines the input arguments for CreateUser function.
type CreateUserArgs struct {
	// Name is the name of the service account.
	Name string
	// Email is the email address associated with the user.
	Email pulumi.StringInput
	// DefaultProjectID is the default project ID for the API key.
	DefaultProjectID pulumi.StringPtrInput
	// Labels are key/value pairs to tag the user with.
	Labels []string
	// PulumiOptions are additional options to pass to Pulumi resource creation.
	PulumiOptions []pulumi.ResourceOption
}

// CreateUser creates a new service account and key.
// ctx: Pulumi context.
// args: Arguments for creating the service account.
func CreateUser(
	ctx *pulumi.Context,
	args *CreateUserArgs,
) (*scwmodel.User, error) {
	user, errUser := user.Create(ctx, args.Name, &user.CreateOptions{
		Email:         args.Email,
		Labels:        args.Labels,
		PulumiOptions: args.PulumiOptions,
	})
	if errUser != nil {
		return nil, errUser
	}

	key, errKey := apikey.Create(ctx, args.Name, &apikey.CreateOptions{
		UserID:           user.ID(),
		DefaultProjectID: args.DefaultProjectID,
		PulumiOptions: []pulumi.ResourceOption{
			pulumi.DependsOn([]pulumi.Resource{user}),
		},
	})
	if errKey != nil {
		return nil, errKey
	}

	return &scwmodel.User{
		User: user,
		Key:  key,
	}, nil
}
