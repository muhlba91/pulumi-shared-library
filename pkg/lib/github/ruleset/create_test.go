package ruleset_test

import (
	"testing"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libruleset "github.com/muhlba91/pulumi-shared-library/pkg/lib/github/ruleset"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateRuleset(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		repo, err := github.NewRepository(ctx, "repo-basic", &github.RepositoryArgs{
			Name: pulumi.String("myrepo"),
		})
		require.NoError(t, err)
		require.NotNil(t, repo)

		opts := &libruleset.CreateOptions{
			Repository: repo,
			Patterns:   []string{"main", "release"},
		}

		rs, err := libruleset.Create(ctx, "basic", opts)
		require.NoError(t, err)
		require.NotNil(t, rs)

		rs.Repository.ApplyT(func(r string) error {
			assert.Equal(t, "myrepo", r)
			return nil
		})
		rs.Target.ApplyT(func(target string) error {
			assert.Equal(t, "branch", target)
			return nil
		})
		rs.Enforcement.ApplyT(func(e string) error {
			assert.Equal(t, "active", e)
			return nil
		})

		rs.Conditions.ApplyT(func(c *github.RepositoryRulesetConditions) error {
			includes := c.RefName.Includes
			assert.Contains(t, includes, "main")
			assert.Contains(t, includes, "release")
			return nil
		})

		rs.BypassActors.ApplyT(func(b []github.RepositoryRulesetBypassActor) error {
			assert.GreaterOrEqual(t, len(b), 2)
			return nil
		})

		rs.Rules.ApplyT(func(r github.RepositoryRulesetRules) error {
			if r.MergeQueue != nil {
				assert.Fail(t, "expected MergeQueue to be nil for default")
			}
			assert.True(t, *r.Creation)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateRuleset_Custom(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		repo, err := github.NewRepository(ctx, "repo-custom", &github.RepositoryArgs{
			Name: pulumi.String("custom-repo"),
		})
		require.NoError(t, err)
		require.NotNil(t, repo)

		restrictCreation := false
		allowForcePush := true
		signedCommits := true
		codeOwnerReview := true
		conversationResolution := false
		lastPushApproval := false
		reviewerCount := 2
		enableMergeQueue := true
		deleteOnDestroy := true
		allowBypass := true
		allowBypassIntegrations := []int{10, 3}
		updatedBeforeMerge := false
		requiredChecks := []string{"ci", "lint"}
		wipIntegration := false

		opts := &libruleset.CreateOptions{
			Repository:               repo,
			Patterns:                 []string{"feature/*", "hotfix/*"},
			RestrictCreation:         &restrictCreation,
			AllowForcePush:           &allowForcePush,
			SignedCommits:            &signedCommits,
			CodeOwnerReview:          &codeOwnerReview,
			ConversationResolution:   &conversationResolution,
			LastPushApproval:         &lastPushApproval,
			ReviewerCount:            &reviewerCount,
			EnableMergeQueue:         &enableMergeQueue,
			DeleteOnDestroy:          &deleteOnDestroy,
			AllowBypass:              &allowBypass,
			AllowBypassIntegrations:  allowBypassIntegrations,
			UpdatedBranchBeforeMerge: &updatedBeforeMerge,
			RequiredChecks:           requiredChecks,
			WIPIntegration:           &wipIntegration,
		}

		rs, err := libruleset.Create(ctx, "custom", opts)
		require.NoError(t, err)
		require.NotNil(t, rs)

		rs.Rules.ApplyT(func(r github.RepositoryRulesetRules) error {
			assert.False(t, *r.Creation)
			assert.False(t, *r.NonFastForward)
			assert.True(t, *r.RequiredSignatures)

			if r.PullRequest != nil {
				assert.True(t, *r.PullRequest.RequireCodeOwnerReview)
				assert.Equal(t, 2, *r.PullRequest.RequiredApprovingReviewCount)
				assert.False(t, *r.PullRequest.RequiredReviewThreadResolution)
				assert.False(t, *r.PullRequest.RequireLastPushApproval)
			} else {
				assert.Fail(t, "PullRequest is nil")
			}

			if r.MergeQueue == nil {
				assert.Fail(t, "expected MergeQueue to be present")
			}
			return nil
		})

		rs.BypassActors.ApplyT(func(b []github.RepositoryRulesetBypassActor) error {
			assert.GreaterOrEqual(t, len(b), 4)
			return nil
		})

		rs.Rules.ApplyT(func(r github.RepositoryRulesetRules) error {
			if r.RequiredStatusChecks != nil && len(r.RequiredStatusChecks.RequiredChecks) > 0 {
				found := map[string]bool{}
				for _, c := range r.RequiredStatusChecks.RequiredChecks {
					found[c.Context] = true
				}
				assert.True(t, found["ci"])
				assert.True(t, found["lint"])
				assert.False(t, *r.RequiredStatusChecks.StrictRequiredStatusChecksPolicy)
			} else {
				assert.Fail(t, "RequiredStatusChecks missing")
			}
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
