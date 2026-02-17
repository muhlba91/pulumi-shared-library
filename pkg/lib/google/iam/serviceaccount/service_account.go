package serviceaccount

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/projects"
	giam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/google/iam/role"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

// CreateOptions represents the options for creating a service account.
type CreateOptions struct {
	// Name is the name of the service account.
	Name string
	// Roles are the roles to assign to the service account.
	Roles []string
	// Project is the GCP project ID where the service account will be created.
	Project pulumi.StringInput
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// CreateServiceAccount creates a Google Service Account and (optionally) attaches IAM members for the provided roles.
// ctx: Pulumi context.
// opts: CreateOptions containing name, roles, project, and optional Pulumi options.
func CreateServiceAccount(
	ctx *pulumi.Context,
	opts *CreateOptions,
) (*giam.ServiceAccount, []*projects.IAMMember, error) {
	sa, errSa := giam.NewServiceAccount(ctx,
		fmt.Sprintf("gcp-sa-%s", sanitize.Text(opts.Name)),
		&giam.ServiceAccountArgs{
			AccountId:   pulumi.String(opts.Name),
			DisplayName: pulumi.String(opts.Name),
			Project:     opts.Project,
		},
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
