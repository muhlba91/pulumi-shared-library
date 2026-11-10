// Package serviceaccount creates Google service account users (account, roles, key) with the google-native provider.
//
// Deprecated: use github.com/muhlba91/pulumi-shared-library/pkg/util/google/iam/serviceaccount (gcp provider); the
// migration is described on CreateServiceAccountUser.
package serviceaccount

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/google/iam/serviceaccount"
	//nolint:staticcheck // the deprecated native user builds on the deprecated native account
	nativeserviceaccount "github.com/muhlba91/pulumi-shared-library/pkg/lib/google/native/iam/serviceaccount"
	//nolint:staticcheck // the deprecated native user returns the deprecated native model
	gmodel "github.com/muhlba91/pulumi-shared-library/pkg/model/google/native/iam/serviceaccount"
)

// CreateOptions represents the options for creating a Google service account user.
type CreateOptions struct {
	// Name is the name of the service account.
	Name string
	// Project is the GCP project ID where the service account will be created.
	Project pulumi.StringInput
	// Roles is a list of roles to assign to the service account.
	Roles []string
	// PulumiOptions are additional Pulumi resource options for the service account and its IAM members (e.g.
	// pulumi.RetainOnDelete for the migration to the gcp provider). Optional.
	PulumiOptions []pulumi.ResourceOption
}

// CreateServiceAccountUser creates a new service account, its project roles, and a key.
// ctx: Pulumi context.
// opts: CreateOptions for creating the service account user.
//
// Deprecated: google-native is deprecated upstream. Use CreateServiceAccountUser of
// github.com/muhlba91/pulumi-shared-library/pkg/util/google/iam/serviceaccount (gcp provider) instead.
// It keeps the resource names (gcp-sa-<name>, gcp-sa-key-<name>), the account ID, the display name, the roles' IAM
// members and the key, so an existing account and its key move over without being recreated, in three deployments:
//
//  1. Add pulumi.RetainOnDelete(true) to PulumiOptions here and deploy: removing the native resource later leaves
//     the account in GCP.
//  2. Switch the import to the gcp package (same options; its User holds a gcp serviceaccount.Account), set
//     CreateIgnoreAlreadyExists to true and deploy: the gcp resource adopts the existing account and the native
//     resource is dropped from the state without deleting the account. The key and the IAM members keep their names
//     and inputs (the account's email and its resource name projects/<project>/serviceAccounts/<email> are the same)
//     and stay untouched; check the preview shows only the create of the gcp account and the delete of the native
//     one.
//  3. Remove CreateIgnoreAlreadyExists and deploy (an in-place update of a provider-only field).
//
// Once no stack uses this package, it and the google-native dependency are removed.
func CreateServiceAccountUser(
	ctx *pulumi.Context,
	opts *CreateOptions,
) (*gmodel.User, error) {
	serviceAccount, _, errSa := nativeserviceaccount.CreateServiceAccount(ctx, &nativeserviceaccount.CreateOptions{
		Name:          opts.Name,
		Roles:         opts.Roles,
		Project:       opts.Project,
		PulumiOptions: opts.PulumiOptions,
	})
	if errSa != nil {
		return nil, errSa
	}

	key, errKey := serviceaccount.CreateKey(ctx, opts.Name, &serviceaccount.KeyOptions{
		ServiceAccount: serviceAccount.Name,
		PulumiOptions: []pulumi.ResourceOption{
			pulumi.DependsOn([]pulumi.Resource{serviceAccount}),
		},
	})
	if errKey != nil {
		return nil, errKey
	}

	return &gmodel.User{
		ServiceAccount: serviceAccount,
		Key:            key,
	}, nil
}
