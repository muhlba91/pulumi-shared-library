package setupkey_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/lib/netbird/setupkey"
	rModel "github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

const (
	day         = 86400
	setupKeyTok = "netbird:index/setupKey:SetupKey"
	rotatingTok = "time:index/rotating:Rotating"
)

func ptr[T any](v T) *T { return &v }

func TestCreate_Defaults(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		key, err := setupkey.Create(ctx, "test", &setupkey.CreateOptions{Type: pulumi.String("reusable")})
		require.NoError(t, err)
		require.NotNil(t, key)

		key.Name.ApplyT(func(v string) error { assert.Equal(t, "test", v); return nil })
		key.Type.ApplyT(func(v string) error { assert.Equal(t, "reusable", v); return nil })
		key.UsageLimit.ApplyT(func(v int) error { assert.Equal(t, 0, v); return nil })
		key.Ephemeral.ApplyT(func(v bool) error { assert.False(t, v); return nil })
		key.ExpirySeconds.ApplyT(func(v int) error { assert.Equal(t, 365*day, v); return nil })
		key.Key.ApplyT(func(v string) error {
			assert.Equal(t, "mocked-netbird-setup-key-netbird-setup-key-test", v)
			return nil
		})
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Equal(t, []string{"netbird-setup-key-test"}, counter.Resources[setupKeyTok])
	assert.Nil(t, counter.Resources[rotatingTok])
}

func TestCreate_CustomOptions(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		key, err := setupkey.Create(ctx, "custom", &setupkey.CreateOptions{
			Type:          pulumi.String("one-off"),
			AutoGroups:    pulumi.StringArray{pulumi.String("g1"), pulumi.String("g2")},
			UsageLimit:    ptr(5),
			Ephemeral:     ptr(true),
			ExpirySeconds: ptr(7 * day),
		})
		require.NoError(t, err)

		key.Type.ApplyT(func(v string) error { assert.Equal(t, "one-off", v); return nil })
		key.AutoGroups.ApplyT(func(v []string) error { assert.Equal(t, []string{"g1", "g2"}, v); return nil })
		key.UsageLimit.ApplyT(func(v int) error { assert.Equal(t, 5, v); return nil })
		key.Ephemeral.ApplyT(func(v bool) error { assert.True(t, v); return nil })
		key.ExpirySeconds.ApplyT(func(v int) error { assert.Equal(t, 7*day, v); return nil })
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)
}

func TestCreate_ExpiryResolution(t *testing.T) {
	tests := []struct {
		name     string
		expiry   *int
		rotation *rModel.Options
		want     int
	}{
		{"no rotation, no expiry", nil, nil, 365 * day},
		{"rotation default days", nil, &rModel.Options{}, 60 * day},
		{"rotation 30 days", nil, &rModel.Options{Days: 30}, 60 * day},
		{"rotation clamped to max", nil, &rModel.Options{Days: 200}, 365 * day},
		{"rotation at max", nil, &rModel.Options{Days: 365}, 365 * day},
		{"explicit overrides rotation", ptr(10 * day), &rModel.Options{Days: 30}, 10 * day},
		{"explicit clamped to min", ptr(60), nil, day},
		{"explicit clamped to max", ptr(1000 * day), nil, 365 * day},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := pulumi.RunErr(func(ctx *pulumi.Context) error {
				key, err := setupkey.Create(ctx, "exp", &setupkey.CreateOptions{
					Type:          pulumi.String("reusable"),
					ExpirySeconds: tc.expiry,
					Rotation:      tc.rotation,
				})
				require.NoError(t, err)
				key.ExpirySeconds.ApplyT(func(v int) error { assert.Equal(t, tc.want, v); return nil })
				return nil
			}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
			require.NoError(t, err)
		})
	}
}

func TestCreate_Rotation(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		_, err := setupkey.Create(ctx, "rot", &setupkey.CreateOptions{
			Type:     pulumi.String("reusable"),
			Rotation: &rModel.Options{Days: 7},
		})
		return err
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Len(t, counter.Resources[setupKeyTok], 1)
	assert.Equal(t, []string{"rotation-netbird-setup-key-rot"}, counter.Resources[rotatingTok])
}

func TestCreate_RotationTooLong(t *testing.T) {
	counter := mocks.NewCounter()

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		key, err := setupkey.Create(ctx, "long", &setupkey.CreateOptions{
			Type:     pulumi.String("reusable"),
			Rotation: &rModel.Options{Days: 366},
		})
		require.ErrorContains(t, err, "must not exceed 365")
		assert.Nil(t, key)
		return nil
	}, pulumi.WithMocks("project", "stack", counter))
	require.NoError(t, err)

	assert.Nil(t, counter.Resources[setupKeyTok])
	assert.Nil(t, counter.Resources[rotatingTok])
}
