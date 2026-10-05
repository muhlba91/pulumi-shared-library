package policy_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/netbird/policy"
	pModel "github.com/muhlba91/pulumi-shared-library/pkg/model/netbird/policy"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

const policyTok = "netbird:index/policy:Policy"

func ptr[T any](v T) *T { return &v }

func TestCreate_Defaults(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		p, err := policy.Create(ctx, "test", &policy.CreateOptions{
			Rules: []pModel.Rule{{Name: "rule"}},
		})
		require.NoError(t, err)
		require.NotNil(t, p)

		p.Name.ApplyT(func(v string) error { assert.Equal(t, "test", v); return nil })
		p.Enabled.ApplyT(func(v bool) error { assert.True(t, v); return nil })
		p.Rule.Action().ApplyT(func(v *string) error { assert.Equal(t, "accept", *v); return nil })
		p.Rule.Enabled().ApplyT(func(v *bool) error { assert.True(t, *v); return nil })
		p.Rule.Bidirectional().ApplyT(func(v *bool) error { assert.True(t, *v); return nil })
		p.Rule.Protocol().ApplyT(func(v *string) error { assert.Equal(t, "all", *v); return nil })
		p.Rule.Ports().ApplyT(func(v []string) error { assert.Empty(t, v); return nil })
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Equal(t, []string{"netbird-policy-test"}, counter.Resources[policyTok])
}

func TestCreate_CustomOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		p, err := policy.Create(ctx, "custom", &policy.CreateOptions{
			Description: ptr("desc"),
			Enabled:     ptr(false),
			Rules: []pModel.Rule{{
				Name:          "rule",
				Action:        ptr("drop"),
				Enabled:       ptr(false),
				Bidirectional: ptr(false),
				Protocol:      ptr("tcp"),
				Ports:         []string{"80", "443"},
				Sources:       pulumi.StringArray{pulumi.String("s1")},
				Destinations:  pulumi.StringArray{pulumi.String("d1")},
			}},
		})
		require.NoError(t, err)

		p.Description.ApplyT(func(v string) error { assert.Equal(t, "desc", v); return nil })
		p.Enabled.ApplyT(func(v bool) error { assert.False(t, v); return nil })
		p.Rule.Action().ApplyT(func(v *string) error { assert.Equal(t, "drop", *v); return nil })
		p.Rule.Enabled().ApplyT(func(v *bool) error { assert.False(t, *v); return nil })
		p.Rule.Bidirectional().ApplyT(func(v *bool) error { assert.False(t, *v); return nil })
		p.Rule.Protocol().ApplyT(func(v *string) error { assert.Equal(t, "tcp", *v); return nil })
		p.Rule.Ports().ApplyT(func(v []string) error { assert.Equal(t, []string{"80", "443"}, v); return nil })
		p.Rule.Sources().ApplyT(func(v []string) error { assert.Equal(t, []string{"s1"}, v); return nil })
		p.Rule.Destinations().ApplyT(func(v []string) error { assert.Equal(t, []string{"d1"}, v); return nil })
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}

func TestCreate_NoRules(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		p, err := policy.Create(ctx, "empty", &policy.CreateOptions{})
		require.NoError(t, err)
		require.NotNil(t, p)
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	require.NoError(t, err)
}

func TestCreate_TooManyRules(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		p, err := policy.Create(ctx, "many", &policy.CreateOptions{
			Rules: []pModel.Rule{{Name: "a"}, {Name: "b"}},
		})
		require.ErrorContains(t, err, "at most 1")
		assert.Nil(t, p)
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Nil(t, counter.Resources[policyTok])
}

func TestCreate_PortsRequireTCPOrUDP(t *testing.T) {
	for _, protocol := range []*string{nil, ptr("all"), ptr("icmp")} {
		counter := mocks.NewCounter()

		err := pulumi.RunErr(func(ctx *pulumi.Context) error {
			p, err := policy.Create(ctx, "ports", &policy.CreateOptions{
				Rules: []pModel.Rule{{Name: "rule", Protocol: protocol, Ports: []string{"443"}}},
			})
			require.ErrorContains(t, err, "ports require protocol")
			assert.Nil(t, p)
			return nil
		}, pulumi.WithMocks("project", "stack", counter))
		require.NoError(t, err)

		assert.Nil(t, counter.Resources[policyTok])
	}
}
