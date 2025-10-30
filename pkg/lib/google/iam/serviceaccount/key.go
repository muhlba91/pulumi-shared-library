package serviceaccount

import (
	"fmt"

	gsa "github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// KeyArgs represents the arguments for creating a service account key.
type KeyArgs struct {
	// ServiceAccount is the service account ID to create the key for.
	ServiceAccount pulumi.StringInput
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// CreateKey creates a new service account key.
// ctx: Pulumi context.
// name: name suffix for the key resource.
// args: KeyArgs containing the service account ID and optional project ID.
func CreateKey(
	ctx *pulumi.Context,
	name string,
	args *KeyArgs,
) (*gsa.Key, error) {
	return gsa.NewKey(ctx,
		fmt.Sprintf("gcp-sa-key-%s", name),
		&gsa.KeyArgs{
			ServiceAccountId: args.ServiceAccount,
		},
		args.PulumiOptions...,
	)
}
