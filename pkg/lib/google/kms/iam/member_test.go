package iam_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/google/kms/iam"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateKMSMember(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		cryptoKeyID := "projects/proj/locations/global/keyRings/ring/cryptoKeys/key"
		member := "user:alice@example.com"
		role := "roles/cloudkms.cryptoKeyEncrypterDecrypter"

		args := &iam.MemberArgs{
			CryptoKeyID: cryptoKeyID,
			Member:      member,
			Role:        role,
		}

		res, err := iam.CreateMember(ctx, args)
		require.NoError(t, err)
		require.NotNil(t, res)

		res.CryptoKeyId.ApplyT(func(id string) error {
			assert.Equal(t, cryptoKeyID, id)
			return nil
		})
		res.Member.ApplyT(func(m string) error {
			assert.Equal(t, member, m)
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

func TestCreateKMSMember_WithOptionalArgs(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		cryptoKeyID := "projects/proj/locations/global/keyRings/ring/cryptoKeys/key"
		member := "user:alice@example.com"
		role := "roles/cloudkms.cryptoKeyEncrypterDecrypter"

		args := &iam.MemberArgs{
			CryptoKeyID:   cryptoKeyID,
			Member:        member,
			Role:          role,
			PulumiOptions: []pulumi.ResourceOption{},
		}

		res, err := iam.CreateMember(ctx, args)
		require.NoError(t, err)
		require.NotNil(t, res)

		res.CryptoKeyId.ApplyT(func(id string) error {
			assert.Equal(t, cryptoKeyID, id)
			return nil
		})
		res.Member.ApplyT(func(m string) error {
			assert.Equal(t, member, m)
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

func TestCreateKeyringMember(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		keyRingID := "projects/proj/locations/global/keyRings/ring"
		member := "user:alice@example.com"
		role := "roles/cloudkms.cryptoKeyEncrypterDecrypter"

		args := &iam.KeyringMemberArgs{
			KeyRingID: keyRingID,
			Member:    member,
			Role:      role,
		}

		res, err := iam.CreateKeyringMember(ctx, args)
		require.NoError(t, err)
		require.NotNil(t, res)

		res.KeyRingId.ApplyT(func(id string) error {
			assert.Equal(t, keyRingID, id)
			return nil
		})
		res.Member.ApplyT(func(m string) error {
			assert.Equal(t, member, m)
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

func TestCreateKeyringMember_WithOptionalArgs(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		keyRingID := "projects/proj/locations/global/keyRings/ring"
		member := "user:alice@example.com"
		role := "roles/cloudkms.cryptoKeyEncrypterDecrypter"

		args := &iam.KeyringMemberArgs{
			KeyRingID:     keyRingID,
			Member:        member,
			Role:          role,
			PulumiOptions: []pulumi.ResourceOption{},
		}

		res, err := iam.CreateKeyringMember(ctx, args)
		require.NoError(t, err)
		require.NotNil(t, res)

		res.KeyRingId.ApplyT(func(id string) error {
			assert.Equal(t, keyRingID, id)
			return nil
		})
		res.Member.ApplyT(func(m string) error {
			assert.Equal(t, member, m)
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
