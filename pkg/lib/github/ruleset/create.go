package ruleset

import (
	"fmt"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
)

// CreateOptions defines the options for creating a GitHub Repository Ruleset.
type CreateOptions struct {
	// Repository is the repository to which the ruleset will be applied.
	Repository *github.Repository
	// Patterns are the branch name patterns to which the ruleset will apply (see DefaultBranchRulesetPattern).
	Patterns []string
	// RestrictCreation indicates whether to restrict branch creation. Optional, defaults to true.
	RestrictCreation *bool
	// AllowForcePush indicates whether to allow force pushes. Optional, defaults to false.
	AllowForcePush *bool
	// SignedCommits indicates whether to require signed commits. Optional, defaults to false.
	SignedCommits *bool
	// CodeOwnerReview indicates whether to require code owner review. Optional, defaults to false.
	CodeOwnerReview *bool
	// ConversationResolution indicates whether to require conversation resolution. Optional, defaults to true.
	ConversationResolution *bool
	// LastPushApproval indicates whether to require approval for the last push. Optional, defaults to true.
	LastPushApproval *bool
	// ReviewerCount is the number of required approving reviews. Optional, defaults to 0.
	ReviewerCount *int
	// EnableMergeQueue indicates whether to enable the merge queue. Optional, defaults to false.
	EnableMergeQueue *bool
	// DeleteOnDestroy indicates whether to delete the ruleset on destroy. Optional, defaults to false (retained).
	DeleteOnDestroy *bool
	// AllowBypass indicates whether repository maintainers (pull requests only) and admins may bypass the ruleset.
	// Optional, defaults to true.
	AllowBypass *bool
	// AllowBypassIntegrations are the IDs of integrations allowed to bypass the ruleset. Only applies if bypass is allowed.
	AllowBypassIntegrations []int
	// UpdatedBranchBeforeMerge indicates whether to require an updated branch before merging. Optional, defaults to true.
	// Only applies if there are required status checks.
	UpdatedBranchBeforeMerge *bool
	// RequiredChecks are the required status checks for the ruleset, expected to be reported by GitHub Actions.
	RequiredChecks []string
	// CopilotReview indicates whether to enable Copilot code review on push. Optional, defaults to true.
	CopilotReview *bool
	// WIPIntegration indicates whether to require the status check of the WIP integration. Optional, defaults to true.
	WIPIntegration *bool
	// PulumiOptions are additional options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new GitHub Repository Ruleset with the given options.
// ctx: The Pulumi context.
// name: The logical name for the Pulumi resource (prefixed with "github-repository-ruleset-").
// opts: The options for creating the ruleset.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*github.RepositoryRuleset, error) {
	optsWithRepoSpecifics := append([]pulumi.ResourceOption{}, opts.PulumiOptions...)
	optsWithRepoSpecifics = append(
		optsWithRepoSpecifics,
		pulumi.RetainOnDelete(!defaults.GetOrDefault(opts.DeleteOnDestroy, false)),
		pulumi.DependsOn([]pulumi.Resource{opts.Repository}),
	)
	optsWithRepoSpecifics = append(optsWithRepoSpecifics, pulumi.Parent(opts.Repository))

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
				RequiredLinearHistory:     pulumi.Bool(true),
				RequiredSignatures:        pulumi.Bool(defaults.GetOrDefault(opts.SignedCommits, false)),
				Update:                    pulumi.Bool(false),
				UpdateAllowsFetchAndMerge: pulumi.Bool(false),
				CopilotCodeReview: &github.RepositoryRulesetRulesCopilotCodeReviewArgs{
					ReviewDraftPullRequests: pulumi.Bool(false),
					ReviewOnPush:            pulumi.Bool(defaults.GetOrDefault(opts.CopilotReview, true)),
				},
				PullRequest: &github.RepositoryRulesetRulesPullRequestArgs{
					AllowedMergeMethods:          pulumi.StringArray{pulumi.String("rebase")},
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
