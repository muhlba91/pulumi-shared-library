package file

import (
	"os"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rs/zerolog/log"
)

const defaultPermissions os.FileMode = 0o644

// WriteContents writes content to the given path with the provided permissions.
// path: file path to write to.
// content: content to write.
// permissions: file permissions (os.FileMode, default is 0644), optional.
func WriteContents(path string, content string, permissions ...os.FileMode) (string, error) {
	perm := defaultPermissions
	if len(permissions) > 0 {
		perm = permissions[0]
	}

	err := os.WriteFile(path, []byte(content), perm)
	return content, err
}

// WritePulumi writes a Pulumi StringInput to a file and returns a pulumi.StringOutput
// that resolves to the written content. Optional permissions can be passed; default is 0644.
// path: file path to write to.
// content: Pulumi StringInput content to write.
// permissions: optional file permissions (os.FileMode).
func WritePulumi(path string, content pulumi.StringInput, permissions ...os.FileMode) pulumi.StringOutput {
	return content.ToStringOutput().ApplyT(func(v string) string { //nolint:errcheck // errors are logged
		_, err := WriteContents(path, v, permissions...)
		if err != nil {
			log.Error().Msgf("failed to write file %s: %v", path, err)
		}

		return v
	}).(pulumi.StringOutput)
}
