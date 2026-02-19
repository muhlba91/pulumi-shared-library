package bucket

import (
	"fmt"

	storage "github.com/pulumi/pulumi-google-native/sdk/go/google/storage/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateNativeOptions holds optional parameters for Create.
type CreateNativeOptions struct {
	// Location is the GCP region where the bucket will be created.
	Location pulumi.StringInput
	// CORs configuration for the bucket.
	CORS *CreateNativeCorsOptions
	// Labels are optional key/value pairs to tag the bucket.
	Labels map[string]string
	// PulumiOptions are optional resource options passed to the RecordSet.
	PulumiOptions []pulumi.ResourceOption
}

// CreateNativeCorsOptions defines CORS configuration data for an GCS bucket.
type CreateNativeCorsOptions struct {
	// MaxAgeSeconds is the maximum age of the CORS preflight request.
	MaxAgeSeconds *int
	// Method is the list of allowed HTTP methods.
	Method []string
	// Origin is the list of allowed origins.
	Origin []string
	// ResponseHeader is the list of allowed response headers.
	ResponseHeader []string
}

// CreateNative creates a GCS bucket with the given parameters.
// ctx: Pulumi context.
// name: Name for the bucket.
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
