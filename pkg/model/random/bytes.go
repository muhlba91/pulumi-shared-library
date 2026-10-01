package random

import (
	randomsdk "github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// BytesData is a helper struct to bundle bytes resource and its outputs.
type BytesData struct {
	// Resource is the bytes resource.
	Resource *randomsdk.RandomId
	// Hex is the output of the bytes as hexadecimal digits.
	Hex pulumi.StringOutput
	// Base64Std is the output of the bytes as standard base64.
	Base64Std pulumi.StringOutput
	// Base64URL is the output of the bytes as URL-friendly base64.
	Base64URL pulumi.StringOutput
}
