package accesskey

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	rModel "github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/rotation"
)

// CreateOptions defines the options for creating an IAM access key.
type CreateOptions struct {
	// UserName is the name of the IAM user to create the access key for.
	UserName string
	// User is the IAM user resource to associate the access key with.
	User pulumi.Resource
	// Rotation defines the rotation options for the access key.
	Rotation *rModel.Options
	// PulumiOptions are additional options to pass to the resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new IAM access key for the specified user.
// ctx: Pulumi context.
// opts: CreateOptions for creating the access key.
func Create(ctx *pulumi.Context, opts *CreateOptions) (*iam.AccessKey, error) {
	resName := fmt.Sprintf("aws-access-key-%s", opts.UserName)

	optsWithDepends := append([]pulumi.ResourceOption{}, opts.PulumiOptions...)
	optsWithDepends = append(optsWithDepends, pulumi.DependsOn([]pulumi.Resource{opts.User}))

	if trigger, _ := rotation.Trigger(ctx, resName, opts.Rotation); trigger != nil {
		optsWithDepends = append(optsWithDepends, pulumi.ReplacementTrigger(trigger))
	}

	return iam.NewAccessKey(
		ctx,
		resName,
		&iam.AccessKeyArgs{
			User: pulumi.String(opts.UserName),
		},
		optsWithDepends...,
	)
}
