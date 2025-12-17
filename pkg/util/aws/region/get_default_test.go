package region_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/aws/region"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestGetOrDefault_WithExplicitRegion(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		reg := "explicit-region"
		out := region.GetOrDefault(ctx, &reg)
		assert.Equal(t, "explicit-region", *out)
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestGetOrDefault_NoRegionReturnsNil(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		reg := ""
		out := region.GetOrDefault(ctx, &reg)
		assert.Nil(t, out)
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestGetOrDefault_NilRegionReturnsNil(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		var reg *string
		out := region.GetOrDefault(ctx, reg)
		assert.Nil(t, out)
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestGetOrDefault_UsesAWSConfigWhenPresent(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		out := region.GetOrDefault(ctx, nil)
		assert.Equal(t, "configured-region", *out)
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)), mocks.WithConfig(map[string]string{
		"aws:region": "configured-region",
	}))
	require.NoError(t, err)
}
