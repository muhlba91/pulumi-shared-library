package ruleset

import (
	"sort"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
)

// Repository role IDs of the bypass actors.
const (
	maintainerActorID = 2
	adminActorID      = 5
)

// Status check context and integration ID of the WIP integration.
const (
	wipIntegrationContext = "WIP"
	wipIntegrationID      = 3414
)

// githubActionsIntegrationID is the integration ID of GitHub Actions, which reports all other required checks.
const githubActionsIntegrationID = 15368

// buildMergeQueueArgs builds the merge queue arguments for the ruleset based on the provided options.
// Returns nil if the merge queue is not enabled.
// opts: The options for creating the ruleset.
//
//nolint:mnd // magic number is acceptable here for configuration defaults
func buildMergeQueueArgs(opts *CreateOptions) *github.RepositoryRulesetRulesMergeQueueArgs {
	if !defaults.GetOrDefault(opts.EnableMergeQueue, false) {
		return nil
	}
	return &github.RepositoryRulesetRulesMergeQueueArgs{
		CheckResponseTimeoutMinutes:  pulumi.Int(60),
		GroupingStrategy:             pulumi.String("ALLGREEN"),
		MaxEntriesToBuild:            pulumi.Int(5),
		MaxEntriesToMerge:            pulumi.Int(5),
		MergeMethod:                  pulumi.String("REBASE"),
		MinEntriesToMerge:            pulumi.Int(1),
		MinEntriesToMergeWaitMinutes: pulumi.Int(5),
	}
}

// buildBypassActorsArgs builds the bypass actors arguments for the ruleset based on the provided options.
// Maintainers may bypass for pull requests only, admins and integrations always. Returns no actors if bypass is not allowed.
// opts: The options for creating the ruleset.
func buildBypassActorsArgs(opts *CreateOptions) github.RepositoryRulesetBypassActorArray {
	var bypassActors github.RepositoryRulesetBypassActorArray
	if !defaults.GetOrDefault(opts.AllowBypass, true) {
		return bypassActors
	}

	actors := []int{
		maintainerActorID,
		adminActorID,
	}
	actors = append(actors, opts.AllowBypassIntegrations...)
	sort.Ints(actors)

	for _, actorID := range actors {
		var actorType string
		var bypassMode string
		switch actorID {
		case maintainerActorID:
			actorType = "RepositoryRole"
			bypassMode = "pull_request"
		case adminActorID:
			actorType = "RepositoryRole"
			bypassMode = "always"
		default:
			actorType = "Integration"
			bypassMode = "always"
		}
		bypassActors = append(bypassActors, &github.RepositoryRulesetBypassActorArgs{
			ActorId:    pulumi.Int(actorID),
			ActorType:  pulumi.String(actorType),
			BypassMode: pulumi.String(bypassMode),
		})
	}

	return bypassActors
}

// buildRequiredStatusChecksArgs builds the required status checks args for the ruleset.
// Returns nil if there are no required checks and the WIP integration is disabled.
// opts: The options for creating the ruleset.
func buildRequiredStatusChecksArgs(opts *CreateOptions) *github.RepositoryRulesetRulesRequiredStatusChecksArgs {
	var reqStatusChecks *github.RepositoryRulesetRulesRequiredStatusChecksArgs
	wipIntegration := defaults.GetOrDefault(opts.WIPIntegration, true)
	if len(opts.RequiredChecks) > 0 || wipIntegration {
		contexts := []string{}
		if wipIntegration {
			contexts = append(contexts, wipIntegrationContext)
		}
		contexts = append(contexts, opts.RequiredChecks...)
		sort.Strings(contexts)

		checks := github.RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArray{}
		for _, c := range contexts {
			var integrationID int
			if wipIntegration && c == wipIntegrationContext {
				integrationID = wipIntegrationID
			} else {
				integrationID = githubActionsIntegrationID
			}
			checks = append(checks, &github.RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs{
				Context:       pulumi.String(c),
				IntegrationId: pulumi.Int(integrationID),
			})
		}

		reqStatusChecks = &github.RepositoryRulesetRulesRequiredStatusChecksArgs{
			RequiredChecks:                   checks,
			StrictRequiredStatusChecksPolicy: pulumi.Bool(defaults.GetOrDefault(opts.UpdatedBranchBeforeMerge, true)),
		}
	}
	return reqStatusChecks
}
