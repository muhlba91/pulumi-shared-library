package ruleset_test

import (
	"testing"

	"github.com/pulumi/pulumi-gitlab/sdk/v10/go/gitlab"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libruleset "github.com/muhlba91/pulumi-shared-library/pkg/lib/gitlab/ruleset"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateRuleset(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		repo, err := gitlab.NewProject(ctx, "repo-basic", &gitlab.ProjectArgs{
			Name: pulumi.String("myrepo"),
		})
		require.NoError(t, err)
		require.NotNil(t, repo)

		opts := &libruleset.CreateOptions{
			Repository: repo,
			Branch:     "main",
		}

		rs, err := libruleset.Create(ctx, "basic", opts)
		require.NoError(t, err)
		require.NotNil(t, rs)

		rs.Project.ApplyT(func(r string) error {
			assert.Equal(t, "repo-basic_id", r)
			return nil
		})
		rs.Branch.ApplyT(func(target string) error {
			assert.Equal(t, "main", target)
			return nil
		})

		rs.AllowForcePush.ApplyT(func(e bool) error {
			assert.False(t, e)
			return nil
		})
		rs.CodeOwnerApprovalRequired.ApplyT(func(e bool) error {
			assert.False(t, e)
			return nil
		})
		rs.AllowedToMerges.ApplyT(func(a []gitlab.BranchProtectionAllowedToMerge) error {
			assert.Empty(t, a)
			return nil
		})
		rs.AllowedToPushes.ApplyT(func(a []gitlab.BranchProtectionAllowedToPush) error {
			assert.Empty(t, a)
			return nil
		})

		// rs.PushAccessLevel.ApplyT(func(e string) error {
		// 	assert.Equal(t, "maintainer", e)
		// 	return nil
		// })
		// rs.MergeAccessLevel.ApplyT(func(e string) error {
		// 	assert.Equal(t, "developer", e)
		// 	return nil
		// })
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}

func TestCreateRuleset_Custom(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		repo, err := gitlab.NewProject(ctx, "repo-basic", &gitlab.ProjectArgs{
			Name: pulumi.String("custom-repo"),
		})
		require.NoError(t, err)
		require.NotNil(t, repo)

		allowForcePush := true
		signedCommits := true
		codeOwnerReview := true
		memberCheck := true
		deleteOnDestroy := true
		reviewerCount := 1
		allowedToMerge := "developer"
		allowedToPush := "maintainer"

		opts := &libruleset.CreateOptions{
			Repository:      repo,
			Branch:          "main",
			AllowForcePush:  &allowForcePush,
			SignedCommits:   &signedCommits,
			CodeOwnerReview: &codeOwnerReview,
			DeleteOnDestroy: &deleteOnDestroy,
			MemberCheck:     &memberCheck,
			ReviewerCount:   &reviewerCount,
			AllowedToMerge:  &allowedToMerge,
			AllowedToPush:   &allowedToPush,
		}

		rs, err := libruleset.Create(ctx, "custom", opts)
		require.NoError(t, err)
		require.NotNil(t, rs)

		rs.AllowForcePush.ApplyT(func(e bool) error {
			assert.True(t, e)
			return nil
		})
		rs.CodeOwnerApprovalRequired.ApplyT(func(e bool) error {
			assert.True(t, e)
			return nil
		})
		rs.AllowedToMerges.ApplyT(func(a []gitlab.BranchProtectionAllowedToMerge) error {
			require.Len(t, a, 1)
			assert.Equal(t, "developer", *a[0].AccessLevel)
			return nil
		})
		rs.AllowedToPushes.ApplyT(func(a []gitlab.BranchProtectionAllowedToPush) error {
			require.Len(t, a, 1)
			assert.Equal(t, "maintainer", *a[0].AccessLevel)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}
