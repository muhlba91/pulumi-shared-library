//nolint:revive,nolintlint // package name is fine as is
package tls

import (
	"fmt"
	"strings"

	tls "github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	rModel "github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/rotation"
)

// Algorithm is a private key algorithm supported by the TLS provider.
type Algorithm string

const (
	// AlgorithmRSA creates an RSA key (see CreateOptions.RSABits).
	AlgorithmRSA Algorithm = "RSA"
	// AlgorithmECDSA creates an ECDSA key (see CreateOptions.ECDSACurve).
	AlgorithmECDSA Algorithm = "ECDSA"
	// AlgorithmED25519 creates an Ed25519 key.
	AlgorithmED25519 Algorithm = "ED25519"
)

const (
	// defaultRSABits is the default size of RSA keys.
	defaultRSABits = 4096
	// defaultECDSACurve is the default curve of ECDSA keys.
	defaultECDSACurve = "P256"
)

// CreateOptions defines the options for creating a private key.
type CreateOptions struct {
	// Algorithm is the key algorithm. Required.
	Algorithm Algorithm
	// RSABits is the size of an RSA key. Defaults to 4096. Ignored for other algorithms.
	RSABits *int
	// ECDSACurve is the curve of an ECDSA key (P224, P256, P384, P521). Defaults to P256. Ignored for other algorithms.
	ECDSACurve *string
	// Rotation defines the rotation options for the resource. Optional.
	Rotation *rModel.Options
	// PulumiOptions are additional options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new private key.
// The key is available in PEM (PrivateKeyPem, PublicKeyPem) and OpenSSH (PrivateKeyOpenssh, PublicKeyOpenssh) format.
// Returns an error if the algorithm is not supported.
// ctx: The Pulumi context.
// name: The logical name for the Pulumi resource (prefixed with "<algorithm>-key-", e.g. "ed25519-key-").
// opts: The options for creating the key. Must not be nil.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*tls.PrivateKey, error) {
	return createKey(ctx, fmt.Sprintf("%s-key-%s", strings.ToLower(string(opts.Algorithm)), name), opts)
}

// createKey creates a new private key with the given resource name.
// ctx: The Pulumi context.
// resName: The name of the Pulumi resource.
// opts: The options for creating the key. Must not be nil.
func createKey(ctx *pulumi.Context, resName string, opts *CreateOptions) (*tls.PrivateKey, error) {
	args := &tls.PrivateKeyArgs{
		Algorithm: pulumi.String(string(opts.Algorithm)),
	}

	switch opts.Algorithm {
	case AlgorithmRSA:
		args.RsaBits = pulumi.Int(defaults.GetOrDefault(opts.RSABits, defaultRSABits))
	case AlgorithmECDSA:
		args.EcdsaCurve = pulumi.String(defaults.GetOrDefault(opts.ECDSACurve, defaultECDSACurve))
	case AlgorithmED25519:
	default:
		return nil, fmt.Errorf("unsupported key algorithm: %q", opts.Algorithm)
	}

	pulumiOpts := append([]pulumi.ResourceOption{}, opts.PulumiOptions...)
	if trigger, _ := rotation.Trigger(ctx, resName, opts.Rotation); trigger != nil {
		pulumiOpts = append(pulumiOpts, pulumi.ReplacementTrigger(trigger))
	}

	return tls.NewPrivateKey(ctx, resName, args, pulumiOpts...)
}
