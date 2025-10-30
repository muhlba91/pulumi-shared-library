package namespace

import (
	"fmt"

	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateOptions defines the options for creating a Kubernetes Namespace.
type CreateOptions struct {
	// Name is the name of the Kubernetes Namespace to create.
	Name string
	// PulumiOptions are additional options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a Kubernetes Namespace with the given name using the provided provider.
// ctx: The Pulumi context.
// opts: The options for creating the Namespace.
func Create(ctx *pulumi.Context, opts *CreateOptions) (*corev1.Namespace, error) {
	return corev1.NewNamespace(ctx, fmt.Sprintf("k8s-namespace-%s", opts.Name), &corev1.NamespaceArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name: pulumi.String(opts.Name),
		},
	}, opts.PulumiOptions...)
}
