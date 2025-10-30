package convert

import (
	"strconv"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// IDToInt converts a Pulumi IDOutput to an IntOutput by parsing the string ID to an integer.
// id: The Pulumi IDOutput to convert.
func IDToInt(id pulumi.IDOutput) pulumi.IntOutput {
	cID, _ := id.ApplyT(func(id string) int {
		intID, _ := strconv.Atoi(id)
		return intID
	}).(pulumi.IntOutput)
	return cID
}
