package serviceaccount

import (
	svcaccount "github.com/pulumi/pulumi-gcp/sdk/v10/go/gcp/serviceaccount"
)

// User bundles a Google service account and its key.
type User struct {
	// ServiceAccount is the Pulumi service account resource.
	ServiceAccount *svcaccount.Account
	// Key is the Pulumi service account key resource.
	Key *svcaccount.Key
}
