package project

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/projects"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// EnableServicesArgs represents the arguments for enabling services.
type EnableServicesArgs struct {
	// Project is the GCP project ID.
	Project string
	// Services is the list of services to enable.
	Services []string
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// EnableServices enables the specified Google services for a given project.
// ctx: Pulumi context.
// args: EnableProjectServicesArgs containing project and services to enable.
func EnableServices(ctx *pulumi.Context, args *EnableServicesArgs) ([]*projects.Service, error) {
	res := make([]*projects.Service, 0, len(args.Services))

	for _, svc := range args.Services {
		name := fmt.Sprintf("gcp-project-service-%s-%s", args.Project, svc)
		s, err := projects.NewService(ctx, name, &projects.ServiceArgs{
			Service: pulumi.String(svc),
			Project: pulumi.String(args.Project),
		}, args.PulumiOptions...)
		if err != nil {
			return nil, err
		}
		res = append(res, s)
	}

	return res, nil
}
