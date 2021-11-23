package utils

// Longest returns the length of the longest string in the given array
func Longest(strs []string) int {
	longest := 0
	for _, str := range strs {
		if len(str) > longest {
			longest = len(str)
		}
	}
	return longest
}

// Max returns the largest of the two ints.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the smallest of the two ints.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
