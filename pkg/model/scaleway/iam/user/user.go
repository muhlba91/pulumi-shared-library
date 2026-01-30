package serviceaccount

import (
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/iam"
)

// User defines a user for a service account.
type User struct {
	// User is the Pulumi User resource.
	User *iam.User
	// Key is the Pulumi API key resource.
	Key *iam.ApiKey
}
