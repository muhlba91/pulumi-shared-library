package rotation

// Options defines the options for creating a rotation resource.
type Options struct {
	// Name is the name of the rotation resource, which is named "rotation-<name>".
	// Optional when used through the rotation util, which then uses the name of the rotated resource.
	Name *string
	// Days is the number of days between rotations. If days is <= 0 it defaults to 30.
	Days int
}
