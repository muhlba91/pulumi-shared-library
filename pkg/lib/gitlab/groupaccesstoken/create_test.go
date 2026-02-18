package groupaccesstoken_test

import (
	"testing"

	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	accesstoken "github.com/muhlba91/pulumi-shared-library/pkg/lib/gitlab/groupaccesstoken"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreate(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		group := "test-group"
		opts := &accesstoken.CreateOptions{
			Name:        pulumi.String("test-token"),
			Description: pulumi.String("test description"),
			Group:       group,
			Scopes:      []string{"api", "read_user"},
		}

		r, err := accesstoken.Create(ctx, "test", opts)
		require.NoError(t, err)
		require.NotNil(t, r)

		r.Name.ApplyT(func(n string) error {
			assert.Equal(t, "test-token", n)
			return nil
		})
		r.Description.ApplyT(func(d string) error {
			assert.Equal(t, "test description", d)
			return nil
		})
		r.Group.ApplyT(func(id string) error {
			assert.Equal(t, group, id)
			return nil
		})
		r.AccessLevel.ApplyT(func(id string) error {
			assert.Equal(t, "owner", id)
			return nil
		})
		r.RotationConfiguration.ApplyT(func(rc *gitlab.GroupAccessTokenRotationConfiguration) error {
			assert.Equal(t, 365, rc.ExpirationDays)
			assert.Equal(t, 30, rc.RotateBeforeDays)
			return nil
		})
		r.Scopes.ApplyT(func(s []string) error {
			assert.ElementsMatch(t, []string{"api", "read_user", "self_rotate"}, s)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreate_Minimal(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		group := "test-group"
		opts := &accesstoken.CreateOptions{
			Name:  pulumi.String("minimal-token"),
			Group: group,
		}

		r, err := accesstoken.Create(ctx, "minimal", opts)
		require.NoError(t, err)
		require.NotNil(t, r)

		r.Group.ApplyT(func(id string) error {
			assert.Equal(t, group, id)
			return nil
		})
		r.Scopes.ApplyT(func(s []string) error {
			assert.Equal(t, []string{"self_rotate"}, s)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
