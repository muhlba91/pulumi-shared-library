package application

import (
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/iam"
)

// Application defines an application for a service account.
type Application struct {
	// Application is the Pulumi Application resource.
	Application *iam.Application
	// Key is the Pulumi API key resource.
	Key *iam.ApiKey
}
