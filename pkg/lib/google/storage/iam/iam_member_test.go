package iam_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/google/storage/iam"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateIAMMember(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		bucket := "my-bucket"
		member := "user:alice@example.com"
		role := "roles/storage.objectViewer"

		opts := &iam.MemberOptions{
			BucketID: bucket,
			Member:   member,
			Role:     role,
		}

		res, err := iam.CreateIAMMember(ctx, opts)
		require.NoError(t, err)
		require.NotNil(t, res)

		res.Bucket.ApplyT(func(b string) error {
			assert.Equal(t, bucket, b)
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

func TestCreateIAMMember_WithOptionalArgs(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		bucket := "my-bucket"
		member := "user:alice@example.com"
		role := "roles/storage.objectViewer"

		opts := &iam.MemberOptions{
			BucketID:      bucket,
			Member:        member,
			Role:          role,
			PulumiOptions: []pulumi.ResourceOption{},
		}

		res, err := iam.CreateIAMMember(ctx, opts)
		require.NoError(t, err)
		require.NotNil(t, res)

		res.Bucket.ApplyT(func(b string) error {
			assert.Equal(t, bucket, b)
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
