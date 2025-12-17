package region

import (
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/config"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GetOrDefault returns the provided region if non-empty,
// otherwise falls back to the Pulumi AWS region configuration value.
// If neither is set, it returns an empty string.
// ctx: Pulumi context.
// region: The region input.
func GetOrDefault(ctx *pulumi.Context, region *string) *string {
	if region != nil && *region != "" {
		return region
	}

	if cfg := config.GetRegion(ctx); cfg != "" {
		return &cfg
	}

	return nil
}
