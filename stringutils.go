package coinpayments

// stringExistsInSlice is a helper to determine whether an element exists in a slice of string
func stringExistsInSlice(strings []string, e string) bool {
	for _, a := range strings {
		if a == e {
			return true
		}
	}
	return false
}
