package bucket

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
)

// CreateOptions holds optional parameters for Create.
type CreateOptions struct {
	// Name prefix for the bucket.
	Name string
	// Prefix for the bucket name.
	Prefix *pulumi.StringPtrInput
	// Labels to apply to the bucket.
	Labels map[string]string
	// Additional Pulumi resource options.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates an S3 bucket with AES256 server-side encryption and standardized tagging.
// ctx: Pulumi context.
// opts: CreateOptions for customizing the bucket creation.
func Create(ctx *pulumi.Context, opts *CreateOptions) (*s3.Bucket, error) {
	b, errCreate := s3.NewBucket(ctx, fmt.Sprintf("s3-bucket-%s", opts.Name), &s3.BucketArgs{
		BucketPrefix: defaults.GetOrDefault(opts.Prefix, pulumi.StringPtrFromPtr(nil)),
		Tags:         pulumi.ToStringMap(opts.Labels),
	}, opts.PulumiOptions...)
	if errCreate != nil {
		return nil, errCreate
	}

	_, errSSE := s3.NewBucketServerSideEncryptionConfiguration(
		ctx,
		fmt.Sprintf("s3-bucket-sse-%s", opts.Name),
		&s3.BucketServerSideEncryptionConfigurationArgs{
			Bucket: b.ID(),
			Rules: s3.BucketServerSideEncryptionConfigurationRuleArray{
				&s3.BucketServerSideEncryptionConfigurationRuleArgs{
					ApplyServerSideEncryptionByDefault: &s3.BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArgs{
						SseAlgorithm: pulumi.String("AES256"),
					},
				},
			},
		},
		opts.PulumiOptions...)
	if errSSE != nil {
		return nil, errSSE
	}

	return b, nil
}
