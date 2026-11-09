package approle_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libapprole "github.com/muhlba91/pulumi-shared-library/pkg/lib/vault/approle"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreate_Defaults(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		r, err := libapprole.Create(ctx, "dev-host", &libapprole.CreateOptions{
			TokenPolicies: pulumi.ToStringArray([]string{"dev-host"}),
		})
		require.NoError(t, err)
		require.NotNil(t, r)

		pulumi.All(r.Backend, r.RoleName, r.TokenPolicies, r.TokenTtl, r.TokenMaxTtl, r.SecretIdTtl).
			ApplyT(func(all []any) error {
				assert.Equal(t, "approle", *all[0].(*string))
				assert.Equal(t, "dev-host", all[1].(string))
				assert.Equal(t, []string{"dev-host"}, all[2].([]string))
				assert.Equal(t, 3600, *all[3].(*int))
				assert.Equal(t, 86400, *all[4].(*int))
				assert.Equal(t, 0, *all[5].(*int))
				return nil
			})
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Equal(t,
		[]string{"vault-approle-dev-host"},
		counter.Resources["vault:appRole/authBackendRole:AuthBackendRole"])
}

func TestCreate_WithOptions(t *testing.T) {
	counter := mocks.NewCounter()
	backend := "machines"
	tokenTTL := 600
	tokenMaxTTL := 1200
	secretIDTTL := 3600

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		r, err := libapprole.Create(ctx, "custom", &libapprole.CreateOptions{
			Backend:            &backend,
			TokenPolicies:      pulumi.ToStringArray([]string{"a", "b"}),
			TokenTTL:           &tokenTTL,
			TokenMaxTTL:        &tokenMaxTTL,
			SecretIDTTL:        &secretIDTTL,
			SecretIDBoundCIDRs: pulumi.ToStringArray([]string{"10.0.0.0/8"}),
			TokenBoundCIDRs:    pulumi.ToStringArray([]string{"10.0.0.0/8"}),
		})
		require.NoError(t, err)

		pulumi.All(r.Backend, r.TokenTtl, r.TokenMaxTtl, r.SecretIdTtl, r.SecretIdBoundCidrs).
			ApplyT(func(all []any) error {
				assert.Equal(t, "machines", *all[0].(*string))
				assert.Equal(t, 600, *all[1].(*int))
				assert.Equal(t, 1200, *all[2].(*int))
				assert.Equal(t, 3600, *all[3].(*int))
				assert.Equal(t, []string{"10.0.0.0/8"}, all[4].([]string))
				return nil
			})
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Equal(t,
		[]string{"vault-approle-custom"},
		counter.Resources["vault:appRole/authBackendRole:AuthBackendRole"])
}
