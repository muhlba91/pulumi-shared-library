package github

// IsPrivateRepository checks if the given visibility indicates a private repository.
// Only the exact value "private" does, a nil visibility does not.
// visibility: A pointer to a string representing the repository visibility.
func IsPrivateRepository(visibility *string) bool {
	if visibility != nil && *visibility == "private" {
		return true
	}

	return false
}
