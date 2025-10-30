package ruleset

import (
	"fmt"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
)

// CreateOptions defines the options for creating a GitHub Repository Ruleset.
type CreateOptions struct {
	// Repository is the name of the repository to which the ruleset will be applied.
	Repository *github.Repository
	// Patterns are the branch name patterns to which the ruleset will apply.
	Patterns []string
	// RestrictCreation indicates whether to restrict branch creation.
	RestrictCreation *bool
	// AllowForcePush indicates whether to allow force pushes.
	AllowForcePush *bool
	// SignedCommits indicates whether to require signed commits.
	SignedCommits *bool
	// CodeOwnerReview indicates whether to require code owner review.
	CodeOwnerReview *bool
	// ConversationResolution indicates whether to require conversation resolution.
	ConversationResolution *bool
	// LastPushApproval indicates whether to require approval for the last push.
	LastPushApproval *bool
	// ReviewerCount is the number of required approving reviews.
	ReviewerCount *int
	// EnableMergeQueue indicates whether to enable the merge queue.
	EnableMergeQueue *bool
	// DeleteOnDestroy indicates whether to delete the ruleset on destroy.
	DeleteOnDestroy *bool
	// AllowBypass indicates whether to allow bypassing the ruleset.
	AllowBypass *bool
	// AllowBypassIntegrations are the IDs of integrations allowed to bypass the ruleset.
	AllowBypassIntegrations []int
	// UpdatedBranchBeforeMerge indicates whether to require an updated branch before merging.
	UpdatedBranchBeforeMerge *bool
	// RequiredChecks are the required status checks for the ruleset.
	RequiredChecks []string
	// WIPIntegration indicates whether to enable WIP integration.
	WIPIntegration *bool
	// PulumiOptions are additional options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new GitHub Repository Ruleset with the given options.
// ctx: The Pulumi context.
// name: The name of the ruleset.
// opts: The options for creating the ruleset.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*github.RepositoryRuleset, error) {
	optsWithRepoSpecifics := append([]pulumi.ResourceOption{}, opts.PulumiOptions...)
	optsWithRepoSpecifics = append(
		optsWithRepoSpecifics,
		pulumi.RetainOnDelete(!defaults.GetOrDefault(opts.DeleteOnDestroy, false)),
		pulumi.DependsOn([]pulumi.Resource{opts.Repository}),
	)

	mergeQueue := buildMergeQueueArgs(opts)
	bypassActors := buildBypassActorsArgs(opts)
	reqStatusChecks := buildRequiredStatusChecksArgs(opts)

	return github.NewRepositoryRuleset(
		ctx,
		fmt.Sprintf("github-repository-ruleset-%s", name),
		&github.RepositoryRulesetArgs{
			Repository:  opts.Repository.Name,
			Target:      pulumi.String("branch"),
			Enforcement: pulumi.String("active"),
			Conditions: &github.RepositoryRulesetConditionsArgs{
				RefName: &github.RepositoryRulesetConditionsRefNameArgs{
					Excludes: pulumi.ToStringArray([]string{}),
					Includes: pulumi.ToStringArray(opts.Patterns),
				},
			},
			BypassActors: bypassActors,
			Rules: &github.RepositoryRulesetRulesArgs{
				Creation:                  pulumi.Bool(defaults.GetOrDefault(opts.RestrictCreation, true)),
				Deletion:                  pulumi.Bool(true),
				NonFastForward:            pulumi.Bool(!defaults.GetOrDefault(opts.AllowForcePush, false)),
				RequiredDeployments:       &github.RepositoryRulesetRulesRequiredDeploymentsArgs{},
				RequiredLinearHistory:     pulumi.Bool(true),
				RequiredSignatures:        pulumi.Bool(defaults.GetOrDefault(opts.SignedCommits, false)),
				Update:                    pulumi.Bool(false),
				UpdateAllowsFetchAndMerge: pulumi.Bool(false),
				PullRequest: &github.RepositoryRulesetRulesPullRequestArgs{
					DismissStaleReviewsOnPush:    pulumi.Bool(true),
					RequireCodeOwnerReview:       pulumi.Bool(defaults.GetOrDefault(opts.CodeOwnerReview, false)),
					RequiredApprovingReviewCount: pulumi.Int(defaults.GetOrDefault(opts.ReviewerCount, 0)),
					RequiredReviewThreadResolution: pulumi.Bool(
						defaults.GetOrDefault(opts.ConversationResolution, true),
					),
					RequireLastPushApproval: pulumi.Bool(defaults.GetOrDefault(opts.LastPushApproval, true)),
				},
				RequiredStatusChecks: reqStatusChecks,
				MergeQueue:           mergeQueue,
			},
		},
		optsWithRepoSpecifics...)
}
