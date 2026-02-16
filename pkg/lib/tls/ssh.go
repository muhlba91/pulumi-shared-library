//nolint:revive // package name is fine as is
package tls

import (
	"fmt"

	tls "github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateSSHKey creates a new SSH private key with the specified number of bits.
// ctx: The Pulumi context.
// name: The name to use for the key resource.
// bits: The number of bits for the SSH key. Defaults to 4096 if set to 0.
func CreateSSHKey(ctx *pulumi.Context, name string, bits int) (*tls.PrivateKey, error) {
	return createKey(ctx, fmt.Sprintf("ssh-key-%s", name), bits)
}
