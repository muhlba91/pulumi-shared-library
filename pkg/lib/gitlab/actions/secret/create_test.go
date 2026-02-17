package secret_test

import (
	"testing"

	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/gitlab/actions/secret"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreate_GitLabActionsSecret(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		repoName := "my-repo"
		repo, err := gitlab.NewProject(ctx, "repo", &gitlab.ProjectArgs{
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

			as, ok := v.(*gitlab.ProjectVariable)
			assert.True(t, ok)
			assert.NotNil(t, as)

			as.Project.ApplyT(func(r string) error {
				assert.Equal(t, "repo_id", r)
				return nil
			})
			as.Key.ApplyT(func(n string) error {
				assert.Equal(t, key, n)
				return nil
			})
			as.VariableType.ApplyT(func(n string) error {
				assert.Equal(t, "env_var", n)
				return nil
			})

			as.Hidden.ApplyT(func(n bool) error {
				assert.True(t, n)
				return nil
			})
			as.Masked.ApplyT(func(n bool) error {
				assert.True(t, n)
				return nil
			})
			as.Protected.ApplyT(func(n bool) error {
				assert.False(t, n)
				return nil
			})
			as.Raw.ApplyT(func(n bool) error {
				assert.False(t, n)
				return nil
			})
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreate_GitLabActionsSecret_WithOptionalArgs(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		repoName := "my-repo"
		repo, err := gitlab.NewProject(ctx, "repo", &gitlab.ProjectArgs{
			Name: pulumi.String(repoName),
		})
		require.NoError(t, err)
		require.NotNil(t, repo)

		key := "MY_SECRET"
		value := "s3cr3t"
		trueValue := true
		variableType := "file"

		opts := &secret.CreateOptions{
			Key:                      key,
			Value:                    pulumi.String(value),
			Repository:               repo,
			Protected:                &trueValue,
			DisableVariableExpansion: &trueValue,
			VariableType:             &variableType,
			PulumiOptions: []pulumi.ResourceOption{
				pulumi.DependsOn([]pulumi.Resource{repo}),
			},
		}

		out := secret.Create(ctx, opts)
		assert.NotNil(t, out)

		out.ApplyT(func(v any) error {
			assert.NotNil(t, v)

			as, ok := v.(*gitlab.ProjectVariable)
			assert.True(t, ok)
			assert.NotNil(t, as)

			as.Project.ApplyT(func(r string) error {
				assert.Equal(t, "repo_id", r)
				return nil
			})
			as.Key.ApplyT(func(n string) error {
				assert.Equal(t, key, n)
				return nil
			})
			as.VariableType.ApplyT(func(n string) error {
				assert.Equal(t, "file", n)
				return nil
			})

			as.Hidden.ApplyT(func(n bool) error {
				assert.True(t, n)
				return nil
			})
			as.Masked.ApplyT(func(n bool) error {
				assert.True(t, n)
				return nil
			})
			as.Protected.ApplyT(func(n bool) error {
				assert.True(t, n)
				return nil
			})
			as.Raw.ApplyT(func(n bool) error {
				assert.True(t, n)
				return nil
			})
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
