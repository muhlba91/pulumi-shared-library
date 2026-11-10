// Package serviceaccount creates Google service accounts with the google-native provider.
//
// Deprecated: use github.com/muhlba91/pulumi-shared-library/pkg/lib/google/iam/serviceaccount (gcp provider); the
// migration is described on CreateServiceAccount.
package serviceaccount

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v10/go/gcp/projects"
	giam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/google/iam/role"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

// CreateOptions represents the options for creating a service account.
type CreateOptions struct {
	// Name is the name of the service account, used as account ID and display name.
	// Must be 6 to 30 characters long, as required for GCP service account IDs.
	Name string
	// Roles are the project roles to assign to the service account. Optional.
	Roles []string
	// Project is the GCP project ID where the service account will be created.
	Project pulumi.StringInput
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// CreateServiceAccount creates a Google Service Account and (optionally) attaches IAM members for the provided roles.
// ctx: Pulumi context.
// opts: CreateOptions containing name, roles, project, and optional Pulumi options.
// Returns the service account and the created IAM members (nil if no roles are provided).
//
// Deprecated: google-native is deprecated upstream. Use CreateServiceAccount of
// github.com/muhlba91/pulumi-shared-library/pkg/lib/google/iam/serviceaccount (gcp provider) instead.
// It keeps the resource name (gcp-sa-<name>), the account ID, the display name and the roles' IAM members, so an
// existing account moves over without being recreated, in three deployments:
//
//  1. Add pulumi.RetainOnDelete(true) to PulumiOptions here and deploy: removing the native resource later leaves
//     the account in GCP.
//  2. Switch the import to the gcp package (same options; its CreateServiceAccount returns a gcp
//     serviceaccount.Account), set CreateIgnoreAlreadyExists to true and deploy: the gcp resource adopts the
//     existing account and the native resource is dropped from the state without deleting the account. IAM
//     members and keys (CreateKey) keep their names and inputs (the account's email and its resource name
//     projects/<project>/serviceAccounts/<email> are the same) and stay untouched; check the preview shows only the
//     create of the gcp account and the delete of the native one.
//  3. Remove CreateIgnoreAlreadyExists and deploy (an in-place update of a provider-only field).
//
// Once no stack uses this package, it and the google-native dependency are removed.
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
