package tokens

// Estimate returns an approximate token count for the given data.
func Estimate(data []byte) int {
	return len(data) / 4
}
