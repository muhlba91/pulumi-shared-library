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
// If the upload fails, the error is logged and the output resolves to nil.
// ctx: Pulumi context.
// opts: the options for writing the file and uploading it.
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
