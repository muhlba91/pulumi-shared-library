package project

import (
	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/config"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GetOrDefault returns the provided project if non-empty,
// otherwise falls back to the Pulumi GCP project configuration value.
// If neither is set, it returns an empty string.
// ctx: Pulumi context.
// project: The project input.
func GetOrDefault(ctx *pulumi.Context, project *string) *string {
	if project != nil && *project != "" {
		return project
	}

	if cfg := config.GetProject(ctx); cfg != "" {
		return &cfg
	}

	return nil
}
