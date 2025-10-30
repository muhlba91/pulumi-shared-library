package defaults

// GetOrDefault returns the value pointed to by v if v is non-nil, otherwise it returns def.
// v: Pointer to a value of any type.
// def: Default value to return if v is nil.
func GetOrDefault[T any](v *T, def T) T {
	if v == nil {
		return def
	}
	return *v
}
