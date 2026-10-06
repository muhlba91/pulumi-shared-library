package serviceaccount

import (
	svcaccount "github.com/pulumi/pulumi-gcp/sdk/v10/go/gcp/serviceaccount"
	iam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
)

// User bundles a Google service account and its key.
type User struct {
	// ServiceAccount is the Pulumi ServiceAccount resource.
	ServiceAccount *iam.ServiceAccount
	// Key is the Pulumi ServiceAccount key resource.
	Key *svcaccount.Key
}
