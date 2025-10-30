package convert_test

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/pulumi/convert"
)

func TestIDToInt(t *testing.T) {
	idOut := pulumi.ID("123").ToIDOutput()
	intOut := convert.IDToInt(idOut)

	intOut.ApplyT(func(v int) error {
		assert.Equal(t, 123, v)
		return nil
	})
}

func TestIDToInt_NonNumeric(t *testing.T) {
	idOut := pulumi.ID("not-an-int").ToIDOutput()
	intOut := convert.IDToInt(idOut)

	intOut.ApplyT(func(v int) error {
		assert.Equal(t, 0, v)
		return nil
	})
}
