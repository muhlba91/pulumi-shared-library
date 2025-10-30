package role

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/projects"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

// MemberArgs represents the arguments for creating a service account member.
type MemberArgs struct {
	// Member is the member ID to create the IAM member for.
	Member pulumi.StringInput
	// Roles are the roles to assign to the IAM member.
	Roles []string
	// Project is the GCP project ID where the IAM member will be created.
	Project pulumi.StringInput
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// CreateMember creates a gcp.projects.IAMMember for each role.
// ctx: Pulumi context.
// name: Name prefix for the IAM member resources.
// args: MemberArgs containing member, roles, project, and optional Pulumi options.
func CreateMember(
	ctx *pulumi.Context,
	name string,
	args *MemberArgs,
) ([]*projects.IAMMember, error) {
	var created []*projects.IAMMember

	for _, role := range args.Roles {
		res, err := projects.NewIAMMember(ctx,
			fmt.Sprintf("gcp-iam-member-%s-%s", name, sanitize.Text(role)),
			&projects.IAMMemberArgs{
				Member:  args.Member,
				Role:    pulumi.String(role),
				Project: args.Project,
			},
			args.PulumiOptions...,
		)
		if err != nil {
			return nil, err
		}
		created = append(created, res)
	}

	return created, nil
}
