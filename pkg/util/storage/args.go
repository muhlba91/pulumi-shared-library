package storage

import (
	"os"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// WriteFileAndUploadOptions represents the options for writing a file and uploading it.
type WriteFileAndUploadOptions struct {
	// Name is the name of the object in the bucket, also used as the file name in the output path.
	Name string
	// Content is the content to write and upload.
	Content pulumi.StringInput
	// OutputPath is the local directory to write the content to.
	OutputPath string
	// BucketID is the ID of the bucket.
	BucketID string
	// BucketPath is the path (prefix) within the bucket to upload the object to.
	BucketPath string
	// Labels are the labels to assign to the bucket object.
	Labels map[string]string
	// Permissions are optional file permissions for the written file. Defaults to 0644.
	Permissions []os.FileMode
}
