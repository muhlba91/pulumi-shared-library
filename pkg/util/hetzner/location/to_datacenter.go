package location

const defaultDatacenter = "fsn1-dc14"

// ToDatacenter converts a Hetzner location key to the datacenter identifier.
// If the location is unknown it returns the default datacenter.
// location: e.g. "fsn1", "nbg1"
func ToDatacenter(location *string) string {
	if location == nil {
		return defaultDatacenter
	}

	locationsDatacenter := map[string]string{
		"fsn1": "fsn1-dc14",
		"nbg1": "nbg1-dc3",
	}

	if dc, ok := locationsDatacenter[*location]; ok {
		return dc
	}

	return defaultDatacenter
}
