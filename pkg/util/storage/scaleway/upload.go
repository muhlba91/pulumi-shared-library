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
// The object is registered immediately with the (possibly secret) content as an input, so secrecy is preserved,
// previews show the object, and a failed content or a failed local write fails the program.
// It returns a Pulumi Output that resolves to the created Item once the local file has been written.
// ctx: Pulumi context.
// opts: the options for writing the file and uploading it.
func WriteFileAndUpload(
	ctx *pulumi.Context,
	opts *storage.WriteFileAndUploadOptions,
) pulumi.Output {
	written := fileutil.WritePulumiStrict(filepath.Join(opts.OutputPath, opts.Name), opts.Content, opts.Permissions...)

	item, err := scwutil.Upload(ctx, &scwutil.UploadOptions{
		BucketID:     opts.BucketID,
		ContentInput: written,
		Key:          path.Join(opts.BucketPath, opts.Name),
		Labels:       opts.Labels,
	})
	if err != nil {
		log.Error().Msgf("Failed to upload object to Scaleway: %v", err)
		return pulumi.ToOutput(nil)
	}

	return written.ApplyT(func(_ string) *object.Item { return item })
}
