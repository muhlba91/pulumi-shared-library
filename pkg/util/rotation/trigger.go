package rotation

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	rLib "github.com/muhlba91/pulumi-shared-library/pkg/lib/rotation"
	rModel "github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
)

// Trigger creates a rotation schedule if the options are set and returns its timestamp to be used as a trigger.
// Sets the name in the options to the provided name if it is not set, which modifies the provided options.
// ctx: Pulumi context.
// name: Name to use for the rotation schedule if the options do not provide one.
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
