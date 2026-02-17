package repository

import (
	"fmt"
	"sort"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
	utilgitlab "github.com/muhlba91/pulumi-shared-library/pkg/util/gitlab"
)

const (
	defaultCiDefaultGitDepth          = 1
	defaultCiDeletePipelinesInSeconds = 3600
	defaultVisibility                 = "public"
)

// CreateOptions defines the options for creating a GitLab repository.
type CreateOptions struct {
	// Name is the name of the repository.
	Name pulumi.StringInput
	// Description is the description of the repository.
	Description pulumi.StringInput
	// NamespaceID is the ID of the namespace under which the repository will be created (group or user).
	NamespaceID pulumi.IntPtrInput
	// EnableWiki indicates whether to enable the wiki for the repository.
	EnableWiki *bool
	// Topics is a list of topics to associate with the repository.
	Topics []string
	// Visibility is the visibility level of the repository. Can be "public" or "private".
	Visibility *string
	// ConversationResolution indicates whether to require conversation resolution.
	ConversationResolution *bool
	// AutoDevopsEnabled indicates whether to enable Auto DevOps for the repository.
	AutoDevopsEnabled *bool
	// EnableMergeQueue indicates whether to enable the merge queue.
	EnableMergeQueue *bool
	// Protected indicates whether the repository is protected.
	Protected bool
	// AllowRepositoryDeletion indicates whether the repository should be protected from deletion.
	AllowRepositoryDeletion bool
	// RetainOnDelete indicates whether the repository should be retained on deletion.
	RetainOnDelete *bool
	// PulumiOptions are additional options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new GitLab repository with the specified options.
// ctx: The Pulumi context.
// name: The logical name for the Pulumi resource.
// opts: The options for creating the repository.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*gitlab.Project, error) {
	visibility := defaults.GetOrDefault(opts.Visibility, defaultVisibility)
	visibilitySelector := "enabled"
	wikiVisibilitySelector := "disabled"

	optsWithRepoSpecifics := append([]pulumi.ResourceOption{}, opts.PulumiOptions...)
	optsWithRepoSpecifics = append(
		optsWithRepoSpecifics,
		pulumi.Protect(!opts.AllowRepositoryDeletion),
		pulumi.RetainOnDelete(defaults.GetOrDefault(opts.RetainOnDelete, true)),
		pulumi.IgnoreChanges([]string{}),
	)

	sort.Strings(opts.Topics)

	if !utilgitlab.IsPrivateRepository(visibility) {
		visibilitySelector = "private"
	}
	if opts.EnableWiki != nil && *opts.EnableWiki {
		wikiVisibilitySelector = visibility
	}

	return gitlab.NewProject(ctx, fmt.Sprintf("gitlab-project-%s", name), &gitlab.ProjectArgs{
		Name:                                   opts.Name,
		Description:                            opts.Description,
		AllowMergeOnSkippedPipeline:            pulumi.Bool(true),
		AnalyticsAccessLevel:                   pulumi.String(visibilitySelector),
		ArchiveOnDestroy:                       pulumi.Bool(opts.Protected),
		Archived:                               pulumi.Bool(false),
		AutoCancelPendingPipelines:             pulumi.String("enabled"),
		AutoDevopsEnabled:                      pulumi.Bool(defaults.GetOrDefault(opts.AutoDevopsEnabled, false)),
		AutocloseReferencedIssues:              pulumi.Bool(true),
		BuildGitStrategy:                       pulumi.String("fetch"),
		BuildsAccessLevel:                      pulumi.String(visibilitySelector),
		CiDefaultGitDepth:                      pulumi.Int(defaultCiDefaultGitDepth),
		CiDeletePipelinesInSeconds:             pulumi.Int(defaultCiDeletePipelinesInSeconds),
		CiForwardDeploymentEnabled:             pulumi.Bool(true),
		CiForwardDeploymentRollbackAllowed:     pulumi.Bool(false),
		CiPipelineVariablesMinimumOverrideRole: pulumi.String("owner"),
		CiPushRepositoryForJobTokenAllowed:     pulumi.Bool(true),
		ContainerRegistryAccessLevel:           pulumi.String(visibilitySelector),
		EnvironmentsAccessLevel:                pulumi.String(visibilitySelector),
		FeatureFlagsAccessLevel:                pulumi.String(visibilitySelector),
		ForkingAccessLevel:                     pulumi.String(visibilitySelector),
		GroupRunnersEnabled:                    pulumi.Bool(true),
		InfrastructureAccessLevel:              pulumi.String(visibilitySelector),
		IssuesAccessLevel:                      pulumi.String(visibilitySelector),
		KeepLatestArtifact:                     pulumi.Bool(true),
		MergeMethod:                            pulumi.String("ff"),
		MergePipelinesEnabled:                  pulumi.Bool(defaults.GetOrDefault(opts.EnableMergeQueue, false)),
		MergeRequestsAccessLevel:               pulumi.String(visibilitySelector),
		MergeTrainsEnabled:                     pulumi.Bool(defaults.GetOrDefault(opts.EnableMergeQueue, false)),
		MergeTrainsSkipTrainAllowed:            pulumi.Bool(true),
		ModelExperimentsAccessLevel:            pulumi.String("disabled"),
		ModelRegistryAccessLevel:               pulumi.String("disabled"),
		MonitorAccessLevel:                     pulumi.String(visibilitySelector),
		NamespaceId:                            opts.NamespaceID,
		OnlyAllowMergeIfAllDiscussionsAreResolved: pulumi.Bool(
			defaults.GetOrDefault(opts.ConversationResolution, true),
		),
		OnlyAllowMergeIfPipelineSucceeds: pulumi.Bool(true),
		PackagesEnabled:                  pulumi.Bool(true),
		PagesAccessLevel:                 pulumi.String(visibilitySelector),
		PrintingMergeRequestLinkEnabled:  pulumi.Bool(true),
		PublicJobs:                       pulumi.Bool(!utilgitlab.IsPrivateRepository(visibility)),
		ReleasesAccessLevel:              pulumi.String(visibilitySelector),
		RemoveSourceBranchAfterMerge:     pulumi.Bool(true),
		RepositoryAccessLevel:            pulumi.String(visibilitySelector),
		RequestAccessEnabled:             pulumi.Bool(false),
		RequirementsAccessLevel:          pulumi.String(visibilitySelector),
		ResolveOutdatedDiffDiscussions:   pulumi.Bool(true),
		SecurityAndComplianceAccessLevel: pulumi.String(visibilitySelector),
		SharedRunnersEnabled:             pulumi.Bool(true),
		SnippetsAccessLevel:              pulumi.String("disabled"),
		SquashOption:                     pulumi.String("never"),
		Topics:                           pulumi.ToStringArray(opts.Topics),
		VisibilityLevel:                  pulumi.String(visibility),
		WikiAccessLevel:                  pulumi.String(wikiVisibilitySelector),
	}, optsWithRepoSpecifics...)
}
