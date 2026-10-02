package scaleway

import (
	"path"
	"path/filepath"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/object"
	"github.com/rs/zerolog/log"

	scwutil "github.com/muhlba91/pulumi-shared-library/pkg/lib/scaleway/storage"
	fileutil "github.com/muhlba91/pulumi-shared-library/pkg/util/file"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/storage"
)

// WriteFileAndUpload writes content to a local file and uploads it to a Scaleway bucket.
// It returns a Pulumi Output that resolves to the created Item.
// If the upload fails, the error is logged and the output resolves to nil.
// ctx: Pulumi context.
// opts: the options for writing the file and uploading it.
func WriteFileAndUpload(
	ctx *pulumi.Context,
	opts *storage.WriteFileAndUploadOptions,
) pulumi.Output {
	written := fileutil.WritePulumi(filepath.Join(opts.OutputPath, opts.Name), opts.Content, opts.Permissions...)

	return written.ApplyT(func(v string) *object.Item {
		bo, err := scwutil.Upload(ctx, &scwutil.UploadOptions{
			BucketID: opts.BucketID,
			Content:  &v,
			Key:      path.Join(opts.BucketPath, opts.Name),
			Labels:   opts.Labels,
		})
		if err != nil {
			log.Error().Msgf("Failed to upload object to Scaleway: %v", err)
			return nil
		}

		return bo
	})
}
