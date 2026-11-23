package project_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/google/project"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestGetNumber(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		number := project.GetNumber(ctx, pulumi.String("proj-123"))
		number.ApplyT(func(n string) error {
			assert.Equal(t, "123456789", n)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}
