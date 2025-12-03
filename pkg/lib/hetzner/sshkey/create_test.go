package sshkey_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libssh "github.com/muhlba91/pulumi-shared-library/pkg/lib/hetzner/sshkey"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateSSHKey(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &libssh.CreateOptions{
			Name:      "mykey",
			PublicKey: pulumi.String("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCu..."),
			Labels:    map[string]string{"env": "test"},
		}

		res, err := libssh.Create(ctx, "sshkey", opts)
		require.NoError(t, err)
		require.NotNil(t, res)

		res.Name.ApplyT(func(n string) error {
			assert.Contains(t, n, opts.Name)
			return nil
		})
		res.PublicKey.ApplyT(func(pk string) error {
			assert.NotEmpty(t, pk)
			return nil
		})
		res.Labels.ApplyT(func(m map[string]string) error {
			assert.Equal(t, opts.Labels, m)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreateSSHKey_WithOptions(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &libssh.CreateOptions{
			Name:      "protected-key",
			PublicKey: pulumi.String("ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIE..."),
			Labels:    map[string]string{"team": "infra"},
			PulumiOptions: []pulumi.ResourceOption{
				pulumi.Protect(true),
			},
		}

		res, err := libssh.Create(ctx, "sshkey", opts)
		require.NoError(t, err)
		require.NotNil(t, res)

		res.Name.ApplyT(func(n string) error {
			assert.Contains(t, n, opts.Name)
			return nil
		})
		res.PublicKey.ApplyT(func(pk string) error {
			assert.NotEmpty(t, pk)
			return nil
		})
		res.Labels.ApplyT(func(m map[string]string) error {
			assert.Equal(t, opts.Labels, m)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
