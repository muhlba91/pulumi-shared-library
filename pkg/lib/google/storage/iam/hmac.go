package iam

import (
	"fmt"

	gstorage "github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"

	rModel "github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/rotation"
)

// HmacKeyOptions represents the options for creating an HMAC key.
type HmacKeyOptions struct {
	// ServiceAccount is the email of the service account for which to create the HMAC key.
	ServiceAccount string
	// Project is the GCP project ID where the HMAC key will be created. Optional.
	Project pulumi.StringPtrInput
	// Rotation defines the rotation options for the resource.
	Rotation *rModel.Options
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// CreateHmacKey creates an HMAC key for a Google service account.
// ctx: Pulumi context.
// opts: The options for creating the HMAC key.
func CreateHmacKey(ctx *pulumi.Context, opts *HmacKeyOptions) (*gstorage.HmacKey, error) {
	resName := fmt.Sprintf("gcp-hmac-%s", sanitize.Text(opts.ServiceAccount))

	pulumiOpts := append([]pulumi.ResourceOption{}, opts.PulumiOptions...)
	if trigger, _ := rotation.Trigger(ctx, resName, opts.Rotation); trigger != nil {
		pulumiOpts = append(pulumiOpts, pulumi.ReplacementTrigger(trigger))
	}

	return gstorage.NewHmacKey(ctx,
		resName,
		&gstorage.HmacKeyArgs{
			ServiceAccountEmail: pulumi.String(opts.ServiceAccount),
			Project:             opts.Project,
		},
		pulumiOpts...,
	)
}
