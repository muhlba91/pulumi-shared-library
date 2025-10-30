package tls

import (
	"fmt"

	tls "github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateRSAKey creates a new RSA private key with the specified number of bits.
// ctx: The Pulumi context.
// name: The name to use for the key resource.
// bits: The number of bits for the RSA key. Defaults to 4096 if set to 0.
func CreateRSAKey(ctx *pulumi.Context, name string, bits int) (*tls.PrivateKey, error) {
	return createKey(ctx, fmt.Sprintf("rsa-key-%s", name), bits)
}

func createKey(ctx *pulumi.Context, name string, bits int) (*tls.PrivateKey, error) {
	if bits == 0 {
		bits = 4096
	}

	return tls.NewPrivateKey(ctx, name, &tls.PrivateKeyArgs{
		Algorithm: pulumi.String("RSA"),
		RsaBits:   pulumi.Int(bits),
	})
}
