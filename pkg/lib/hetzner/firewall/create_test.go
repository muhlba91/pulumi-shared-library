package firewall_test

import (
	"testing"

	"github.com/pulumi/pulumi-hcloud/sdk/go/hcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libfw "github.com/muhlba91/pulumi-shared-library/pkg/lib/hetzner/firewall"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateFirewall(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &libfw.CreateOptions{
			Name: "fw-basic",
			Rules: []libfw.Rule{
				{
					Direction: "in",
					Protocol:  "tcp",
					Port:      "22",
					SourceIps: []pulumi.StringInput{pulumi.String("203.0.113.0/24")},
				},
				{
					Direction: "in",
					Protocol:  "tcp",
					Port:      "80",
				},
			},
			Labels: map[string]string{"env": "test"},
		}

		fw, err := libfw.Create(ctx, opts)
		require.NoError(t, err)
		require.NotNil(t, fw)

		fw.Name.ApplyT(func(n string) error {
			assert.Equal(t, "fw-basic", n)
			return nil
		})
		fw.Labels.ApplyT(func(m map[string]string) error {
			assert.Equal(t, opts.Labels, m)
			return nil
		})
		fw.Rules.ApplyT(func(r []hcloud.FirewallRule) error {
			assert.Len(t, r, 2)
			assert.Equal(t, "tcp", r[0].Protocol)
			assert.Equal(t, "22", *r[0].Port)
			assert.Equal(t, "tcp", r[1].Protocol)
			assert.Equal(t, "80", *r[1].Port)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateFirewall_WithOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &libfw.CreateOptions{
			Name: "fw-protected",
			Rules: []libfw.Rule{
				{
					Direction: "in",
					Protocol:  "udp",
					Port:      "53",
					SourceIps: []pulumi.StringInput{pulumi.String("198.51.100.0/24")},
				},
			},
			Labels: map[string]string{"team": "infra"},
			PulumiOptions: []pulumi.ResourceOption{
				pulumi.Protect(true),
			},
		}

		fw, err := libfw.Create(ctx, opts)
		require.NoError(t, err)
		require.NotNil(t, fw)

		fw.Labels.ApplyT(func(m map[string]string) error {
			assert.Equal(t, opts.Labels, m)
			return nil
		})
		fw.Rules.ApplyT(func(r []hcloud.FirewallRule) error {
			assert.Len(t, r, 1)
			assert.Equal(t, "udp", r[0].Protocol)
			assert.Equal(t, "53", *r[0].Port)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
