package role

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v10/go/gcp/projects"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

// MemberCondition represents an IAM condition restricting a role binding.
type MemberCondition struct {
	// Title is the title of the condition.
	Title string
	// Description is the description of the condition. Optional.
	Description *string
	// Expression is the condition expression in Common Expression Language syntax.
	Expression pulumi.StringInput
}

// MemberOptions represents the options for creating a service account member.
type MemberOptions struct {
	// Member is the member ID to create the IAM member for.
	Member pulumi.StringInput
	// Roles are the roles to assign to the IAM member.
	Roles []string
	// Project is the GCP project ID where the IAM member will be created.
	Project pulumi.StringInput
	// Condition restricts the role bindings. Optional.
	Condition *MemberCondition
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

	var condition *projects.IAMMemberConditionArgs
	if opts.Condition != nil {
		condition = &projects.IAMMemberConditionArgs{
			Title:       pulumi.String(opts.Condition.Title),
			Description: pulumi.StringPtrFromPtr(opts.Condition.Description),
			Expression:  opts.Condition.Expression,
		}
	}

	for _, role := range opts.Roles {
		res, err := projects.NewIAMMember(ctx,
			fmt.Sprintf("gcp-iam-member-%s-%s", name, sanitize.Text(role)),
			&projects.IAMMemberArgs{
				Member:    opts.Member,
				Role:      pulumi.String(role),
				Project:   opts.Project,
				Condition: condition,
			},
			opts.PulumiOptions...,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create the iam member for the role %q: %w", role, err)
		}
		created = append(created, res)
	}

	return created, nil
}
