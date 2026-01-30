package record

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/domain"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/sanitize"
)

const defaultTTL = 300

// CreateOptions holds optional parameters for Create.
type CreateOptions struct {
	// Domain is the DNS record's domain (e.g. "web.example.com").
	Domain string
	// Zone is the managed zone name (e.g. "example.com").
	Zone string
	// RecordType is the DNS record type (e.g. "A", "TXT").
	RecordType string
	// Record is the data for the record.
	Record pulumi.StringInput
	// TTL for the record. If 0, defaultTTL is used.
	TTL int
	// Project is an optional Scaleway project to set on the Record.
	Project *string
	// NameAppendix is an optional appendix to append to the resource name (for example for uniqueness).
	NameAppendix *string
	// PulumiOptions are optional resource options passed to the Record.
	PulumiOptions []pulumi.ResourceOption
}

// Create creates a Scaleway DNS Record with the given parameters.
// ctx: Pulumi context.
// opts: CreateOptions with parameters for the Record.
func Create(
	ctx *pulumi.Context,
	opts *CreateOptions,
) (*domain.Record, error) {
	ttl := defaultTTL
	if opts.TTL != 0 {
		ttl = opts.TTL
	}

	name := fmt.Sprintf("dns-record-%s-%s", sanitize.Text(opts.RecordType), sanitize.Text(opts.Domain))
	if opts.NameAppendix != nil && *opts.NameAppendix != "" {
		name = fmt.Sprintf("%s-%s", name, sanitize.Text(*opts.NameAppendix))
	}

	domainName := strings.ReplaceAll(opts.Domain, fmt.Sprintf(".%s", opts.Zone), "")
	if domainName == opts.Zone {
		domainName = ""
	}

	return domain.NewRecord(ctx, name, &domain.RecordArgs{
		DnsZone:   pulumi.String(opts.Zone),
		Name:      pulumi.String(domainName),
		Type:      pulumi.String(opts.RecordType),
		Data:      opts.Record,
		Ttl:       pulumi.Int(ttl),
		ProjectId: pulumi.StringPtrFromPtr(opts.Project),
	}, opts.PulumiOptions...)
}
