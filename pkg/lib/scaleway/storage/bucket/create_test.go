package bucket_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/object"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libbucket "github.com/muhlba91/pulumi-shared-library/pkg/lib/scaleway/storage/bucket"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateBucket_Basic(t *testing.T) {
	req := require.New(t)
	assert := assert.New(t)

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &libbucket.CreateOptions{
			Location: pulumi.String("EU"),
			Labels:   map[string]string{"env": "test"},
		}

		b, err := libbucket.Create(ctx, "mybucket", opts)
		req.NoError(err)
		req.NotNil(b)

		b.Region.ApplyT(func(loc *string) error {
			assert.Equal("EU", *loc)
			return nil
		})
		b.CorsRules.ApplyT(func(cors []object.BucketCorsRule) error {
			assert.Empty(cors)
			return nil
		})
		b.LifecycleRules.ApplyT(func(rules []object.BucketLifecycleRule) error {
			assert.Len(rules, 2)

			rule1 := rules[0]
			assert.True(rule1.Enabled)
			assert.Nil(rule1.Transitions)
			assert.Equal("expire-incomplete-multipart-uploads", *rule1.Prefix)
			assert.Equal(3, *rule1.AbortIncompleteMultipartUploadDays)

			rule2 := rules[1]
			assert.True(rule2.Enabled)
			assert.Equal("move-to-one-zone", *rule2.Prefix)
			assert.Nil(rule2.AbortIncompleteMultipartUploadDays)
			assert.Len(rule2.Transitions, 1)
			transition := rule2.Transitions[0]
			assert.Equal("ONEZONE_IA", transition.StorageClass)
			assert.Equal(30, *transition.Days)

			return nil
		})
		b.Tags.ApplyT(func(m map[string]string) error {
			assert.Equal(map[string]string{"env": "test"}, m)
			return nil
		})
		b.ID().ApplyT(func(id string) error {
			assert.NotEmpty(id)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	req.NoError(err)
}

func TestCreateBucket_WithOptions(t *testing.T) {
	req := require.New(t)
	assert := assert.New(t)

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		maxAgeSeconds := 3600
		transitionDays := 5

		opts := &libbucket.CreateOptions{
			Location:              pulumi.String("US"),
			OneZoneTransitionDays: &transitionDays,
			CORS: &libbucket.CreateCorsOptions{
				MaxAgeSeconds:  &maxAgeSeconds,
				Method:         []string{"GET", "POST"},
				Origin:         []string{"https://example.com"},
				ResponseHeader: []string{"Content-Type"},
			},
			Labels:        map[string]string{"owner": "ci"},
			PulumiOptions: []pulumi.ResourceOption{pulumi.Protect(true)},
		}

		b, err := libbucket.Create(ctx, "withopts", opts)
		req.NoError(err)
		req.NotNil(b)

		b.Region.ApplyT(func(loc *string) error {
			assert.Equal("US", *loc)
			return nil
		})
		b.CorsRules.ApplyT(func(cors []object.BucketCorsRule) error {
			assert.Len(cors, 1)
			rule := cors[0]
			assert.Equal(3600, *rule.MaxAgeSeconds)
			assert.Equal([]string{"GET", "POST"}, rule.AllowedMethods)
			assert.Equal([]string{"https://example.com"}, rule.AllowedOrigins)
			assert.Equal([]string{"Content-Type"}, rule.AllowedHeaders)
			return nil
		})
		b.LifecycleRules.ApplyT(func(rules []object.BucketLifecycleRule) error {
			assert.Len(rules, 2)

			rule2 := rules[1]
			assert.True(rule2.Enabled)
			assert.Equal("move-to-one-zone", *rule2.Prefix)
			assert.Nil(rule2.AbortIncompleteMultipartUploadDays)
			assert.Len(rule2.Transitions, 1)
			transition := rule2.Transitions[0]
			assert.Equal("ONEZONE_IA", transition.StorageClass)
			assert.Equal(transitionDays, *transition.Days)

			return nil
		})
		b.Tags.ApplyT(func(m map[string]string) error {
			assert.Equal(map[string]string{"owner": "ci"}, m)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.NewCounter()))
	req.NoError(err)
}
