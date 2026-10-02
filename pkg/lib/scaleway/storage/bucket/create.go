package bucket

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/object"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
)

// incompleteMultipartUploadAbortDays is the number of days after which incomplete multipart uploads will be aborted.
const incompleteMultipartUploadAbortDays = 3

// defaultOneZoneTransitionDays is the default number of days after which objects will be transitioned to ONEZONE_IA.
const defaultOneZoneTransitionDays = 30

// CreateOptions defines the options for creating a Scaleway bucket.
type CreateOptions struct {
	// Location is the Scaleway region where the bucket will be created.
	Location pulumi.StringInput
	// CORS is the CORS configuration for the bucket. Optional, no CORS is configured if nil.
	CORS *CreateCorsOptions
	// OneZoneTransitionDays is the number of days after which objects will be transitioned to the ONEZONE_IA storage class.
	// Optional, defaults to 30.
	OneZoneTransitionDays *int
	// Labels are optional key/value pairs to tag the bucket.
	Labels map[string]string
	// PulumiOptions are optional resource options passed to the Bucket.
	PulumiOptions []pulumi.ResourceOption
}

// CreateCorsOptions defines CORS configuration data for a Scaleway bucket.
type CreateCorsOptions struct {
	// MaxAgeSeconds is the time in seconds browsers may cache the response to a CORS preflight request.
	MaxAgeSeconds *int
	// Method is the list of allowed HTTP methods.
	Method []string
	// Origin is the list of allowed origins.
	Origin []string
	// ResponseHeader is the list of allowed headers.
	ResponseHeader []string
}

// Create creates a Scaleway bucket with the given parameters.
// Incomplete multipart uploads are aborted after 3 days, objects are transitioned to ONEZONE_IA,
// and the objects are deleted with the bucket.
// ctx: Pulumi context.
// name: The logical name for the Pulumi resource (prefixed with "scaleway-bucket-").
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
