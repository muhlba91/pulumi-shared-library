package store_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libstore "github.com/muhlba91/pulumi-shared-library/pkg/lib/vault/store"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateVaultStore(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		args := &libstore.CreateArgs{
			Path:        pulumi.String("secret/path"),
			Description: pulumi.String("test store"),
		}

		mnt, err := libstore.Create(ctx, "basic", args)
		require.NoError(t, err)
		require.NotNil(t, mnt)

		mnt.ID().ApplyT(func(p string) error {
			assert.Equal(t, "vault-store-basic_id", p)
			return nil
		})
		mnt.Path.ApplyT(func(p string) error {
			assert.Equal(t, "secret/path", p)
			return nil
		})
		mnt.Type.ApplyT(func(tp string) error {
			assert.Equal(t, "kv", tp)
			return nil
		})
		mnt.Options.ApplyT(func(o map[string]string) error {
			assert.Equal(t, "2", o["version"])
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateVaultStore_WithOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		prefix := "mount"
		args := &libstore.CreateArgs{
			Path:          pulumi.String("secret/withopts"),
			Description:   pulumi.String("store with options"),
			NamePrefix:    &prefix,
			PulumiOptions: []pulumi.ResourceOption{pulumi.Protect(true)},
		}

		mnt, err := libstore.Create(ctx, "withopts", args)
		require.NoError(t, err)
		require.NotNil(t, mnt)

		mnt.ID().ApplyT(func(p string) error {
			assert.Equal(t, "vault-mount-withopts_id", p)
			return nil
		})
		mnt.Path.ApplyT(func(p string) error {
			assert.Equal(t, "secret/withopts", p)
			return nil
		})
		mnt.Type.ApplyT(func(tp string) error {
			assert.Equal(t, "kv", tp)
			return nil
		})
		mnt.Options.ApplyT(func(o map[string]string) error {
			assert.Equal(t, "2", o["version"])
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
