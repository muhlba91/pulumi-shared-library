package metadata

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

// LabelsToStringMap converts a map of string labels to a Pulumi StringMap.
// Returns nil if the labels are nil.
// labels: map of string labels.
func LabelsToStringMap(labels map[string]string) pulumi.StringMap {
	var result pulumi.StringMap

	if labels != nil {
		result = pulumi.StringMap{}
		for k, v := range labels {
			result[k] = pulumi.String(v)
		}
	}

	return result
}
