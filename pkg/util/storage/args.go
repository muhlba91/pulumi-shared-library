package storage

import (
	"os"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// WriteFileAndUploadOptions represents the options for writing a file and uploading it.
type WriteFileAndUploadOptions struct {
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
	// Labels are the labels to assign to the bucket object.
	Labels map[string]string
	// Permissions are optional file permissions for the written file.
	Permissions []os.FileMode
}
