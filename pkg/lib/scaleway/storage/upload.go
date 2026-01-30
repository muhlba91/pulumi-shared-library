package storage

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/object"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/metadata"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

// UploadArgs holds optional parameters for uploading to Scaleway Object Storage.
type UploadArgs struct {
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
// name: a name prefix for the resource.
// args: optional arguments for the upload and metadata.
func Upload(
	ctx *pulumi.Context,
	args *UploadArgs,
) (*object.Item, error) {
	resName := fmt.Sprintf("scaleway-object-%s-%s", args.BucketID, sanitize.Text(args.Key))

	var source pulumi.StringPtrInput
	if args.File != nil {
		source = pulumi.StringPtr(*args.File)
	}

	var content pulumi.StringPtrInput
	if args.Content != nil {
		content = pulumi.StringPtr(*args.Content)
	}

	labels := metadata.LabelsToStringMap(args.Labels)

	return object.NewItem(ctx, resName, &object.ItemArgs{
		Bucket:       pulumi.String(args.BucketID),
		Key:          pulumi.String(args.Key),
		File:         source,
		Content:      content,
		Metadata:     labels,
		StorageClass: pulumi.String("STANDARD"),
		Visibility:   pulumi.String("private"),
	}, args.PulumiOptions...)
}
