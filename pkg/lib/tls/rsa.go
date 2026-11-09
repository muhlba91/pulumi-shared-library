//nolint:revive,nolintlint // package name is fine as is
package tls

import (
	tls "github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateRSAKey creates a new RSA private key with the specified number of bits.
// ctx: The Pulumi context.
// name: The name to use for the key resource (prefixed with "rsa-key-").
// bits: The number of bits for the RSA key. Defaults to 4096 if set to 0.
func CreateRSAKey(ctx *pulumi.Context, name string, bits int) (*tls.PrivateKey, error) {
	opts := &CreateOptions{Algorithm: AlgorithmRSA}
	if bits != 0 {
		opts.RSABits = &bits
	}

	return Create(ctx, name, opts)
}
