package mocks

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// WithConfig returns a RunOption that sets the configuration values
// for the Pulumi program.
// config: A map of configuration keys to values.
func WithConfig(config map[string]string) pulumi.RunOption {
	return func(info *pulumi.RunInfo) {
		info.Config = config
	}
}
