package setupkey

import (
	"fmt"

	"github.com/KitStream/netbird-pulumi-provider/sdk/go/netbird"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	rModel "github.com/muhlba91/pulumi-shared-library/pkg/model/rotation"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/defaults"
	"github.com/muhlba91/pulumi-shared-library/pkg/util/rotation"
)

const (
	// secondsPerDay is the number of seconds in a day.
	secondsPerDay = 24 * 60 * 60
	// minExpirySeconds is the minimum expiry NetBird accepts for a setup key (1 day).
	minExpirySeconds = secondsPerDay
	// maxExpirySeconds is the maximum expiry NetBird accepts for a setup key (365 days).
	// It is also the expiry used if neither an explicit expiry nor a rotation is set.
	maxExpirySeconds = 365 * secondsPerDay

	// maxRotationDays is the maximum rotation period in days.
	// It equals the maximum expiry, as a longer period would let the key expire before it is rotated.
	maxRotationDays = 365
	// defaultRotationDays is the rotation period in days used if the rotation options do not set a positive value.
	// It mirrors the default of the rotation package.
	defaultRotationDays = 30
	// expiryRotationFactor is the multiplier applied to the rotation period to derive the expiry.
	// The key therefore does not expire before the next rotation, even if a rotation is delayed by up to one period.
	expiryRotationFactor = 2

	// defaultUsageLimit is the default usage limit for a setup key (0 is unlimited).
	defaultUsageLimit = 0
	// defaultEphemeral is the default ephemeral setting for a setup key (false).
	defaultEphemeral = false
)

// CreateOptions defines the options for creating a NetBird setup key.
type CreateOptions struct {
	// Type is the setup key type ("one-off" or "reusable"). Required.
	// The value is not validated by the library, but by the NetBird provider.
	Type pulumi.StringInput
	// AutoGroups are the IDs of the groups automatically assigned to peers enrolled with this key. Optional.
	AutoGroups pulumi.StringArrayInput
	// UsageLimit is the maximum number of uses (0 is unlimited). Defaults to 0.
	UsageLimit *int
	// ExpirySeconds is the validity of the key in seconds and takes precedence over the rotation period.
	// Values outside the range NetBird accepts (1 to 365 days) are clamped to the nearest limit.
	// Defaults to twice the rotation period (clamped to 365 days) if rotation is set, otherwise 365 days.
	ExpirySeconds *int
	// Ephemeral marks enrolled peers as ephemeral, i.e., they are removed after a period of inactivity. Defaults to false.
	Ephemeral *bool
	// Rotation defines the rotation options for the resource. Optional.
	// If set, the setup key is replaced after the rotation period, which must not exceed 365 days.
	Rotation *rModel.Options
	// PulumiOptions are additional options to pass to the Pulumi resource.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a new NetBird setup key.
// Returns an error if the rotation period exceeds the maximum expiry of 365 days.
// ctx: The Pulumi context.
// name: The logical name for the Pulumi resource (prefixed with "netbird-setup-key-"), also used as the setup key name.
// opts: The options for creating the setup key. Must not be nil.
func Create(ctx *pulumi.Context, name string, opts *CreateOptions) (*netbird.SetupKey, error) {
	resName := fmt.Sprintf("netbird-setup-key-%s", name)

	expiry, err := expirySeconds(opts)
	if err != nil {
		return nil, err
	}

	pulumiOpts := append([]pulumi.ResourceOption{}, opts.PulumiOptions...)
	if trigger, _ := rotation.Trigger(ctx, resName, opts.Rotation); trigger != nil {
		pulumiOpts = append(pulumiOpts, pulumi.ReplacementTrigger(trigger))
	}

	return netbird.NewSetupKey(ctx, resName, &netbird.SetupKeyArgs{
		Name:          pulumi.String(name),
		Type:          opts.Type.ToStringOutput().ToStringPtrOutput(),
		AutoGroups:    opts.AutoGroups,
		UsageLimit:    pulumi.Int(defaults.GetOrDefault(opts.UsageLimit, defaultUsageLimit)),
		Ephemeral:     pulumi.Bool(defaults.GetOrDefault(opts.Ephemeral, defaultEphemeral)),
		ExpirySeconds: pulumi.Int(expiry),
	}, pulumiOpts...)
}

// expirySeconds resolves the expiry of the key in seconds and clamps it to the range NetBird accepts (1 to 365 days).
// An explicit value takes precedence over the value derived from the rotation period.
// Returns an error if the rotation period exceeds the maximum rotation days, even if an explicit expiry is set.
// opts: The options for creating the setup key.
func expirySeconds(opts *CreateOptions) (int, error) {
	derived := maxExpirySeconds
	if opts.Rotation != nil {
		days := opts.Rotation.Days
		if days <= 0 {
			days = defaultRotationDays
		}
		if days > maxRotationDays {
			return 0, fmt.Errorf("rotation days must not exceed %d, got %d", maxRotationDays, days)
		}
		derived = days * secondsPerDay * expiryRotationFactor
	}

	return min(max(defaults.GetOrDefault(opts.ExpirySeconds, derived), minExpirySeconds), maxExpirySeconds), nil
}
