package namespace_test

import (
	"testing"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libns "github.com/muhlba91/pulumi-shared-library/pkg/lib/kubernetes/namespace"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateNamespace(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "testns"

		ns, err := libns.Create(ctx, &libns.CreateOptions{
			Name: name,
		})
		require.NoError(t, err)
		require.NotNil(t, ns)

		// Metadata is metav1.ObjectMetaOutput; check the Name field.
		ns.Metadata.ApplyT(func(m metav1.ObjectMeta) error {
			assert.Equal(t, name, *m.Name)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateNamespace_WithOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "protected-ns"

		ns, err := libns.Create(ctx, &libns.CreateOptions{
			Name: name,
			PulumiOptions: []pulumi.ResourceOption{
				pulumi.Protect(true),
			},
		})
		require.NoError(t, err)
		require.NotNil(t, ns)

		ns.Metadata.ApplyT(func(m metav1.ObjectMeta) error {
			assert.Equal(t, name, *m.Name)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
