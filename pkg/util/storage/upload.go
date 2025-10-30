package storage

import (
	"os"
	"path"

	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rs/zerolog/log"

	gcsutil "github.com/muhlba91/pulumi-shared-library/pkg/lib/google/storage"
	fileutil "github.com/muhlba91/pulumi-shared-library/pkg/util/file"
)

// WriteFileAndUploadArgs defines the arguments for WriteFileAndUpload function.
type WriteFileAndUploadArgs struct {
	// Name is the name of the object in the bucket.
	Name string
	// Content is the content to write and upload.
	Content pulumi.StringInput
	// OutputPath is the local file path to write the content to.
	OutputPath string
	// BucketID is the ID of the GCS bucket.
	BucketID string
	// BucketPath is the path within the bucket to upload the object to.
	BucketPath string
	// Permissions are optional file permissions for the written file.
	Permissions []os.FileMode
}

// WriteFileAndUpload writes content to a local file and uploads it to a GCS bucket.
// It returns a Pulumi Output that resolves to the created BucketObject.
// ctx: Pulumi context.
// name: the name of the object in the bucket.
// content: the content to write and upload.
// outputPath: local file path to write the content to.
// bucketID: the ID of the GCS bucket.
// bucketPath: the path within the bucket to upload the object to.
// permissions: optional file permissions for the written file.
func WriteFileAndUpload(
	ctx *pulumi.Context,
	args *WriteFileAndUploadArgs,
) pulumi.Output {
	written := fileutil.WritePulumi(args.OutputPath, args.Content, args.Permissions...)

	return written.ApplyT(func(v string) *storage.BucketObject {
		bo, err := gcsutil.Upload(ctx, &gcsutil.UploadArgs{
			BucketID: args.BucketID,
			Content:  &v,
			Key:      path.Join(args.BucketPath, args.Name),
		})
		if err != nil {
			log.Error().Msgf("Failed to upload object to GCS: %v", err)
			return nil
		}

		return bo
	})
}
