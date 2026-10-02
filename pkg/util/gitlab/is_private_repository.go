package gitlab

// IsPrivateRepository checks if the given visibility indicates a repository that is not public.
// Both "private" and "internal" do.
// visibility: The repository visibility.
func IsPrivateRepository(visibility string) bool {
	return visibility == "private" || visibility == "internal"
}
