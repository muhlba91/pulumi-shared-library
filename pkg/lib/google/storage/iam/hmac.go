package iam

import (
	"fmt"

	gstorage "github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

// HmacKeyArgs represents the arguments for creating an HMAC key.
type HmacKeyArgs struct {
	// ServiceAccount is the email of the service account for which to create the HMAC key.
	ServiceAccount string
	// Project is the GCP project ID where the HMAC key will be created. Optional.
	Project pulumi.StringPtrInput
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// CreateHmacKey creates an HMAC key for a Google service account.
// ctx: Pulumi context.
// args: The arguments for creating the HMAC key.
func CreateHmacKey(ctx *pulumi.Context, args *HmacKeyArgs) (*gstorage.HmacKey, error) {
	return gstorage.NewHmacKey(ctx,
		fmt.Sprintf("gcp-hmac-%s", sanitize.Text(args.ServiceAccount)),
		&gstorage.HmacKeyArgs{
			ServiceAccountEmail: pulumi.String(args.ServiceAccount),
			Project:             args.Project,
		},
		args.PulumiOptions...,
	)
}
