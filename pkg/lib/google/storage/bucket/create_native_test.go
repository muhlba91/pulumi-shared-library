package bucket_test

import (
	"testing"

	storage "github.com/pulumi/pulumi-google-native/sdk/go/google/storage/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libbucket "github.com/muhlba91/pulumi-shared-library/pkg/lib/google/storage/bucket"
	"github.com/muhlba91/pulumi-shared-library/test/mocks"
)

func TestCreateNativeBucket_Basic(t *testing.T) {
	req := require.New(t)
	assert := assert.New(t)

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &libbucket.CreateNativeOptions{
			Location: pulumi.String("EU"),
			Labels:   map[string]string{"env": "test"},
		}

		b, err := libbucket.CreateNative(ctx, "mybucket", opts)
		req.NoError(err)
		req.NotNil(b)

		b.Location.ApplyT(func(loc string) error {
			assert.Equal("EU", loc)
			return nil
		})

		b.Labels.ApplyT(func(m map[string]string) error {
			assert.Equal(map[string]string{"env": "test"}, m)
			return nil
		})

		b.ID().ApplyT(func(id string) error {
			assert.NotEmpty(id)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	req.NoError(err)
}

func TestCreateNativeBucket_WithOptions(t *testing.T) {
	req := require.New(t)
	assert := assert.New(t)

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		opts := &libbucket.CreateNativeOptions{
			Location:      pulumi.String("US"),
			Labels:        map[string]string{"owner": "ci"},
			PulumiOptions: []pulumi.ResourceOption{pulumi.Protect(true)},
		}

		b, err := libbucket.CreateNative(ctx, "withopts", opts)
		req.NoError(err)
		req.NotNil(b)

		b.Location.ApplyT(func(loc string) error {
			assert.Equal("US", loc)
			return nil
		})
		b.Labels.ApplyT(func(m map[string]string) error {
			assert.Equal(map[string]string{"owner": "ci"}, m)
			return nil
		})

		b.StorageClass.ApplyT(func(sc string) error {
			assert.Equal("STANDARD", sc)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	req.NoError(err)
}

func TestCreateNativeBucket_WithCORS(t *testing.T) {
	req := require.New(t)
	assert := assert.New(t)

	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		maxAge := 3600
		opts := &libbucket.CreateNativeOptions{
			Location: pulumi.String("US"),
			CORS: &libbucket.CreateNativeCorsOptions{
				MaxAgeSeconds:  &maxAge,
				Method:         []string{"GET", "POST"},
				Origin:         []string{"*"},
				ResponseHeader: []string{"Content-Type"},
			},
		}

		b, err := libbucket.CreateNative(ctx, "withcors", opts)
		req.NoError(err)
		req.NotNil(b)

		b.Cors.ApplyT(func(cors []storage.BucketCorsItemResponse) error {
			req.Len(cors, 1)
			assert.Equal(maxAge, cors[0].MaxAgeSeconds)
			assert.ElementsMatch([]string{"GET", "POST"}, cors[0].Method)
			assert.ElementsMatch([]string{"*"}, cors[0].Origin)
			assert.ElementsMatch([]string{"Content-Type"}, cors[0].ResponseHeader)
			return nil
		})

		return nil
	}, pulumi.WithMocks("project", "stack", mocks.Mocks(0)))
	req.NoError(err)
}
