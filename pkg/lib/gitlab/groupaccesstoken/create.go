package groupaccesstoken

import (
	"fmt"
	"slices"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab"
)

const (
	defaultExpirationDays               = 365
	defaultRotationBeforeExpirationDays = 30
)

// CreateOptions defines the options for creating a GitLab group access token.
type CreateOptions struct {
	// Name is the name of the group access token.
	Name pulumi.StringInput
	// Description is the description of the group access token.
	Description pulumi.StringInput
	// Group is the ID or full path of the group for which the access token will be created.
	Group string
	// Scopes is a list of scopes to assign to the group access token.
	Scopes []string
	// PulumiOptions are additional options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new GitLab group access token with the specified options.
// ctx: The Pulumi context.
// name: The logical name for the Pulumi resource.
// opts: The options for creating the group access token.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*gitlab.GroupAccessToken, error) {
	scopes := make([]string, 0, len(opts.Scopes)+1)
	scopes = append(scopes, opts.Scopes...)
	scopes = append(scopes, "self_rotate")
	slices.Sort(scopes)

	return gitlab.NewGroupAccessToken(ctx, fmt.Sprintf("gitlab-gat-%s", name), &gitlab.GroupAccessTokenArgs{
		Name:        opts.Name,
		Description: opts.Description,
		Group:       pulumi.String(opts.Group),
		Scopes:      pulumi.ToStringArray(scopes),
		AccessLevel: pulumi.String("maintainer"),
		RotationConfiguration: &gitlab.GroupAccessTokenRotationConfigurationArgs{
			ExpirationDays:   pulumi.Int(defaultExpirationDays),
			RotateBeforeDays: pulumi.Int(defaultRotationBeforeExpirationDays),
		},
	}, opts.PulumiOptions...)
}
