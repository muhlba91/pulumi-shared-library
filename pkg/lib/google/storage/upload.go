package storage

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/metadata"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

// UploadOptions holds the options for uploading to GCS.
type UploadOptions struct {
	// Key is the object key in the bucket.
	Key string
	// BucketID is the ID of the GCS bucket.
	BucketID string
	// File is a local path to upload (optional).
	File *string
	// Content is raw content to store (optional).
	Content *string
	// Labels are metadata labels to set on the object (optional).
	Labels map[string]string
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// Upload uploads a file or content to a GCS bucket and returns the BucketObject.
// ctx: Pulumi context.
// opts: UploadOptions containing the upload and metadata options.
func Upload(
	ctx *pulumi.Context,
	opts *UploadOptions,
) (*storage.BucketObject, error) {
	resName := fmt.Sprintf("gcs-object-%s-%s", opts.BucketID, sanitize.Text(opts.Key))

	var source pulumi.AssetOrArchiveInput
	if opts.File != nil {
		source = pulumi.NewFileAsset(*opts.File)
	}

	var content pulumi.StringPtrInput
	if opts.Content != nil {
		content = pulumi.StringPtr(*opts.Content)
	}

	labels := metadata.LabelsToStringMap(opts.Labels)

	return storage.NewBucketObject(ctx, resName, &storage.BucketObjectArgs{
		Bucket:   pulumi.String(opts.BucketID),
		Name:     pulumi.String(opts.Key),
		Source:   source,
		Content:  content,
		Metadata: labels,
	}, opts.PulumiOptions...)
}
