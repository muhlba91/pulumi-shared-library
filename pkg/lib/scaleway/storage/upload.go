package storage

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/object"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/metadata"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

// UploadOptions holds the options for uploading to Scaleway Object Storage.
type UploadOptions struct {
	// Key is the object key in the bucket.
	Key string
	// BucketID is the ID of the Scaleway bucket.
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

// Upload uploads a file or content to a Scaleway bucket and returns the Item.
// ctx: Pulumi context.
// opts: UploadOptions containing the upload and metadata options.
func Upload(
	ctx *pulumi.Context,
	opts *UploadOptions,
) (*object.Item, error) {
	resName := fmt.Sprintf("scaleway-object-%s-%s", opts.BucketID, sanitize.Text(opts.Key))

	var source pulumi.StringPtrInput
	if opts.File != nil {
		source = pulumi.StringPtr(*opts.File)
	}

	var content pulumi.StringPtrInput
	if opts.Content != nil {
		content = pulumi.StringPtr(*opts.Content)
	}

	labels := metadata.LabelsToStringMap(opts.Labels)

	return object.NewItem(ctx, resName, &object.ItemArgs{
		Bucket:       pulumi.String(opts.BucketID),
		Key:          pulumi.String(opts.Key),
		File:         source,
		Content:      content,
		Metadata:     labels,
		StorageClass: pulumi.String("STANDARD"),
		Visibility:   pulumi.String("private"),
	}, opts.PulumiOptions...)
}
