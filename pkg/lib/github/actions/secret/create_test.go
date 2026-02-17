package secret_test

import (
	"testing"

	gh "github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/github/actions/secret"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreate_GithubActionsSecret(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		repoName := "my-repo"
		repo, err := gh.NewRepository(ctx, "repo", &gh.RepositoryArgs{
			Name: pulumi.String(repoName),
		})
		require.NoError(t, err)
		require.NotNil(t, repo)

		key := "MY_SECRET"
		value := "s3cr3t"

		opts := &secret.CreateOptions{
			Key:        key,
			Value:      pulumi.String(value),
			Repository: repo,
		}

		out := secret.Create(ctx, opts)
		assert.NotNil(t, out)

		out.ApplyT(func(v any) error {
			assert.NotNil(t, v)

			as, ok := v.(*gh.ActionsSecret)
			assert.True(t, ok)
			assert.NotNil(t, as)

			as.Repository.ApplyT(func(r string) error {
				assert.Equal(t, repoName, r)
				return nil
			})
			as.SecretName.ApplyT(func(n string) error {
				assert.Equal(t, key, n)
				return nil
			})
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreate_GithubActionsSecret_WithOptionalArgs(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		repoName := "my-repo"
		repo, err := gh.NewRepository(ctx, "repo", &gh.RepositoryArgs{
			Name: pulumi.String(repoName),
		})
		require.NoError(t, err)
		require.NotNil(t, repo)

		key := "MY_SECRET"
		value := "s3cr3t"

		opts := &secret.CreateOptions{
			Key:        key,
			Value:      pulumi.String(value),
			Repository: repo,
			PulumiOptions: []pulumi.ResourceOption{
				pulumi.DependsOn([]pulumi.Resource{repo}),
			},
		}

		out := secret.Create(ctx, opts)
		assert.NotNil(t, out)

		out.ApplyT(func(v any) error {
			assert.NotNil(t, v)

			as, ok := v.(*gh.ActionsSecret)
			assert.True(t, ok)
			assert.NotNil(t, as)

			as.Repository.ApplyT(func(r string) error {
				assert.Equal(t, repoName, r)
				return nil
			})
			as.SecretName.ApplyT(func(n string) error {
				assert.Equal(t, key, n)
				return nil
			})
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
