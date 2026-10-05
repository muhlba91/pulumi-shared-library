package network_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/netbird/network"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

const networkTok = "netbird:index/network:Network"

func ptr[T any](v T) *T { return &v }

func TestCreate(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		n, err := network.Create(ctx, "test", &network.CreateOptions{})
		require.NoError(t, err)
		require.NotNil(t, n)

		n.Name.ApplyT(func(v string) error { assert.Equal(t, "test", v); return nil })
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Equal(t, []string{"netbird-network-test"}, counter.Resources[networkTok])
}

func TestCreate_Description(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		n, err := network.Create(ctx, "custom", &network.CreateOptions{Description: ptr("desc")})
		require.NoError(t, err)

		n.Description.ApplyT(func(v string) error { assert.Equal(t, "desc", v); return nil })
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}
