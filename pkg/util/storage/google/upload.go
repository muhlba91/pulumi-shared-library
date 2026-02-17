package google

import (
	"path"
	"path/filepath"

	gstorage "github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rs/zerolog/log"

	gcsutil "github.com/muhlba91/pulumi-shared-library/pkg/lib/google/storage"
	fileutil "github.com/muhlba91/pulumi-shared-library/pkg/util/file"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/storage"
)

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
	opts *storage.WriteFileAndUploadOptions,
) pulumi.Output {
	written := fileutil.WritePulumi(filepath.Join(opts.OutputPath, opts.Name), opts.Content, opts.Permissions...)

	return written.ApplyT(func(v string) *gstorage.BucketObject {
		bo, err := gcsutil.Upload(ctx, &gcsutil.UploadOptions{
			BucketID: opts.BucketID,
			Content:  &v,
			Key:      path.Join(opts.BucketPath, opts.Name),
			Labels:   opts.Labels,
		})
		if err != nil {
			log.Error().Msgf("Failed to upload object to GCS: %v", err)
			return nil
		}

		return bo
	})
}
