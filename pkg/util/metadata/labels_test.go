package metadata_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/metadata"
)

func TestLabelsToStringMap_Nil(t *testing.T) {
	var input map[string]string
	got := metadata.LabelsToStringMap(input)
	assert.Nil(t, got)
}

func TestLabelsToStringMap_Empty(t *testing.T) {
	input := map[string]string{}
	got := metadata.LabelsToStringMap(input)
	assert.NotNil(t, got)
	assert.Empty(t, got)
}

func TestLabelsToStringMap_MultipleEntries(t *testing.T) {
	input := map[string]string{
		"env":   "prod",
		"owner": "team-a",
	}
	got := metadata.LabelsToStringMap(input)
	assert.NotNil(t, got)
	assert.Len(t, got, len(input))
	assert.Equal(t, pulumi.String("prod"), got["env"])
	assert.Equal(t, pulumi.String("team-a"), got["owner"])
}

func TestLabelsToStringMap_InputIsolation(t *testing.T) {
	input := map[string]string{"k": "v"}
	got := metadata.LabelsToStringMap(input)

	input["k"] = "changed"
	assert.Equal(t, pulumi.String("v"), got["k"])
}
