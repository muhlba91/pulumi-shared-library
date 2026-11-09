package approle_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libapprole "github.com/muhlba91/pulumi-shared-library/pkg/lib/vault/approle"
	rModel "github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateSecretID(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		s, err := libapprole.CreateSecretID(ctx, "dev-host", &libapprole.SecretIDOptions{
			RoleName: pulumi.String("dev-host"),
		})
		require.NoError(t, err)
		require.NotNil(t, s)

		pulumi.All(s.Backend, s.RoleName).ApplyT(func(all []any) error {
			assert.Equal(t, "approle", *all[0].(*string))
			assert.Equal(t, "dev-host", all[1].(string))
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Equal(t,
		[]string{"vault-approle-secret-id-dev-host"},
		counter.Resources["vault:appRole/authBackendRoleSecretId:AuthBackendRoleSecretId"])
	assert.Empty(t, counter.Resources["time:index/rotating:Rotating"])
}

func TestCreateSecretID_WithRotation(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		_, err := libapprole.CreateSecretID(ctx, "dev-host", &libapprole.SecretIDOptions{
			RoleName: pulumi.String("dev-host"),
			Metadata: pulumi.String(`{"host":"dev-host"}`),
			Rotation: &rModel.Options{Days: 90},
		})
		require.NoError(t, err)
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Len(t, counter.Resources["time:index/rotating:Rotating"], 1)
}
