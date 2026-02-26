package iam_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/google/storage/iam"
	"github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateHmacKey(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		svc := "service-account@example.com"

		opts := &iam.HmacKeyOptions{
			ServiceAccount: svc,
		}

		hk, err := iam.CreateHmacKey(ctx, opts)
		require.NoError(t, err)
		require.NotNil(t, hk)

		hk.ServiceAccountEmail.ApplyT(func(email string) error {
			assert.Equal(t, svc, email)
			return nil
		})
		hk.Project.ApplyT(func(project string) error {
			assert.Empty(t, project)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Nil(t, counter.Resources["time:index/rotating:Rotating"])
	assert.Len(t, counter.Resources["gcp:storage/hmacKey:HmacKey"], 1)
}

func TestCreateHmacKey_WithOptionalArgs(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		svc := "service-account@example.com"
		proj := pulumi.String("proj-123")

		opts := &iam.HmacKeyOptions{
			ServiceAccount: svc,
			Project:        proj,
			Rotation:       &rotation.Options{},
			PulumiOptions:  []pulumi.ResourceOption{},
		}

		hk, err := iam.CreateHmacKey(ctx, opts)
		require.NoError(t, err)
		require.NotNil(t, hk)

		hk.ServiceAccountEmail.ApplyT(func(email string) error {
			assert.Equal(t, svc, email)
			return nil
		})
		hk.Project.ApplyT(func(project string) error {
			assert.Equal(t, "proj-123", project)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Len(t, counter.Resources["time:index/rotating:Rotating"], 1)
	assert.Len(t, counter.Resources["gcp:storage/hmacKey:HmacKey"], 1)
}
