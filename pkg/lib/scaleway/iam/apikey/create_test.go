package apikey_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libapikey "github.com/muhlba91/pulumi-shared-library/pkg/lib/scaleway/iam/apikey"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateAPIKey(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		userName := "test-user"

		ak, err := libapikey.Create(ctx, "api-key", &libapikey.CreateOptions{
			UserID: pulumi.String(userName),
		})
		require.NoError(t, err)
		require.NotNil(t, ak)

		ak.UserId.ApplyT(func(u *string) error {
			assert.Equal(t, userName, *u)
			return nil
		})
		ak.Description.ApplyT(func(u *string) error {
			assert.Nil(t, u)
			return nil
		})
		ak.DefaultProjectId.ApplyT(func(u string) error {
			assert.Empty(t, u)
			return nil
		})
		ak.ApplicationId.ApplyT(func(u *string) error {
			assert.Nil(t, u)
			return nil
		})
		ak.AccessKey.ApplyT(func(s string) error {
			assert.NotEmpty(t, s)
			return nil
		})
		ak.SecretKey.ApplyT(func(s string) error {
			assert.NotEmpty(t, s)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateAPIKey_WithOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		userName := "protected-user"

		ak, err := libapikey.Create(ctx, "api-key", &libapikey.CreateOptions{
			ApplicationID:    pulumi.String(userName),
			Description:      pulumi.String(userName),
			DefaultProjectID: pulumi.String(userName),
			PulumiOptions: []pulumi.ResourceOption{
				pulumi.Protect(true),
			},
		})
		require.NoError(t, err)
		require.NotNil(t, ak)

		ak.UserId.ApplyT(func(u *string) error {
			assert.Nil(t, u)
			return nil
		})
		ak.Description.ApplyT(func(u *string) error {
			assert.Equal(t, userName, *u)
			return nil
		})
		ak.DefaultProjectId.ApplyT(func(u string) error {
			assert.Equal(t, userName, u)
			return nil
		})
		ak.ApplicationId.ApplyT(func(u *string) error {
			assert.Equal(t, userName, *u)
			return nil
		})
		ak.AccessKey.ApplyT(func(s string) error {
			assert.NotEmpty(t, s)
			return nil
		})
		ak.SecretKey.ApplyT(func(s string) error {
			assert.NotEmpty(t, s)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
