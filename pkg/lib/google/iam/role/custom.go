package role

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/projects"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CustomRoleOptions represents the options for creating a GCP IAM custom role.
type CustomRoleOptions struct {
	// ID is the roleId for the custom role (e.g. "myCustomRole").
	ID pulumi.StringInput
	// Title is the human-readable title for the role.
	Title pulumi.StringInput
	// Description is the role description.
	Description pulumi.StringInput
	// Permissions is the list of permissions to include in the role.
	Permissions []pulumi.StringInput
	// Project is the GCP project to create the role in.
	Project pulumi.StringInput
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// CreateCustomRole creates a GCP IAM custom role.
// ctx: Pulumi context.
// name: Name prefix for the custom role resource.
// opts: CustomRoleOptions containing ID, Title, Description, Permissions, Project, and optional Pulumi options.
func CreateCustomRole(ctx *pulumi.Context, name string, opts *CustomRoleOptions) (*projects.IAMCustomRole, error) {
	return projects.NewIAMCustomRole(
		ctx,
		fmt.Sprintf("gcp-iam-role-%s", name),
		&projects.IAMCustomRoleArgs{
			RoleId:      opts.ID,
			Title:       opts.Title,
			Description: opts.Description,
			Stage:       pulumi.String("GA"),
			Permissions: pulumi.StringArray(opts.Permissions),
			Project:     opts.Project,
		},
		opts.PulumiOptions...,
	)
}
