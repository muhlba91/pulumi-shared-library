package repository_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	librepo "github.com/muhlba91/pulumi-shared-library/pkg/lib/gitlab/repository"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateRepository_Public(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		visibility := "public"
		trueValue := true

		opts := &librepo.CreateOptions{
			Name:                    pulumi.String("repo-name"),
			Description:             pulumi.String("desc"),
			EnableWiki:              &trueValue,
			Topics:                  []string{"z", "a"},
			Visibility:              &visibility,
			AutoDevopsEnabled:       &trueValue,
			ConversationResolution:  &trueValue,
			EnableMergeQueue:        &trueValue,
			NamespaceID:             pulumi.IntPtr(123),
			Protected:               false,
			AllowRepositoryDeletion: false,
			RetainOnDelete:          &trueValue,
		}

		r, err := librepo.Create(ctx, "res-public", opts)
		require.NoError(t, err)
		require.NotNil(t, r)

		r.Name.ApplyT(func(n string) error {
			assert.Equal(t, "repo-name", n)
			return nil
		})
		r.VisibilityLevel.ApplyT(func(v string) error {
			assert.Equal(t, "public", v)
			return nil
		})
		r.Topics.ApplyT(func(ts []string) error {
			assert.Equal(t, []string{"a", "z"}, ts)
			return nil
		})
		r.ArchiveOnDestroy.ApplyT(func(ad *bool) error {
			assert.False(t, *ad)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateRepository_Private(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		visibility := "private"
		wiki := true

		opts := &librepo.CreateOptions{
			Name:                    pulumi.String("repo-private"),
			Description:             pulumi.String("private repo"),
			EnableWiki:              &wiki,
			Topics:                  []string{"b", "a"},
			Visibility:              &visibility,
			Protected:               true,
			AllowRepositoryDeletion: true,
		}

		r, err := librepo.Create(ctx, "res-private", opts)
		require.NoError(t, err)
		require.NotNil(t, r)

		r.Name.ApplyT(func(n string) error {
			assert.Equal(t, "repo-private", n)
			return nil
		})
		r.VisibilityLevel.ApplyT(func(v string) error {
			assert.Equal(t, "private", v)
			return nil
		})
		r.Topics.ApplyT(func(ts []string) error {
			assert.Equal(t, []string{"a", "b"}, ts)
			return nil
		})
		r.ArchiveOnDestroy.ApplyT(func(ad *bool) error {
			assert.True(t, *ad)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateRepository_DefaultVisibility(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &librepo.CreateOptions{
			Name:                    pulumi.String("repo-default"),
			Description:             pulumi.String("default visibility"),
			Topics:                  []string{},
			Protected:               false,
			AllowRepositoryDeletion: false,
		}

		r, err := librepo.Create(ctx, "res-default", opts)
		require.NoError(t, err)
		require.NotNil(t, r)

		r.VisibilityLevel.ApplyT(func(v string) error {
			assert.Equal(t, "public", v)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
