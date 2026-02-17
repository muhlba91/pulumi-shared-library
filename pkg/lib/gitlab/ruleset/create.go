package ruleset

import (
	"fmt"

	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
)

// CreateOptions defines the options for creating a GitLab Repository Ruleset.
type CreateOptions struct {
	// Repository is the name of the repository to which the ruleset will be applied.
	Repository *gitlab.Project
	// Branch is the branch pattern to which the ruleset will apply.
	Branch string
	// AllowForcePush indicates whether to allow force pushes.
	AllowForcePush *bool
	// SignedCommits indicates whether to require signed commits.
	SignedCommits *bool
	// MemberCheck indicates whether to require member e-mails for pushing.
	MemberCheck *bool
	// CodeOwnerReview indicates whether to require code owner review.
	CodeOwnerReview *bool
	// DeleteOnDestroy indicates whether to delete the ruleset on destroy.
	DeleteOnDestroy *bool
	// PulumiOptions are additional options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new GitLab Repository Ruleset with the given options.
// ctx: The Pulumi context.
// name: The name of the ruleset.
// opts: The options for creating the ruleset.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*gitlab.BranchProtection, error) {
	optsWithRepoSpecifics := append([]pulumi.ResourceOption{}, opts.PulumiOptions...)
	optsWithRepoSpecifics = append(
		optsWithRepoSpecifics,
		pulumi.RetainOnDelete(!defaults.GetOrDefault(opts.DeleteOnDestroy, false)),
		pulumi.DependsOn([]pulumi.Resource{opts.Repository}),
	)

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
			MergeAccessLevel:          pulumi.String("developer"),
			PushAccessLevel:           pulumi.String("maintainer"),
			UnprotectAccessLevel:      pulumi.String("admin"),
		},
		optsWithRepoSpecifics...)
}
