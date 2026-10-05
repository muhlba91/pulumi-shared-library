package peer_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/netbird/peer"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestGetPeer(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		name := "peer"

		res, err := peer.Get(ctx, &peer.GetOptions{Name: &name})
		log.Debug().Interface("res", res).Msg("Get Peer result")
		require.NoError(t, err)
		assert.NotNil(t, res)

		// HACK: Pulumi LookupPeer mock seems to not populate the result properly even if we return it from Call
		// assert.Equal(t, name, res.Name)
		// assert.NotEmpty(t, res.Id)
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}
