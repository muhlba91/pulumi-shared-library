package rotation

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-time/sdk/go/time"

	rModel "github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
)

// Create creates a rotating resource.
// If days is <= 0 it defaults to 30.
// ctx: Pulumi context.
// opts: Options for creating the rotation resource.
func Create(ctx *pulumi.Context, opts *rModel.Options) (pulumi.CustomResource, error) {
	days := opts.Days
	if days <= 0 {
		days = 30
	}

	return time.NewRotating(ctx,
		fmt.Sprintf("rotation-%s", opts.Name),
		&time.RotatingArgs{
			RotationDays: pulumi.Int(days),
		},
	)
}
