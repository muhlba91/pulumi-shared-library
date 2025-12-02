package bucket

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateOptions holds optional parameters for Create.
type CreateOptions struct {
	// Location is the GCP region where the bucket will be created.
	Location pulumi.StringInput
	// Labels are optional key/value pairs to tag the bucket.
	Labels map[string]string
	// PulumiOptions are optional resource options passed to the RecordSet.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a GCP bucket with the given parameters.
// ctx: Pulumi context.
// name: Name for the bucket.
// opts: CreateOptions with parameters for the bucket.
func Create(
	ctx *pulumi.Context,
	name string,
	opts *CreateOptions,
) (*storage.Bucket, error) {
	return storage.NewBucket(ctx, fmt.Sprintf("gcp-bucket-%s", name), &storage.BucketArgs{
		Location:                 opts.Location,
		StorageClass:             pulumi.String("STANDARD"),
		PublicAccessPrevention:   pulumi.String("enforced"),
		UniformBucketLevelAccess: pulumi.Bool(true),
		ForceDestroy:             pulumi.Bool(true),
		Labels:                   pulumi.ToStringMap(opts.Labels),
	}, opts.PulumiOptions...)
}
