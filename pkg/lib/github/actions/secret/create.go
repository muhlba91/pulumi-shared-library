package secret

import (
	"fmt"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rs/zerolog/log"
)

// WriteArgs defines the input arguments for the Write function.
type WriteArgs struct {
	// Key is the name of the secret.
	Key string
	// Value is the value of the secret.
	Value pulumi.StringInput
	// Repository is the GitHub repository to which the secret will be added.
	Repository *github.Repository
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// Write stores a value in GitHub Actions secrets.
// ctx: Pulumi context.
// args: WriteArgs containing the key, value, and repository.
func Write(
	ctx *pulumi.Context,
	args *WriteArgs,
) pulumi.Output {
	opts := args.PulumiOptions
	if opts == nil {
		opts = []pulumi.ResourceOption{
			pulumi.DependsOn([]pulumi.Resource{args.Repository}),
		}
	} else {
		opts = append(opts, pulumi.DependsOn([]pulumi.Resource{args.Repository}))
	}

	return args.Repository.Name.ApplyT(func(repositoryName string) *github.ActionsSecret {
		name := fmt.Sprintf("github-actions-secret-%s-%s", repositoryName, args.Key)

		as, err := github.NewActionsSecret(ctx, name, &github.ActionsSecretArgs{
			Repository:     pulumi.String(repositoryName),
			SecretName:     pulumi.String(args.Key),
			PlaintextValue: args.Value,
		}, opts...)
		if err != nil {
			log.Error().Msgf("Failed to create GitHub Actions secret: %v", err)
			return nil
		}

		return as
	})
}
