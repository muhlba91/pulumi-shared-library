//nolint:revive // package name is fine as is
package user

import (
	"fmt"

	"github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/random"
	pgModel "github.com/muhlba91/pulumi-shared-library/pkg/model/postgresql"
)

const defaultPasswordLength = 32

// CreateOptions represents the created PostgreSQL user data.
type CreateOptions struct {
	// Username is the name of the PostgreSQL user to create.
	Username string
	// PulumiOptions are additional options to pass to the Postgresql Role resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a Postgresql role and a random password for the given username.
// ctx: Pulumi context
// options: CreateOptions for customizing the user creation
func Create(ctx *pulumi.Context, opts *CreateOptions) (*pgModel.UserData, error) {
	pw, err := random.CreatePassword(ctx, fmt.Sprintf("password-pg-user-%s", opts.Username), &random.PasswordOptions{
		Length:  defaultPasswordLength,
		Special: false,
	})
	if err != nil {
		return nil, err
	}

	role, err := postgresql.NewRole(ctx, fmt.Sprintf("pg-db-user-%s", opts.Username), &postgresql.RoleArgs{
		Name:           pulumi.String(opts.Username),
		Password:       pw.Password,
		CreateDatabase: pulumi.Bool(false),
		CreateRole:     pulumi.Bool(false),
		Login:          pulumi.Bool(true),
	}, opts.PulumiOptions...)
	if err != nil {
		return nil, err
	}

	return &pgModel.UserData{
		Password: pw.Password,
		User:     role,
	}, nil
}
