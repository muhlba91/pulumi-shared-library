package secret_test

import (
	"testing"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libsecret "github.com/muhlba91/pulumi-shared-library/pkg/lib/kubernetes/secret"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateSecret(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "mysecret"
		namespace := "default"

		s, err := libsecret.Create(ctx, &libsecret.CreateOptions{
			Name:      name,
			Namespace: namespace,
			StringData: map[string]pulumi.StringInput{
				"key": pulumi.String("value"),
			},
		})
		require.NoError(t, err)
		require.NotNil(t, s)

		// Check metadata name and namespace
		s.Metadata.ApplyT(func(m metav1.ObjectMeta) error {
			assert.Equal(t, name, *m.Name)
			assert.Equal(t, namespace, *m.Namespace)
			return nil
		})

		// Check stringData contains expected key/value
		s.StringData.ApplyT(func(sd map[string]string) error {
			assert.Equal(t, "value", sd["key"])
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateSecret_WithOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "protected-secret"
		namespace := "ns-1"

		s, err := libsecret.Create(ctx, &libsecret.CreateOptions{
			Name:      name,
			Namespace: namespace,
			Data: map[string]pulumi.StringInput{
				"raw": pulumi.String("rawval"),
			},
			PulumiOptions: []pulumi.ResourceOption{
				pulumi.Protect(true),
			},
		})
		require.NoError(t, err)
		require.NotNil(t, s)

		s.Metadata.ApplyT(func(m metav1.ObjectMeta) error {
			assert.Equal(t, name, *m.Name)
			assert.Equal(t, namespace, *m.Namespace)
			return nil
		})

		// Data is stored as map[string]string on output
		s.Data.ApplyT(func(d map[string]string) error {
			assert.Equal(t, "rawval", d["raw"])
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
