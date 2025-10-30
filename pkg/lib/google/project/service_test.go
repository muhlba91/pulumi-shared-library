package project_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	gproject "github.com/muhlba91/pulumi-shared-library/pkg/lib/google/project"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestEnableServices_Single(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		args := &gproject.EnableServicesArgs{
			Project: "proj-single",
			Services: []string{
				"compute.googleapis.com",
			},
		}

		res, err := gproject.EnableServices(ctx, args)
		require.NoError(t, err)
		require.Len(t, res, 1)

		res[0].Service.ApplyT(func(s string) error {
			assert.Equal(t, "compute.googleapis.com", s)
			return nil
		})
		res[0].Project.ApplyT(func(p string) error {
			assert.Equal(t, "proj-single", p)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestEnableServices_Multiple(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		services := []string{
			"compute.googleapis.com",
			"iam.googleapis.com",
			"storage.googleapis.com",
		}
		args := &gproject.EnableServicesArgs{
			Project:  "proj-multi",
			Services: services,
		}

		res, err := gproject.EnableServices(ctx, args)
		require.NoError(t, err)
		require.Len(t, res, len(services))

		for i, svc := range services {
			// capture index/svc for closure
			idx := i
			expected := svc
			res[idx].Service.ApplyT(func(s string) error {
				assert.Equal(t, expected, s)
				return nil
			})
			res[idx].Project.ApplyT(func(p string) error {
				assert.Equal(t, "proj-multi", p)
				return nil
			})
		}
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestEnableServices_Single_WithOptionalArgs(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		args := &gproject.EnableServicesArgs{
			Project: "proj-single",
			Services: []string{
				"compute.googleapis.com",
			},
			PulumiOptions: []pulumi.ResourceOption{},
		}

		res, err := gproject.EnableServices(ctx, args)
		require.NoError(t, err)
		require.Len(t, res, 1)

		res[0].Service.ApplyT(func(s string) error {
			assert.Equal(t, "compute.googleapis.com", s)
			return nil
		})
		res[0].Project.ApplyT(func(p string) error {
			assert.Equal(t, "proj-single", p)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
