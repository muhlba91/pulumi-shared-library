package random

import (
	randomsdk "github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// PasswordData is a helper struct to bundle password resource and its output.
type PasswordData struct {
	// Resource is the password resource.
	Resource *randomsdk.RandomPassword
	// Password is the output of the password.
	Password pulumi.StringOutput
}
