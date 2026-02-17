package secret

import (
	"fmt"

	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rs/zerolog/log"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
)

// CreateOptions defines the options for creating a GitLab project variable.
type CreateOptions struct {
	// Key is the name of the secret.
	Key string
	// Value is the value of the secret.
	Value pulumi.StringInput
	// Repository is the GitLab project to which the secret will be added.
	Repository *gitlab.Project
	// VariableType is the type of the variable (e.g., "env_var", "file"). Optional, defaults to "env_var".
	VariableType *string
	// Protected indicates whether the secret should be protected and be available on protected branches and tags.
	Protected *bool
	// DisableVariableExpansion indicates whether to disable variable expansion for the secret.
	DisableVariableExpansion *bool
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// Create stores a value in a GitLab Variable.
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
	pulumiOpts = append(pulumiOpts, pulumi.DeleteBeforeReplace(true))

	return opts.Repository.Name.ApplyT(func(repositoryName string) *gitlab.ProjectVariable {
		name := fmt.Sprintf("gitlab-project-variable-%s-%s", repositoryName, opts.Key)

		as, err := gitlab.NewProjectVariable(ctx, name, &gitlab.ProjectVariableArgs{
			Project:      opts.Repository.ID(),
			Key:          pulumi.String(opts.Key),
			Value:        opts.Value,
			VariableType: pulumi.String(defaults.GetOrDefault(opts.VariableType, "env_var")),
			Hidden:       pulumi.Bool(true),
			Masked:       pulumi.Bool(true),
			Protected:    pulumi.Bool(defaults.GetOrDefault(opts.Protected, false)),
			Raw:          pulumi.Bool(defaults.GetOrDefault(opts.DisableVariableExpansion, false)),
		}, pulumiOpts...)
		if err != nil {
			log.Error().Msgf("Failed to create GitLab Project Variable: %v", err)
			return nil
		}

		return as
	})
}
