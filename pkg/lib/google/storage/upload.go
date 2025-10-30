package storage

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/metadata"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

// UploadArgs holds optional parameters for uploading to GCS.
type UploadArgs struct {
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
// name: a name prefix for the resource.
// args: optional arguments for the upload and metadata.
func Upload(
	ctx *pulumi.Context,
	args *UploadArgs,
) (*storage.BucketObject, error) {
	resName := fmt.Sprintf("gcs-object-%s-%s", args.BucketID, sanitize.Text(args.Key))

	var source pulumi.AssetOrArchiveInput
	if args.File != nil {
		source = pulumi.NewFileAsset(*args.File)
	}

	var content pulumi.StringPtrInput
	if args.Content != nil {
		content = pulumi.StringPtr(*args.Content)
	}

	labels := metadata.LabelsToStringMap(args.Labels)

	return storage.NewBucketObject(ctx, resName, &storage.BucketObjectArgs{
		Bucket:   pulumi.String(args.BucketID),
		Name:     pulumi.String(args.Key),
		Source:   source,
		Content:  content,
		Metadata: labels,
	}, args.PulumiOptions...)
}
