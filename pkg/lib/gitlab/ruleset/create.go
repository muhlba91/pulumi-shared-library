package ruleset

import (
	"fmt"

	"github.com/pulumi/pulumi-gitlab/sdk/v10/go/gitlab"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
)

// CreateOptions defines the options for creating a GitLab Repository Ruleset.
type CreateOptions struct {
	// Repository is the repository to which the ruleset will be applied.
	Repository *gitlab.Project
	// Branch is the branch name or wildcard pattern to which the ruleset will apply (see DefaultBranch).
	Branch string
	// ReviewerCount is the number of required approving reviews. Optional, no approval rule is created if nil or 0.
	ReviewerCount *int
	// AllowForcePush indicates whether to allow force pushes. Optional, defaults to false.
	AllowForcePush *bool
	// SignedCommits indicates whether to reject unsigned commits. Optional, defaults to false.
	SignedCommits *bool
	// MemberCheck indicates whether to require commit authors to be existing GitLab users and committers to use a verified
	// e-mail address of their own. Optional, defaults to true.
	MemberCheck *bool
	// CodeOwnerReview indicates whether to require code owner review. Optional, defaults to false.
	CodeOwnerReview *bool
	// DeleteOnDestroy indicates whether to delete the ruleset on destroy. Optional, defaults to false (retained).
	DeleteOnDestroy *bool
	// PulumiOptions are additional options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new GitLab Repository Ruleset with the given options.
// Consists of a branch protection, push rules, and, if reviewers are required, an approval rule.
// Returns the branch protection.
// ctx: The Pulumi context.
// name: The logical name for the Pulumi resources of the ruleset.
// opts: The options for creating the ruleset.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*gitlab.BranchProtection, error) {
	optsWithRepoSpecifics := append([]pulumi.ResourceOption{}, opts.PulumiOptions...)
	optsWithRepoSpecifics = append(
		optsWithRepoSpecifics,
		pulumi.RetainOnDelete(!defaults.GetOrDefault(opts.DeleteOnDestroy, false)),
		pulumi.DependsOn([]pulumi.Resource{opts.Repository}),
	)

	if opts.ReviewerCount != nil && *opts.ReviewerCount > 0 {
		_, bpErr := gitlab.NewProjectApprovalRule(
			ctx,
			fmt.Sprintf("gitlab-project-approval-rule-%s", name),
			&gitlab.ProjectApprovalRuleArgs{
				Project:                       opts.Repository.ID(),
				Name:                          pulumi.String(fmt.Sprintf("gitlab-project-approval-rule-%s", name)),
				AppliesToAllProtectedBranches: pulumi.Bool(true),
				RuleType:                      pulumi.String("regular"),
				ApprovalsRequired:             pulumi.Int(*opts.ReviewerCount),
			},
			optsWithRepoSpecifics...)
		if bpErr != nil {
			return nil, bpErr
		}
	}

	_, pprErr := gitlab.NewProjectPushRules(
		ctx,
		fmt.Sprintf("gitlab-project-push-rules-%s", name),
		&gitlab.ProjectPushRulesArgs{
			Project:               opts.Repository.ID(),
			MemberCheck:           pulumi.Bool(defaults.GetOrDefault(opts.MemberCheck, true)),
			CommitCommitterCheck:  pulumi.Bool(defaults.GetOrDefault(opts.MemberCheck, true)),
			RejectUnsignedCommits: pulumi.Bool(defaults.GetOrDefault(opts.SignedCommits, false)),
		},
		optsWithRepoSpecifics...)
	if pprErr != nil {
		return nil, pprErr
	}

	return gitlab.NewBranchProtection(
		ctx,
		fmt.Sprintf("gitlab-branch-protection-%s", name),
		&gitlab.BranchProtectionArgs{
			Project:                   opts.Repository.ID(),
			Branch:                    pulumi.String(opts.Branch),
			AllowForcePush:            pulumi.Bool(defaults.GetOrDefault(opts.AllowForcePush, false)),
			CodeOwnerApprovalRequired: pulumi.Bool(defaults.GetOrDefault(opts.CodeOwnerReview, false)),
			// MergeAccessLevel:          pulumi.String("developer"),
			// PushAccessLevel:           pulumi.String("maintainer"),
		},
		optsWithRepoSpecifics...)
}
