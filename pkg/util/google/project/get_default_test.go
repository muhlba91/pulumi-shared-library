package project_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/google/project"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestGetOrDefault_WithExplicitProject(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		proj := "explicit-project"
		out := project.GetOrDefault(ctx, &proj)
		assert.Equal(t, "explicit-project", *out)
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestGetOrDefault_NoProjectReturnsNil(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		proj := ""
		out := project.GetOrDefault(ctx, &proj)
		assert.Nil(t, out)
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestGetOrDefault_NilProjectReturnsNil(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		var proj *string
		out := project.GetOrDefault(ctx, proj)
		assert.Nil(t, out)
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestGetOrDefault_UsesGCPConfigWhenPresent(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		out := project.GetOrDefault(ctx, nil)
		assert.Equal(t, "configured-project", *out)
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)), mocks.WithConfig(map[string]string{
		"gcp:project": "configured-project",
	}))
	require.NoError(t, err)
}
