package application

import (
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/iam"
)

// Application bundles a Scaleway IAM application and its API key.
type Application struct {
	// Application is the Pulumi Application resource.
	Application *iam.Application
	// Key is the Pulumi API key resource.
	Key *iam.ApiKey
}
