package repository

import (
	"fmt"
	"sort"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
	utilgithub "github.com/muhlba91/pulumi-shared-library/pkg/util/github"
)

const defaultGitHubPagesBranch = "main"

// CreateOptions defines the options for creating a GitHub repository.
type CreateOptions struct {
	// Name is the name of the repository.
	Name pulumi.StringInput
	// Description is the description of the repository.
	Description pulumi.StringInput
	// EnableDiscussions indicates whether to enable discussions for the repository.
	EnableDiscussions pulumi.BoolPtrInput
	// EnableWiki indicates whether to enable the wiki for the repository.
	EnableWiki pulumi.BoolPtrInput
	// Homepage is the homepage URL of the repository.
	Homepage pulumi.StringPtrInput
	// Topics is a list of topics to associate with the repository.
	Topics []string
	// GitHubPagesBranch is the branch to use for GitHub Pages.
	GitHubPagesBranch *string
	// Visibility is the visibility level of the repository. Can be "public" or "private".
	Visibility *string
	// Protected indicates whether the repository is protected.
	Protected bool
	// AllowRepositoryDeletion indicates whether the repository should be protected from deletion.
	AllowRepositoryDeletion bool
	// RetainOnDelete indicates whether the repository should be retained on deletion.
	RetainOnDelete *bool
	// PulumiOptions are additional options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new GitHub repository with the specified options.
// ctx: The Pulumi context.
// name: The logical name for the Pulumi resource.
// opts: The options for creating the repository.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*github.Repository, error) {
	defaultVisibility := "public"

	optsWithRepoSpecifics := append([]pulumi.ResourceOption{}, opts.PulumiOptions...)
	optsWithRepoSpecifics = append(
		optsWithRepoSpecifics,
		pulumi.Protect(!opts.AllowRepositoryDeletion),
		pulumi.RetainOnDelete(defaults.GetOrDefault(opts.RetainOnDelete, true)),
		pulumi.IgnoreChanges([]string{
			"securityAndAnalysis",
			"template",
		}),
	)

	sort.Strings(opts.Topics)

	var pages *github.RepositoryPagesArgs
	var secAnalysis *github.RepositorySecurityAndAnalysisArgs
	if !utilgithub.IsPrivateRepository(opts.Visibility) {
		pages = &github.RepositoryPagesArgs{
			BuildType: pulumi.String("workflow"),
			Source: &github.RepositoryPagesSourceArgs{
				Branch: pulumi.String(defaults.GetOrDefault(opts.GitHubPagesBranch, defaultGitHubPagesBranch)),
				Path:   pulumi.String("/"),
			},
		}
		secAnalysis = &github.RepositorySecurityAndAnalysisArgs{
			SecretScanning: &github.RepositorySecurityAndAnalysisSecretScanningArgs{
				Status: pulumi.String("enabled"),
			},
			SecretScanningPushProtection: &github.RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs{
				Status: pulumi.String("enabled"),
			},
		}
	}

	return github.NewRepository(ctx, fmt.Sprintf("github-repo-%s", name), &github.RepositoryArgs{
		Name:                     opts.Name,
		Description:              opts.Description,
		HasDiscussions:           opts.EnableDiscussions,
		HasWiki:                  opts.EnableWiki,
		HomepageUrl:              opts.Homepage,
		Topics:                   pulumi.ToStringArray(opts.Topics),
		Visibility:               pulumi.String(defaults.GetOrDefault(opts.Visibility, defaultVisibility)),
		AllowAutoMerge:           pulumi.Bool(false),
		AllowMergeCommit:         pulumi.Bool(false),
		AllowRebaseMerge:         pulumi.Bool(true),
		AllowSquashMerge:         pulumi.Bool(false),
		AllowUpdateBranch:        pulumi.Bool(true),
		Archived:                 pulumi.Bool(false),
		ArchiveOnDestroy:         pulumi.Bool(opts.Protected),
		AutoInit:                 pulumi.Bool(false),
		DeleteBranchOnMerge:      pulumi.Bool(true),
		HasDownloads:             pulumi.Bool(true),
		HasIssues:                pulumi.Bool(true),
		HasProjects:              pulumi.Bool(true),
		MergeCommitMessage:       pulumi.String("PR_TITLE"),
		MergeCommitTitle:         pulumi.String("MERGE_MESSAGE"),
		SquashMergeCommitMessage: pulumi.String("COMMIT_MESSAGES"),
		SquashMergeCommitTitle:   pulumi.String("COMMIT_OR_PR_TITLE"),
		VulnerabilityAlerts:      pulumi.Bool(true),
		Pages:                    pages,
		SecurityAndAnalysis:      secAnalysis,
	}, optsWithRepoSpecifics...)
}
