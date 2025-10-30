package record

import (
	"fmt"

	gcpdns "github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/dns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

const defaultTTL = 300

// CreateOptions holds optional parameters for Create.
type CreateOptions struct {
	// Domain is the DNS record's domain (e.g. "example.com").
	Domain string
	// ZoneId is the managed zone id.
	ZoneID pulumi.StringInput
	// RecordType is the DNS record type (e.g. "A", "TXT").
	RecordType string
	// Records are the rrdatas for the record.
	Records pulumi.StringArrayInput
	// TTL for the record. If 0, defaultTTL is used.
	TTL int
	// Project is an optional GCP project to set on the RecordSet.
	Project *string
	// PulumiOptions are optional resource options passed to the RecordSet.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a GCP DNS RecordSet with the given parameters.
// ctx: Pulumi context.
// opts: CreateOptions with parameters for the RecordSet.
func Create(
	ctx *pulumi.Context,
	opts *CreateOptions,
) (*gcpdns.RecordSet, error) {
	ttl := defaultTTL
	if opts.TTL != 0 {
		ttl = opts.TTL
	}
	name := fmt.Sprintf("dns-record-%s-%s", sanitize.Text(opts.RecordType), sanitize.Text(opts.Domain))

	return gcpdns.NewRecordSet(ctx, name, &gcpdns.RecordSetArgs{
		ManagedZone: opts.ZoneID,
		Name:        pulumi.String(fmt.Sprintf("%s.", opts.Domain)),
		Type:        pulumi.String(opts.RecordType),
		Rrdatas:     opts.Records,
		Ttl:         pulumi.Int(ttl),
		Project:     pulumi.StringPtrFromPtr(opts.Project),
	}, opts.PulumiOptions...)
}
