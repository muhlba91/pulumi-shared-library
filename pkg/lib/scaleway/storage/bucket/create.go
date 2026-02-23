package bucket

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/object"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
)

// incompleteMultipartUploadAbortDays is the number of days after which incomplete multipart uploads will be aborted.
const incompleteMultipartUploadAbortDays = 3

// defaultOneZoneTransitionDays is the number of days after which objects will be transitioned to the ONEZONE storage class.
const defaultOneZoneTransitionDays = 3

// CreateOptions holds optional parameters for Create.
type CreateOptions struct {
	// Location is the Scaleway region where the bucket will be created.
	Location pulumi.StringInput
	// CORs configuration for the bucket.
	CORS *CreateCorsOptions
	// OneZoneTransitionDays is the number of days after which objects will be transitioned to the ONEZONE storage class.
	OneZoneTransitionDays *int
	// Labels are optional key/value pairs to tag the bucket.
	Labels map[string]string
	// PulumiOptions are optional resource options passed to the RecordSet.
	PulumiOptions []pulumi.ResourceOption
}

// CreateCorsOptions defines CORS configuration data for a Scaleway bucket.
type CreateCorsOptions struct {
	// MaxAgeSeconds is the maximum age of the CORS preflight request.
	MaxAgeSeconds *int
	// Method is the list of allowed HTTP methods.
	Method []string
	// Origin is the list of allowed origins.
	Origin []string
	// ResponseHeader is the list of allowed response headers.
	ResponseHeader []string
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
	args := &object.BucketArgs{
		Region:       opts.Location,
		ForceDestroy: pulumi.Bool(true),
		LifecycleRules: &object.BucketLifecycleRuleArray{
			&object.BucketLifecycleRuleArgs{
				Enabled:                            pulumi.Bool(true),
				Prefix:                             pulumi.String("expire-incomplete-multipart-uploads"),
				AbortIncompleteMultipartUploadDays: pulumi.Int(incompleteMultipartUploadAbortDays),
			},
			&object.BucketLifecycleRuleArgs{
				Enabled: pulumi.Bool(true),
				Prefix:  pulumi.String("move-to-one-zone"),
				Transitions: &object.BucketLifecycleRuleTransitionArray{
					&object.BucketLifecycleRuleTransitionArgs{
						StorageClass: pulumi.String("ONEZONE_IA"),
						Days: pulumi.Int(
							defaults.GetOrDefault(opts.OneZoneTransitionDays, defaultOneZoneTransitionDays),
						),
					},
				},
			},
		},
		Tags: pulumi.ToStringMap(opts.Labels),
	}

	if opts.CORS != nil {
		args.CorsRules = &object.BucketCorsRuleArray{
			&object.BucketCorsRuleArgs{
				MaxAgeSeconds:  pulumi.IntPtrFromPtr(opts.CORS.MaxAgeSeconds),
				AllowedMethods: pulumi.ToStringArray(opts.CORS.Method),
				AllowedOrigins: pulumi.ToStringArray(opts.CORS.Origin),
				AllowedHeaders: pulumi.ToStringArray(opts.CORS.ResponseHeader),
			},
		}
	}

	return object.NewBucket(ctx, fmt.Sprintf("scaleway-bucket-%s", name), args, opts.PulumiOptions...)
}
