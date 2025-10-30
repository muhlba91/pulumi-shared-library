package random

import (
	randomsdk "github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/model/random"
)

// StringOptions holds optional parameters.
type StringOptions struct {
	// Length is the desired length of the generated string.
	Length int
	// Special indicates whether to include special characters in the string.
	Special bool
}

// CreateString creates a RandomString resource and returns StringData.
// Defaults: length=16, special=true.
// ctx: Pulumi context.
// name: Name prefix for the resource.
// opts: Optional parameters for string generation.
func CreateString(ctx *pulumi.Context, name string, opts *StringOptions) (*random.StringData, error) {
	length := 16
	special := true
	if opts != nil {
		if opts.Length != 0 {
			length = opts.Length
		}
		special = opts.Special
	}

	pw, err := randomsdk.NewRandomString(ctx, name, &randomsdk.RandomStringArgs{
		Length:  pulumi.Int(length),
		Special: pulumi.Bool(special),
		Lower:   pulumi.Bool(true),
		Upper:   pulumi.Bool(true),
		Number:  pulumi.Bool(true),
	})
	if err != nil {
		return nil, err
	}

	return &random.StringData{
		Resource: pw,
		Text:     pw.Result,
	}, nil
}
