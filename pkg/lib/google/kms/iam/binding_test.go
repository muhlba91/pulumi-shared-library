package iam_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/google/kms/iam"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateKeyringBinding(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		keyRingID := "projects/proj/locations/global/keyRings/ring"
		member := "user:alice@example.com"
		role := "roles/cloudkms.cryptoKeyEncrypterDecrypter"

		args := &iam.KeyringBindingArgs{
			KeyRingID: keyRingID,
			Member:    member,
			Role:      role,
		}

		res, err := iam.CreateKeyringBinding(ctx, args)
		require.NoError(t, err)
		require.NotNil(t, res)

		res.KeyRingId.ApplyT(func(id string) error {
			assert.Equal(t, keyRingID, id)
			return nil
		})
		res.Members.ApplyT(func(m []string) error {
			assert.Equal(t, []string{member}, m)
			return nil
		})
		res.Role.ApplyT(func(r string) error {
			assert.Equal(t, role, r)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateKeyringBinding_WithOptionalArgs(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		keyRingID := "projects/proj/locations/global/keyRings/ring"
		member := "user:alice@example.com"
		role := "roles/cloudkms.cryptoKeyEncrypterDecrypter"

		args := &iam.KeyringBindingArgs{
			KeyRingID:     keyRingID,
			Member:        member,
			Role:          role,
			PulumiOptions: []pulumi.ResourceOption{},
		}

		res, err := iam.CreateKeyringBinding(ctx, args)
		require.NoError(t, err)
		require.NotNil(t, res)

		res.KeyRingId.ApplyT(func(id string) error {
			assert.Equal(t, keyRingID, id)
			return nil
		})
		res.Members.ApplyT(func(m []string) error {
			assert.Equal(t, []string{member}, m)
			return nil
		})
		res.Role.ApplyT(func(r string) error {
			assert.Equal(t, role, r)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
