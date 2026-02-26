package rotation

// Options defines the options for creating a rotation resource.
type Options struct {
	// Name prefix for the rotation resource.
	Name *string
	// Number of days between rotations. If days is <= 0 it defaults to 30.
	Days int
}
