package user

import (
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/iam"
)

// User bundles a Scaleway IAM user and its API key.
type User struct {
	// User is the Pulumi User resource.
	User *iam.User
	// Key is the Pulumi API key resource.
	Key *iam.ApiKey
}
