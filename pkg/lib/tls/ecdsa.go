//nolint:revive,nolintlint // package name is fine as is
package tls

import (
	tls "github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateECDSAKey creates a new ECDSA private key on the specified curve.
// ctx: The Pulumi context.
// name: The name to use for the key resource (prefixed with "ecdsa-key-").
// curve: The curve of the ECDSA key (P224, P256, P384, P521). Defaults to P256 if empty.
func CreateECDSAKey(ctx *pulumi.Context, name string, curve string) (*tls.PrivateKey, error) {
	opts := &CreateOptions{Algorithm: AlgorithmECDSA}
	if curve != "" {
		opts.ECDSACurve = &curve
	}

	return Create(ctx, name, opts)
}
