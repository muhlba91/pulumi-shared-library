package serviceaccount

import (
	"fmt"

	gsa "github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	rModel "github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/rotation"
)

// KeyOptions represents the options for creating a service account key.
type KeyOptions struct {
	// ServiceAccount is the service account ID to create the key for.
	ServiceAccount pulumi.StringInput
	// Rotation defines the rotation options for the resource.
	Rotation *rModel.Options
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// CreateKey creates a new service account key.
// ctx: Pulumi context.
// name: name suffix for the key resource.
// opts: KeyOptions containing the service account ID and optional project ID.
func CreateKey(
	ctx *pulumi.Context,
	name string,
	opts *KeyOptions,
) (*gsa.Key, error) {
	resName := fmt.Sprintf("gcp-sa-key-%s", name)

	pulumiOpts := append([]pulumi.ResourceOption{}, opts.PulumiOptions...)
	if trigger, _ := rotation.Trigger(ctx, resName, opts.Rotation); trigger != nil {
		pulumiOpts = append(pulumiOpts, pulumi.ReplacementTrigger(trigger))
	}

	return gsa.NewKey(ctx,
		resName,
		&gsa.KeyArgs{
			ServiceAccountId: opts.ServiceAccount,
		},
		pulumiOpts...,
	)
}
