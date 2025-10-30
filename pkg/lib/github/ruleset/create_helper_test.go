package ruleset //nolint:testpackage // keep test in same package to access unexported functions

import (
	"testing"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildMergeQueueArgs_EnabledDisabled(t *testing.T) {
	falseVal := false
	opts := &CreateOptions{
		EnableMergeQueue: &falseVal,
	}
	mq := buildMergeQueueArgs(opts)
	assert.Nil(t, mq)

	trueVal := true
	opts.EnableMergeQueue = &trueVal
	mq = buildMergeQueueArgs(opts)
	require.NotNil(t, mq)

	assert.Equal(t, pulumi.Int(60), mq.CheckResponseTimeoutMinutes)
	assert.Equal(t, pulumi.String("ALLGREEN"), mq.GroupingStrategy)
	assert.Equal(t, pulumi.String("REBASE"), mq.MergeMethod)
	assert.Equal(t, pulumi.Int(5), mq.MaxEntriesToBuild)
}

func TestBuildBypassActorsArgs_DefaultAndWithIntegrations(t *testing.T) {
	falseVal := false
	opts := &CreateOptions{
		AllowBypass: &falseVal,
	}
	bypass := buildBypassActorsArgs(opts)
	assert.Empty(t, bypass)

	opts.AllowBypass = nil
	bypass = buildBypassActorsArgs(opts)
	require.Len(t, bypass, 2)

	bypass[0].ToRepositoryRulesetBypassActorOutput().ActorId().ApplyT(func(v *int) error {
		assert.Equal(t, 2, *v)
		return nil
	})
	bypass[0].ToRepositoryRulesetBypassActorOutput().ActorType().ApplyT(func(s string) error {
		assert.Equal(t, "RepositoryRole", s)
		return nil
	})
	bypass[1].ToRepositoryRulesetBypassActorOutput().ActorId().ApplyT(func(v *int) error {
		assert.Equal(t, 5, *v)
		return nil
	})
	bypass[1].ToRepositoryRulesetBypassActorOutput().BypassMode().ToStringOutput().ApplyT(func(s string) error {
		assert.Equal(t, "always", s)
		return nil
	})

	opts.AllowBypass = nil
	opts.AllowBypassIntegrations = []int{10, 3}
	bypass = buildBypassActorsArgs(opts)
	require.Len(t, bypass, 4)
	exp := []int{2, 3, 5, 10}
	for i := range bypass {
		idx := i
		bypass[idx].ToRepositoryRulesetBypassActorOutput().ActorId().ApplyT(func(v *int) error {
			assert.Equal(t, exp[idx], *v)
			return nil
		})
	}
}

func TestBuildRequiredStatusChecksArgs_DefaultsAndCustom(t *testing.T) {
	opts := &CreateOptions{
		RequiredChecks: []string{},
		WIPIntegration: nil,
	}
	req := buildRequiredStatusChecksArgs(opts)
	require.NotNil(t, req)
	require.Len(t, req.RequiredChecks, 1)
	req.ToRepositoryRulesetRulesRequiredStatusChecksOutput().
		RequiredChecks().
		ApplyT(func(cs []github.RepositoryRulesetRulesRequiredStatusChecksRequiredCheck) error {
			assert.Equal(t, "WIP", cs[0].Context)
			return nil
		})
	req.ToRepositoryRulesetRulesRequiredStatusChecksOutput().
		RequiredChecks().
		ApplyT(func(cs []github.RepositoryRulesetRulesRequiredStatusChecksRequiredCheck) error {
			assert.Equal(t, 3414, *cs[0].IntegrationId)
			return nil
		})
	req.ToRepositoryRulesetRulesRequiredStatusChecksOutput().
		StrictRequiredStatusChecksPolicy().
		ApplyT(func(b *bool) error {
			assert.True(t, *b)
			return nil
		})

	falseVal := false
	opts = &CreateOptions{
		RequiredChecks: []string{"z", "a"},
		WIPIntegration: &falseVal,
	}
	req = buildRequiredStatusChecksArgs(opts)
	require.NotNil(t, req)
	require.Len(t, req.RequiredChecks, 2)
	req.ToRepositoryRulesetRulesRequiredStatusChecksOutput().
		RequiredChecks().
		ApplyT(func(cs []github.RepositoryRulesetRulesRequiredStatusChecksRequiredCheck) error {
			assert.Equal(t, "a", cs[0].Context)
			return nil
		})
	req.ToRepositoryRulesetRulesRequiredStatusChecksOutput().
		RequiredChecks().
		ApplyT(func(cs []github.RepositoryRulesetRulesRequiredStatusChecksRequiredCheck) error {
			assert.Equal(t, 15368, *cs[0].IntegrationId)
			return nil
		})
	req.ToRepositoryRulesetRulesRequiredStatusChecksOutput().
		RequiredChecks().
		ApplyT(func(cs []github.RepositoryRulesetRulesRequiredStatusChecksRequiredCheck) error {
			assert.Equal(t, "z", cs[1].Context)
			return nil
		})
}
