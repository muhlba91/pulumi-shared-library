package random

import (
	randomsdk "github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// StringData is a helper struct to bundle string resource and its output.
type StringData struct {
	// Resource is the string resource.
	Resource *randomsdk.RandomString
	// Text is the output of the string.
	Text pulumi.StringOutput
}
