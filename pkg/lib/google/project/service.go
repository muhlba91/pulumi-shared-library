package project

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/projects"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// EnableServicesOptions represents the options for enabling services.
type EnableServicesOptions struct {
	// Project is the GCP project ID.
	Project string
	// Services is the list of services to enable.
	Services []string
	// PulumiOptions are additional Pulumi resource options. Optional.
	PulumiOptions []pulumi.ResourceOption
}

// EnableServices enables the specified Google services for a given project.
// ctx: Pulumi context.
// opts: EnableServicesOptions containing project and services to enable.
func EnableServices(ctx *pulumi.Context, opts *EnableServicesOptions) ([]*projects.Service, error) {
	res := make([]*projects.Service, 0, len(opts.Services))

	for _, svc := range opts.Services {
		name := fmt.Sprintf("gcp-project-service-%s-%s", opts.Project, svc)
		s, err := projects.NewService(ctx, name, &projects.ServiceArgs{
			Service: pulumi.String(svc),
			Project: pulumi.String(opts.Project),
		}, opts.PulumiOptions...)
		if err != nil {
			return nil, err
		}
		res = append(res, s)
	}

	return res, nil
}
