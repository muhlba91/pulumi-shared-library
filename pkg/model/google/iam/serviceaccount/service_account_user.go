package serviceaccount

import (
	svcaccount "github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/serviceaccount"
	iam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
)

// User defines a user for a service account.
type User struct {
	// ServiceAccount is the Pulumi ServiceAccount resource.
	ServiceAccount *iam.ServiceAccount
	// Key is the Pulumi ServiceAccount key resource.
	Key *svcaccount.Key
}
