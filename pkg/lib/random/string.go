package random

import (
	randomsdk "github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/model/random"
)

// StringOptions holds optional parameters.
// If opts is not nil, Special is used as given, i.e., it is false unless set.
type StringOptions struct {
	// Length is the desired length of the generated string. Defaults to 16 if 0.
	Length int
	// Special indicates whether to include special characters in the string. Defaults to true only if opts is nil.
	Special bool
}

// CreateString creates a RandomString resource and returns StringData.
// Defaults: length=16, special=true.
// ctx: Pulumi context.
// name: Name of the resource.
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
		Numeric: pulumi.Bool(true),
	})
	if err != nil {
		return nil, err
	}

	return &random.StringData{
		Resource: pw,
		Text:     pw.Result,
	}, nil
}
