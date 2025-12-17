package role

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/projects"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CustomRoleArgs represents arguments for creating a GCP IAM custom role.
type CustomRoleArgs struct {
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
// args: CustomRoleArgs containing ID, Title, Description, Permissions, Project, and optional Pulumi options.
func CreateCustomRole(ctx *pulumi.Context, name string, args *CustomRoleArgs) (*projects.IAMCustomRole, error) {
	return projects.NewIAMCustomRole(
		ctx,
		fmt.Sprintf("gcp-iam-role-%s", name),
		&projects.IAMCustomRoleArgs{
			RoleId:      args.ID,
			Title:       args.Title,
			Description: args.Description,
			Stage:       pulumi.String("GA"),
			Permissions: pulumi.StringArray(args.Permissions),
			Project:     args.Project,
		},
		args.PulumiOptions...,
	)
}
