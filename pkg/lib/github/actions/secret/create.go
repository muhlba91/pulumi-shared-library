package secret

import (
	"fmt"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rs/zerolog/log"
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
// ctx: Pulumi context.
// opts: CreateOptions containing the key, value, and repository.
func Create(
	ctx *pulumi.Context,
	opts *CreateOptions,
) pulumi.Output {
	pulumiOpts := opts.PulumiOptions
	if pulumiOpts == nil {
		pulumiOpts = []pulumi.ResourceOption{
			pulumi.DependsOn([]pulumi.Resource{opts.Repository}),
		}
	} else {
		pulumiOpts = append(pulumiOpts, pulumi.DependsOn([]pulumi.Resource{opts.Repository}))
	}
	pulumiOpts = append(pulumiOpts, pulumi.DeleteBeforeReplace(true), pulumi.IgnoreChanges([]string{"remoteUpdatedAt"}))

	return opts.Repository.Name.ApplyT(func(repositoryName string) *github.ActionsSecret {
		name := fmt.Sprintf("github-actions-secret-%s-%s", repositoryName, opts.Key)

		as, err := github.NewActionsSecret(ctx, name, &github.ActionsSecretArgs{
			Repository:     pulumi.String(repositoryName),
			SecretName:     pulumi.String(opts.Key),
			PlaintextValue: opts.Value,
			DestroyOnDrift: pulumi.Bool(false),
		}, pulumiOpts...)
		if err != nil {
			log.Error().Msgf("Failed to create GitHub Actions secret: %v", err)
			return nil
		}

		return as
	})
}
