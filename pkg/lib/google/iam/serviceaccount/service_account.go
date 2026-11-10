package serviceaccount

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v10/go/gcp/projects"
	gsa "github.com/pulumi/pulumi-gcp/sdk/v10/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/google/iam/role"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

// CreateOptions represents the options for creating a service account.
type CreateOptions struct {
	// Name is the name of the service account, used as account ID (and display name unless DisplayName is set).
	// Must be 6 to 30 characters long, as required for GCP service account IDs.
	Name string
	// DisplayName is the display name of the service account. Optional, defaults to Name.
	DisplayName *string
	// Description is the description of the service account. Optional.
	Description *string
	// Roles are the project roles to assign to the service account. Optional.
	Roles []string
	// Project is the GCP project ID where the service account will be created.
	Project pulumi.StringInput
	// CreateIgnoreAlreadyExists adopts an existing account with the same ID instead of failing (e.g. to migrate an
	// account created with google-native). Optional.
	CreateIgnoreAlreadyExists *bool
	// PulumiOptions are additional Pulumi resource options for the service account and its IAM members. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// CreateServiceAccount creates a Google Service Account and (optionally) attaches IAM members for the provided roles.
// ctx: Pulumi context.
// opts: CreateOptions containing name, roles, project, and optional display name, description and Pulumi options.
// Returns the service account and the created IAM members (nil if no roles are provided).
func CreateServiceAccount(
	ctx *pulumi.Context,
	opts *CreateOptions,
) (*gsa.Account, []*projects.IAMMember, error) {
	args := &gsa.AccountArgs{
		AccountId:   pulumi.String(opts.Name),
		DisplayName: pulumi.String(defaults.GetOrDefault(opts.DisplayName, opts.Name)),
		Project:     opts.Project,
	}
	if opts.Description != nil {
		args.Description = pulumi.String(*opts.Description)
	}
	if opts.CreateIgnoreAlreadyExists != nil {
		args.CreateIgnoreAlreadyExists = pulumi.Bool(*opts.CreateIgnoreAlreadyExists)
	}

	sa, errSa := gsa.NewAccount(ctx,
		fmt.Sprintf("gcp-sa-%s", sanitize.Text(opts.Name)),
		args,
		opts.PulumiOptions...,
	)
	if errSa != nil {
		return nil, nil, errSa
	}

	if len(opts.Roles) > 0 {
		member, _ := sa.Email.ApplyT(func(email string) string {
			return fmt.Sprintf("serviceAccount:%s", email)
		}).(pulumi.StringOutput)

		created, errRoles := role.CreateMember(ctx, opts.Name, &role.MemberOptions{
			Member:        member,
			Roles:         opts.Roles,
			Project:       opts.Project,
			PulumiOptions: opts.PulumiOptions,
		})
		if errRoles != nil {
			return nil, nil, errRoles
		}
		return sa, created, nil
	}

	return sa, nil, nil
}
