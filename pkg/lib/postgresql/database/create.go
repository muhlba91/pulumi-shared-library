package database

import (
	"fmt"

	"github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	pgModel "github.com/muhlba91/pulumi-shared-library/pkg/model/postgresql"
)

// CreateOptions represents the options for creating a PostgreSQL database.
type CreateOptions struct {
	// Name is the name of the database to create.
	Name string
	// Owner is the owner of the database.
	Owner pgModel.UserData
	// PulumiOptions are additional options to pass to the Postgresql Database resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a PostgreSQL database with the given options.
// ctx: Pulumi context
// opts: CreateOptions for customizing the database creation
func Create(ctx *pulumi.Context, opts *CreateOptions) (*postgresql.Database, error) {
	optsWithDepends := append([]pulumi.ResourceOption{}, opts.PulumiOptions...)
	optsWithDepends = append(optsWithDepends, pulumi.DependsOn([]pulumi.Resource{opts.Owner.User}))

	return postgresql.NewDatabase(ctx, fmt.Sprintf("pg-db-%s", opts.Name), &postgresql.DatabaseArgs{
		Name:  pulumi.String(opts.Name),
		Owner: opts.Owner.User.Name,
	}, optsWithDepends...)
}
