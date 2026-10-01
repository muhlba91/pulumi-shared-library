package random

import (
	randomsdk "github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/model/random"

	rModel "github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/rotation"
)

// BytesOptions holds optional parameters.
type BytesOptions struct {
	// Length is the desired number of random bytes.
	Length int
	// Rotation defines the rotation options for the resource.
	Rotation *rModel.Options
}

// CreateBytes creates a RandomId resource and returns BytesData.
// Defaults: length=16.
// ctx: Pulumi context.
// name: Name prefix for the resource.
// opts: Optional parameters for bytes generation.
func CreateBytes(ctx *pulumi.Context, name string, opts *BytesOptions) (*random.BytesData, error) {
	pulumiOpts := []pulumi.ResourceOption{}
	length := 16
	if opts != nil {
		if trigger, _ := rotation.Trigger(ctx, name, opts.Rotation); trigger != nil {
			pulumiOpts = append(pulumiOpts, pulumi.ReplacementTrigger(trigger))
		}
		if opts.Length != 0 {
			length = opts.Length
		}
	}

	b, err := randomsdk.NewRandomId(ctx, name, &randomsdk.RandomIdArgs{
		ByteLength: pulumi.Int(length),
	},
		pulumiOpts...)
	if err != nil {
		return nil, err
	}

	return &random.BytesData{
		Resource:  b,
		Hex:       b.Hex,
		Base64Std: b.B64Std,
		Base64URL: b.B64Url,
	}, nil
}
