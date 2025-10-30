package github

// IsPrivateRepository checks if the given visibility string indicates a private repository.
// visibility: A pointer to a string representing the repository visibility.
func IsPrivateRepository(visibility *string) bool {
	if visibility != nil && *visibility == "private" {
		return true
	}

	return false
}
