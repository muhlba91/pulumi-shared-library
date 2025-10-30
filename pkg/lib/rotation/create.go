package rotation

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-time/sdk/go/time"
)

// Create creates a rotating resource.
// If days is <= 0 it defaults to 30.
// ctx: Pulumi context.
// name: Name prefix for the rotation resource.
// days: Number of days between rotations.
func Create(ctx *pulumi.Context, name string, days int) (pulumi.CustomResource, error) {
	if days <= 0 {
		days = 30
	}

	return time.NewRotating(ctx,
		fmt.Sprintf("rotation-%s", name),
		&time.RotatingArgs{
			RotationDays: pulumi.Int(days),
		},
	)
}
