package project

import (
	"github.com/pulumi/pulumi-gcp/sdk/v10/go/gcp/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GetNumber looks up the project number of a GCP project.
// ctx: Pulumi context.
// project: The project ID.
// opts: Invoke options, e.g. a provider. Optional.
func GetNumber(ctx *pulumi.Context, project pulumi.StringInput, opts ...pulumi.InvokeOption) pulumi.StringOutput {
	return organizations.LookupProjectOutput(ctx, organizations.LookupProjectOutputArgs{
		ProjectId: project.ToStringOutput().ToStringPtrOutput(),
	}, opts...).Number()
}
