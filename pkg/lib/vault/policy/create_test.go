package policy_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libpolicy "github.com/muhlba91/pulumi-shared-library/pkg/lib/vault/policy"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateVaultPolicy(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		args := &libpolicy.CreateArgs{
			Name:   "test-policy",
			Policy: pulumi.String(`path "secret/*" { capabilities = ["read","list"] }`),
		}

		p, err := libpolicy.Create(ctx, args)
		require.NoError(t, err)
		require.NotNil(t, p)

		p.Name.ApplyT(func(n string) error {
			assert.Equal(t, "test-policy", n)
			return nil
		})
		p.Policy.ApplyT(func(pol string) error {
			assert.Contains(t, pol, "path")
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateVaultPolicy_WithOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		args := &libpolicy.CreateArgs{
			Name:          "protected-policy",
			Policy:        pulumi.String(`path "secret/prod/*" { capabilities = ["create","read"] }`),
			PulumiOptions: []pulumi.ResourceOption{pulumi.Protect(true)},
		}

		p, err := libpolicy.Create(ctx, args)
		require.NoError(t, err)
		require.NotNil(t, p)

		p.Name.ApplyT(func(n string) error {
			assert.Equal(t, "protected-policy", n)
			return nil
		})
		p.Policy.ApplyT(func(pol string) error {
			assert.Contains(t, pol, "secret/prod")
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
