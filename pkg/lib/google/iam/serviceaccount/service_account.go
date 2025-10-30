package serviceaccount

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/projects"
	giam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/google/iam/role"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

// Args represents the arguments for creating a service account.
type Args struct {
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
// args: Args containing name, roles, project, and optional Pulumi options.
func CreateServiceAccount(
	ctx *pulumi.Context,
	args *Args,
) (*giam.ServiceAccount, []*projects.IAMMember, error) {
	sa, errSa := giam.NewServiceAccount(ctx,
		fmt.Sprintf("gcp-sa-%s", sanitize.Text(args.Name)),
		&giam.ServiceAccountArgs{
			AccountId:   pulumi.String(args.Name),
			DisplayName: pulumi.String(args.Name),
			Project:     args.Project,
		},
		args.PulumiOptions...,
	)
	if errSa != nil {
		return nil, nil, errSa
	}

	if len(args.Roles) > 0 {
		member, _ := sa.Email.ApplyT(func(email string) string {
			return fmt.Sprintf("serviceAccount:%s", email)
		}).(pulumi.StringOutput)

		created, errRoles := role.CreateMember(ctx, args.Name, &role.MemberArgs{
			Member:        member,
			Roles:         args.Roles,
			Project:       args.Project,
			PulumiOptions: args.PulumiOptions,
		})
		if errRoles != nil {
			return nil, nil, errRoles
		}
		return sa, created, nil
	}

	return sa, nil, nil
}
