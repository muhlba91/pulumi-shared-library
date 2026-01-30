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
// It returns a Pulumi Output that resolves to the created BucketObject.
// ctx: Pulumi context.
// name: the name of the object in the bucket.
// content: the content to write and upload.
// outputPath: local file path to write the content to.
// bucketID: the ID of the Scaleway bucket.
// bucketPath: the path within the bucket to upload the object to.
// permissions: optional file permissions for the written file.
func WriteFileAndUpload(
	ctx *pulumi.Context,
	args *storage.WriteFileAndUploadArgs,
) pulumi.Output {
	written := fileutil.WritePulumi(filepath.Join(args.OutputPath, args.Name), args.Content, args.Permissions...)

	return written.ApplyT(func(v string) *object.Item {
		bo, err := scwutil.Upload(ctx, &scwutil.UploadArgs{
			BucketID: args.BucketID,
			Content:  &v,
			Key:      path.Join(args.BucketPath, args.Name),
			Labels:   args.Labels,
		})
		if err != nil {
			log.Error().Msgf("Failed to upload object to Scaleway: %v", err)
			return nil
		}

		return bo
	})
}
