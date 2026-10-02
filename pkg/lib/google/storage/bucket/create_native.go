package bucket

import (
	"fmt"

	storage "github.com/pulumi/pulumi-google-native/sdk/go/google/storage/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateNativeOptions defines the options for creating a GCS bucket with the native provider.
type CreateNativeOptions struct {
	// Location is the GCP location (region or multi-region) where the bucket will be created.
	Location pulumi.StringInput
	// CORS is the CORS configuration for the bucket. Optional, no CORS is configured if nil.
	CORS *CreateNativeCorsOptions
	// Labels are optional key/value pairs to tag the bucket.
	Labels map[string]string
	// PulumiOptions are optional resource options passed to the Bucket.
	PulumiOptions []pulumi.ResourceOption
}

// CreateNativeCorsOptions defines CORS configuration data for a GCS bucket.
type CreateNativeCorsOptions struct {
	// MaxAgeSeconds is the time in seconds browsers may cache the response to a CORS preflight request.
	MaxAgeSeconds *int
	// Method is the list of allowed HTTP methods.
	Method []string
	// Origin is the list of allowed origins.
	Origin []string
	// ResponseHeader is the list of allowed response headers.
	ResponseHeader []string
}

// CreateNative creates a GCS bucket with the given parameters using the native provider.
// The bucket uses the STANDARD storage class.
// ctx: Pulumi context.
// name: The logical name for the Pulumi resource (prefixed with "gcs-bucket-").
// opts: CreateNativeOptions with parameters for the bucket.
func CreateNative(
	ctx *pulumi.Context,
	name string,
	opts *CreateNativeOptions,
) (*storage.Bucket, error) {
	args := &storage.BucketArgs{
		Location:     opts.Location,
		StorageClass: pulumi.String("STANDARD"),
		Labels:       pulumi.ToStringMap(opts.Labels),
	}

	if opts.CORS != nil {
		args.Cors = &storage.BucketCorsItemArray{
			&storage.BucketCorsItemArgs{
				MaxAgeSeconds:  pulumi.IntPtrFromPtr(opts.CORS.MaxAgeSeconds),
				Method:         pulumi.ToStringArray(opts.CORS.Method),
				Origin:         pulumi.ToStringArray(opts.CORS.Origin),
				ResponseHeader: pulumi.ToStringArray(opts.CORS.ResponseHeader),
			},
		}
	}

	return storage.NewBucket(ctx, fmt.Sprintf("gcs-bucket-%s", name), args, opts.PulumiOptions...)
}
