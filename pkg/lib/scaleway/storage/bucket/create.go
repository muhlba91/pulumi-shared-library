package bucket

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/object"
)

// CreateOptions holds optional parameters for Create.
type CreateOptions struct {
	// Location is the Scaleway region where the bucket will be created.
	Location pulumi.StringInput
	// Labels are optional key/value pairs to tag the bucket.
	Labels map[string]string
	// PulumiOptions are optional resource options passed to the RecordSet.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a Scaleway bucket with the given parameters.
// ctx: Pulumi context.
// name: Name for the bucket.
// opts: CreateOptions with parameters for the bucket.
func Create(
	ctx *pulumi.Context,
	name string,
	opts *CreateOptions,
) (*object.Bucket, error) {
	return object.NewBucket(ctx, fmt.Sprintf("scaleway-bucket-%s", name), &object.BucketArgs{
		Region:       opts.Location,
		ForceDestroy: pulumi.Bool(true),
		Tags:         pulumi.ToStringMap(opts.Labels),
	}, opts.PulumiOptions...)
}
