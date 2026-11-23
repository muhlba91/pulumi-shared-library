package secret

import (
	"fmt"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateOptions defines the options for creating a GitHub Actions secret.
type CreateOptions struct {
	// Key is the name of the secret.
	Key string
	// Value is the value of the secret.
	Value pulumi.StringInput
	// Repository is the GitHub repository to which the secret will be added.
	Repository *github.Repository
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// Create stores a value in GitHub Actions secrets.
// The secret is a child of the repository, depends on it, and is deleted before it is replaced.
// Returns an output of the secret; a failed creation fails the output and thereby the Pulumi program.
// ctx: Pulumi context.
// opts: CreateOptions containing the key, value, and repository.
func Create(
	ctx *pulumi.Context,
	opts *CreateOptions,
) (pulumi.Output, error) {
	if opts.Repository == nil {
		return nil, ErrMissingRepository
	}

	pulumiOpts := opts.PulumiOptions
	if pulumiOpts == nil {
		pulumiOpts = []pulumi.ResourceOption{
			pulumi.DependsOn([]pulumi.Resource{opts.Repository}),
		}
	} else {
		pulumiOpts = append(pulumiOpts, pulumi.DependsOn([]pulumi.Resource{opts.Repository}))
	}
	pulumiOpts = append(pulumiOpts, pulumi.Parent(opts.Repository))
	pulumiOpts = append(pulumiOpts, pulumi.DeleteBeforeReplace(true), pulumi.IgnoreChanges([]string{"remoteUpdatedAt"}))

	return opts.Repository.Name.ApplyT(func(repositoryName string) (*github.ActionsSecret, error) {
		name := fmt.Sprintf("github-actions-secret-%s-%s", repositoryName, opts.Key)

		as, err := github.NewActionsSecret(ctx, name, &github.ActionsSecretArgs{
			Repository: pulumi.String(repositoryName),
			SecretName: pulumi.String(opts.Key),
			Value:      opts.Value,
		}, pulumiOpts...)
		if err != nil {
			return nil, fmt.Errorf("failed to create the GitHub Actions secret %q: %w", opts.Key, err)
		}

		return as, nil
	}), nil
}
