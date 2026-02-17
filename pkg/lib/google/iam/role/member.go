package role

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/projects"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

// MemberOptions represents the options for creating a service account member.
type MemberOptions struct {
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
// opts: MemberOptions containing member, roles, project, and optional Pulumi options.
func CreateMember(
	ctx *pulumi.Context,
	name string,
	opts *MemberOptions,
) ([]*projects.IAMMember, error) {
	var created []*projects.IAMMember

	for _, role := range opts.Roles {
		res, err := projects.NewIAMMember(ctx,
			fmt.Sprintf("gcp-iam-member-%s-%s", name, sanitize.Text(role)),
			&projects.IAMMemberArgs{
				Member:  opts.Member,
				Role:    pulumi.String(role),
				Project: opts.Project,
			},
			opts.PulumiOptions...,
		)
		if err != nil {
			return nil, err
		}
		created = append(created, res)
	}

	return created, nil
}
