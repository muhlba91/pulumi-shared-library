package gitlab_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	utilgitlab "github.com/muhlba91/pulumi-shared-library/pkg/util/gitlab"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestGetCurrentUserID(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		got := utilgitlab.GetCurrentUserID(ctx)
		require.NotNil(t, got)
		assert.Equal(t, 1, *got)
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	assert.NoError(t, err)
}
