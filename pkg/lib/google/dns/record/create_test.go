package record_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	librecord "github.com/muhlba91/pulumi-shared-library/pkg/lib/google/dns/record"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreate_DefaultTTLAndNoProject(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		domain := "example.com"
		zone := pulumi.String("zone-1")
		recType := "A"
		records := pulumi.StringArray{pulumi.String("1.2.3.4")}

		rs, err := librecord.Create(ctx, &librecord.CreateOptions{
			Domain:     domain,
			ZoneID:     zone,
			RecordType: recType,
			Records:    records,
			TTL:        0,
			Project:    nil,
		})
		require.NoError(t, err)
		require.NotNil(t, rs)

		rs.ManagedZone.ApplyT(func(mz string) error {
			assert.Equal(t, "zone-1", mz)
			return nil
		})
		rs.Name.ApplyT(func(n string) error {
			assert.Equal(t, "example.com.", n)
			return nil
		})
		rs.Type.ApplyT(func(typ string) error {
			assert.Equal(t, recType, typ)
			return nil
		})
		rs.Rrdatas.ApplyT(func(r []string) error {
			assert.Equal(t, []string{"1.2.3.4"}, r)
			return nil
		})
		rs.Ttl.ApplyT(func(tt *int) error {
			assert.Equal(t, 300, *tt)
			return nil
		})
		rs.Project.ApplyT(func(p string) error {
			assert.Empty(t, p)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreate_CustomTTLAndProject(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		domain := "test.local"
		zone := pulumi.String("zone-2")
		recType := "TXT"
		records := pulumi.StringArray{pulumi.String("v=spf1 include:_spf.example.com ~all")}
		ttl := 600
		proj := "proj-123"

		rs, err := librecord.Create(ctx, &librecord.CreateOptions{
			Domain:     domain,
			ZoneID:     zone,
			RecordType: recType,
			Records:    records,
			TTL:        ttl,
			Project:    &proj,
		})
		require.NoError(t, err)
		require.NotNil(t, rs)

		rs.Type.ApplyT(func(typ string) error {
			assert.Equal(t, recType, typ)
			return nil
		})
		rs.Rrdatas.ApplyT(func(r []string) error {
			assert.Equal(t, []string{"v=spf1 include:_spf.example.com ~all"}, r)
			return nil
		})
		rs.Ttl.ApplyT(func(tt *int) error {
			assert.Equal(t, ttl, *tt)
			return nil
		})
		rs.Project.ApplyT(func(p string) error {
			assert.Equal(t, proj, p)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}

func TestCreate_DefaultTTLAndNoProject_WithOptionalArguments(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		domain := "example.com"
		zone := pulumi.String("zone-1")
		recType := "A"
		records := pulumi.StringArray{pulumi.String("1.2.3.4")}

		rs, err := librecord.Create(ctx, &librecord.CreateOptions{
			Domain:        domain,
			ZoneID:        zone,
			RecordType:    recType,
			Records:       records,
			PulumiOptions: []pulumi.ResourceOption{},
		})
		require.NoError(t, err)
		require.NotNil(t, rs)

		rs.ManagedZone.ApplyT(func(mz string) error {
			assert.Equal(t, "zone-1", mz)
			return nil
		})
		rs.Name.ApplyT(func(n string) error {
			assert.Equal(t, "example.com.", n)
			return nil
		})
		rs.Type.ApplyT(func(typ string) error {
			assert.Equal(t, recType, typ)
			return nil
		})
		rs.Rrdatas.ApplyT(func(r []string) error {
			assert.Equal(t, []string{"1.2.3.4"}, r)
			return nil
		})
		rs.Ttl.ApplyT(func(tt *int) error {
			assert.Equal(t, 300, *tt)
			return nil
		})
		rs.Project.ApplyT(func(p string) error {
			assert.Empty(t, p)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	require.NoError(t, err)
}
