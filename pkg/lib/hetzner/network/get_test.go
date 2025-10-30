package network_test

// import (
// 	"testing"

// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/rs/zerolog/log"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"

// 	libnet "github.com/muhlba91/pulumi-shared-library/pkg/lib/hetzner/network"
// 	"github.com/muhlba91/pulumi-shared-library/test/mocks"
// )

// // FIXME: mock does not work
// // func TestGetNetwork(t *testing.T) {
// // 	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
// // 		name := "network"

// // 		res, err := libnet.Get(ctx, name)
// // 		log.Debug().Interface("res", res).Msg("Get Network result")
// // 		require.NoError(t, err)
// // 		assert.NotNil(t, res)

// // 		assert.Equal(t, name, res.Name)
// // 		assert.NotEmpty(t, res.Id)
// // 		return nil
// // 	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
// // 	require.NoError(t, err)
// // }
