package accesstoken

import (
	"fmt"
	"slices"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab"

	utilgitlab "github.com/muhlba91/pulumi-shared-library/pkg/util/gitlab"
)

// CreateOptions defines the options for creating a GitLab personal access token.
type CreateOptions struct {
	// Name is the name of the personal access token.
	Name pulumi.StringInput
	// Description is the description of the personal access token.
	Description pulumi.StringInput
	// UserID is the ID of the user for whom the personal access token is being created.
	UserID *int
	// Scopes is a list of scopes to assign to the personal access token.
	Scopes []string
	// PulumiOptions are additional options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new GitLab personal access token with the specified options.
// ctx: The Pulumi context.
// name: The logical name for the Pulumi resource.
// opts: The options for creating the personal access token.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*gitlab.PersonalAccessToken, error) {
	scopes := make([]string, 0, len(opts.Scopes)+1)
	scopes = append(scopes, opts.Scopes...)
	scopes = append(scopes, "self_rotate")
	slices.Sort(scopes)

	userID := utilgitlab.GetCurrentUserID(ctx)
	if opts.UserID != nil {
		userID = opts.UserID
	}

	return gitlab.NewPersonalAccessToken(ctx, fmt.Sprintf("gitlab-pat-%s", name), &gitlab.PersonalAccessTokenArgs{
		Name:        opts.Name,
		Description: opts.Description,
		UserId:      pulumi.Int(*userID),
		Scopes:      pulumi.ToStringArray(scopes),
	}, opts.PulumiOptions...)
}
