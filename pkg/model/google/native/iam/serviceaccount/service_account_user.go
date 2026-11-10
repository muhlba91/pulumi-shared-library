// Package serviceaccount holds the models of service accounts created with the google-native provider.
//
// Deprecated: use github.com/muhlba91/pulumi-shared-library/pkg/model/google/iam/serviceaccount (gcp provider).
package serviceaccount

import (
	svcaccount "github.com/pulumi/pulumi-gcp/sdk/v10/go/gcp/serviceaccount"
	iam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
)

// User bundles a Google service account (google-native) and its key.
//
// Deprecated: use User of github.com/muhlba91/pulumi-shared-library/pkg/model/google/iam/serviceaccount.
type User struct {
	// ServiceAccount is the Pulumi ServiceAccount resource.
	ServiceAccount *iam.ServiceAccount
	// Key is the Pulumi ServiceAccount key resource.
	Key *svcaccount.Key
}
