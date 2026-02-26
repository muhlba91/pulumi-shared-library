package rotation

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	rLib "github.com/muhlba91/pulumi-shared-library/pkg/lib/rotation"
	rModel "github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
)

// Trigger creates a rotation schedule if the options are set.
// Overwrites the name in the options with the provided name.
// ctx: Pulumi context.
// opts: Rotation options. If nil, no rotation schedule will be created.
func Trigger(ctx *pulumi.Context, name string, opts *rModel.Options) (*pulumi.StringOutput, error) {
	if opts == nil {
		//nolint:nilnil // No rotation options provided, so no trigger is needed.
		return nil, nil
	}

	if opts.Name == nil {
		opts.Name = &name
	}
	rotating, err := rLib.Create(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &rotating.Rfc3339, nil
}
