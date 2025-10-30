package repository_test

import (
	"testing"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	librepo "github.com/muhlba91/pulumi-shared-library/pkg/lib/github/repository"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateRepository_Public(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		visibility := "public"

		opts := &librepo.CreateOptions{
			Name:                    pulumi.String("repo-name"),
			Description:             pulumi.String("desc"),
			EnableDiscussions:       pulumi.Bool(false),
			EnableWiki:              pulumi.Bool(true),
			Homepage:                pulumi.String("https://example.com"),
			Topics:                  []string{"z", "a"},
			Visibility:              &visibility,
			Protected:               false,
			AllowRepositoryDeletion: false,
		}

		r, err := librepo.Create(ctx, "res-public", opts)
		require.NoError(t, err)
		require.NotNil(t, r)

		r.Name.ApplyT(func(n string) error {
			assert.Equal(t, "repo-name", n)
			return nil
		})
		r.Visibility.ApplyT(func(v string) error {
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

		r.Pages.ApplyT(func(p *github.RepositoryPages) error {
			assert.NotNil(t, p)
			assert.Equal(t, "workflow", *p.BuildType)
			assert.Equal(t, "main", p.Source.Branch)
			assert.Equal(t, "/", *p.Source.Path)
			return nil
		})

		r.SecurityAndAnalysis.ApplyT(func(s github.RepositorySecurityAndAnalysis) error {
			assert.NotNil(t, s)
			assert.Nil(t, s.AdvancedSecurity)
			assert.NotNil(t, s.SecretScanning)
			assert.Equal(t, "enabled", s.SecretScanning.Status)
			assert.NotNil(t, s.SecretScanningPushProtection)
			assert.Equal(t, "enabled", s.SecretScanningPushProtection.Status)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateRepository_GhPages(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		visibility := "public"
		ghPagesBranch := "gh-pages"

		opts := &librepo.CreateOptions{
			Name:                    pulumi.String("repo-name"),
			Description:             pulumi.String("desc"),
			EnableDiscussions:       pulumi.Bool(false),
			EnableWiki:              pulumi.Bool(true),
			Homepage:                pulumi.String("https://example.com"),
			Topics:                  []string{"z", "a"},
			Visibility:              &visibility,
			GitHubPagesBranch:       &ghPagesBranch,
			Protected:               false,
			AllowRepositoryDeletion: false,
		}

		r, err := librepo.Create(ctx, "res-public", opts)
		require.NoError(t, err)
		require.NotNil(t, r)

		r.Name.ApplyT(func(n string) error {
			assert.Equal(t, "repo-name", n)
			return nil
		})
		r.Visibility.ApplyT(func(v string) error {
			assert.Equal(t, "public", v)
			return nil
		})

		r.Pages.ApplyT(func(p *github.RepositoryPages) error {
			assert.NotNil(t, p)
			assert.Equal(t, "workflow", *p.BuildType)
			assert.Equal(t, "gh-pages", p.Source.Branch)
			assert.Equal(t, "/", *p.Source.Path)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateRepository_Private(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		visibility := "private"
		gitHubPagesBranch := "gh-pages"

		opts := &librepo.CreateOptions{
			Name:                    pulumi.String("repo-private"),
			Description:             pulumi.String("private repo"),
			EnableDiscussions:       pulumi.Bool(false),
			EnableWiki:              pulumi.Bool(false),
			Homepage:                pulumi.String(""),
			Topics:                  []string{"b", "a"},
			GitHubPagesBranch:       &gitHubPagesBranch,
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
		r.Visibility.ApplyT(func(v string) error {
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

		r.Pages.ApplyT(func(p any) error {
			assert.Nil(t, p)
			return nil
		})

		r.SecurityAndAnalysis.ApplyT(func(s github.RepositorySecurityAndAnalysis) error {
			assert.NotNil(t, s)
			assert.Nil(t, s.AdvancedSecurity)
			assert.Nil(t, s.SecretScanning)
			assert.Nil(t, s.SecretScanningPushProtection)
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
			EnableDiscussions:       pulumi.Bool(false),
			EnableWiki:              pulumi.Bool(false),
			Topics:                  []string{},
			Protected:               false,
			AllowRepositoryDeletion: false,
		}

		r, err := librepo.Create(ctx, "res-default", opts)
		require.NoError(t, err)
		require.NotNil(t, r)

		r.Visibility.ApplyT(func(v string) error {
			assert.Equal(t, "public", v)
			return nil
		})

		r.Pages.ApplyT(func(p any) error {
			assert.NotNil(t, p)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
