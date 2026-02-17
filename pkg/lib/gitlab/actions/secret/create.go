package secret

import (
	"fmt"

	"github.com/pulumi/pulumi-gitlab/sdk/v9/go/gitlab"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/rs/zerolog/log"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
)

// WriteArgs defines the input arguments for the Write function.
type WriteArgs struct {
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

// Write stores a value in a GitLab Variable.
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
	opts = append(opts, pulumi.DeleteBeforeReplace(true))

	return args.Repository.Name.ApplyT(func(repositoryName string) *gitlab.ProjectVariable {
		name := fmt.Sprintf("gitlab-project-variable-%s-%s", repositoryName, args.Key)

		as, err := gitlab.NewProjectVariable(ctx, name, &gitlab.ProjectVariableArgs{
			Project:      args.Repository.ID(),
			Key:          pulumi.String(args.Key),
			Value:        args.Value,
			VariableType: pulumi.String(defaults.GetOrDefault(args.VariableType, "env_var")),
			Hidden:       pulumi.Bool(true),
			Masked:       pulumi.Bool(true),
			Protected:    pulumi.Bool(defaults.GetOrDefault(args.Protected, false)),
			Raw:          pulumi.Bool(defaults.GetOrDefault(args.DisableVariableExpansion, false)),
		}, opts...)
		if err != nil {
			log.Error().Msgf("Failed to create GitLab Project Variable: %v", err)
			return nil
		}

		return as
	})
}
