package secret

import (
	"fmt"

	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateOptions defines the options for creating a Kubernetes Secret.
type CreateOptions struct {
	// Name is the name of the Kubernetes Secret to create.
	Name string
	// Namespace is the namespace in which to create the Secret.
	Namespace string
	// Data is the data to store in the Secret.
	Data map[string]pulumi.StringInput
	// PulumiOptions are additional options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a Kubernetes Secret with the provided data and metadata.
// ctx: Pulumi context.
// opts: CreateOptions for customizing the secret creation.
func Create(ctx *pulumi.Context, opts *CreateOptions) (*corev1.Secret, error) {
	return corev1.NewSecret(ctx, fmt.Sprintf("k8s-secret-%s-%s", opts.Namespace, opts.Name), &corev1.SecretArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name:      pulumi.String(opts.Name),
			Namespace: pulumi.String(opts.Namespace),
		},
		Data: pulumi.StringMap(opts.Data),
	}, opts.PulumiOptions...)
}
