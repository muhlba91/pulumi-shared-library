package iam

import (
	"fmt"

	gstorage "github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

// HmacKeyOptions represents the options for creating an HMAC key.
type HmacKeyOptions struct {
	// ServiceAccount is the email of the service account for which to create the HMAC key.
	ServiceAccount string
	// Project is the GCP project ID where the HMAC key will be created. Optional.
	Project pulumi.StringPtrInput
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// CreateHmacKey creates an HMAC key for a Google service account.
// ctx: Pulumi context.
// opts: The options for creating the HMAC key.
func CreateHmacKey(ctx *pulumi.Context, opts *HmacKeyOptions) (*gstorage.HmacKey, error) {
	return gstorage.NewHmacKey(ctx,
		fmt.Sprintf("gcp-hmac-%s", sanitize.Text(opts.ServiceAccount)),
		&gstorage.HmacKeyArgs{
			ServiceAccountEmail: pulumi.String(opts.ServiceAccount),
			Project:             opts.Project,
		},
		opts.PulumiOptions...,
	)
}
