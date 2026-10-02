package repository

import (
	"fmt"
	"sort"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/pulumi/pulumi-gitlab/sdk/v10/go/gitlab"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
	utilgitlab "github.com/muhlba91/pulumi-shared-library/pkg/util/gitlab"
)

const (
	// defaultCiDefaultGitDepth is the default git depth of CI jobs.
	defaultCiDefaultGitDepth = 1
	// defaultCiDeletePipelinesInSeconds is the default age after which pipelines are deleted (1 year in seconds).
	defaultCiDeletePipelinesInSeconds = 31536000
	// defaultVisibility is the default visibility level of the repository.
	defaultVisibility = "public"
)

// CreateOptions defines the options for creating a GitLab repository.
type CreateOptions struct {
	// Name is the name of the repository.
	Name pulumi.StringInput
	// Description is the description of the repository.
	Description pulumi.StringInput
	// NamespaceID is the ID of the namespace under which the repository will be created (group or user). Optional.
	NamespaceID pulumi.IntPtrInput
	// EnableWiki indicates whether to enable the wiki for the repository. Optional, defaults to false.
	EnableWiki *bool
	// Topics is a list of topics to associate with the repository.
	Topics []string
	// Visibility is the visibility level of the repository ("public", "internal", or "private"). Optional, defaults to "public".
	// Feature access is restricted to project members unless the repository is "private" or "internal".
	Visibility *string
	// ConversationResolution indicates whether to require conversation resolution. Optional, defaults to true.
	ConversationResolution *bool
	// AutoDevopsEnabled indicates whether to enable Auto DevOps for the repository. Optional, defaults to false.
	AutoDevopsEnabled *bool
	// EnableMergeQueue indicates whether to enable merged results pipelines and merge trains. Optional, defaults to false.
	EnableMergeQueue *bool
	// DeletePipelinesInSeconds is the number of seconds after which pipelines should be automatically deleted.
	// Optional, defaults to 1 year.
	DeletePipelinesInSeconds *int
	// Protected indicates whether the repository is archived instead of deleted on destroy.
	Protected bool
	// AllowRepositoryDeletion indicates whether the repository may be deleted.
	// If false, the Pulumi resource is protected from deletion.
	AllowRepositoryDeletion bool
	// RetainOnDelete indicates whether the repository should be retained on deletion. Optional, defaults to true.
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
		Name:                        opts.Name,
		Description:                 opts.Description,
		AllowMergeOnSkippedPipeline: pulumi.Bool(true),
		AnalyticsAccessLevel:        pulumi.String(visibilitySelector),
		ArchiveOnDestroy:            pulumi.Bool(opts.Protected),
		Archived:                    pulumi.Bool(false),
		AutoCancelPendingPipelines:  pulumi.String("enabled"),
		AutoDevopsEnabled:           pulumi.Bool(defaults.GetOrDefault(opts.AutoDevopsEnabled, false)),
		AutocloseReferencedIssues:   pulumi.Bool(true),
		BuildGitStrategy:            pulumi.String("fetch"),
		BuildsAccessLevel:           pulumi.String(visibilitySelector),
		CiDefaultGitDepth:           pulumi.Int(defaultCiDefaultGitDepth),
		CiDeletePipelinesInSeconds: pulumi.Int(
			defaults.GetOrDefault(opts.DeletePipelinesInSeconds, defaultCiDeletePipelinesInSeconds),
		),
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
