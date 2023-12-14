package strings

// Contains checks if a string is present in a slice of strings.
//
// haystack: slice of strings to search
// needle: string to search for
func Contain(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}

	return false
}
