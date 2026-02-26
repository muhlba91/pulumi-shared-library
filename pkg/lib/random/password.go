package random

import (
	randomsdk "github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/model/random"

	rModel "github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/rotation"
)

// PasswordOptions holds optional parameters.
type PasswordOptions struct {
	// Length is the desired length of the generated password.
	Length int
	// Special indicates whether to include special characters in the password.
	Special bool
	// Rotation defines the rotation options for the resource.
	Rotation *rModel.Options
}

// CreatePassword creates a RandomPassword resource and returns PasswordData.
// Defaults: length=16, special=true.
// ctx: Pulumi context.
// name: Name prefix for the resource.
// opts: Optional parameters for password generation.
func CreatePassword(ctx *pulumi.Context, name string, opts *PasswordOptions) (*random.PasswordData, error) {
	pulumiOpts := []pulumi.ResourceOption{}
	if opts != nil {
		if trigger, _ := rotation.Trigger(ctx, name, opts.Rotation); trigger != nil {
			pulumiOpts = append(pulumiOpts, pulumi.ReplacementTrigger(trigger))
		}
	}

	length := 16
	special := true
	if opts != nil {
		if opts.Length != 0 {
			length = opts.Length
		}
		special = opts.Special
	}

	pw, err := randomsdk.NewRandomPassword(ctx, name, &randomsdk.RandomPasswordArgs{
		Length:  pulumi.Int(length),
		Special: pulumi.Bool(special),
		Lower:   pulumi.Bool(true),
		Upper:   pulumi.Bool(true),
		Number:  pulumi.Bool(true),
	},
		pulumiOpts...)
	if err != nil {
		return nil, err
	}

	return &random.PasswordData{
		Resource: pw,
		Password: pw.Result,
	}, nil
}
