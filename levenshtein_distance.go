package ed

func LevenshteinDistance(a, b string) int {
	if len(a) == 0 && len(b) != 0 {
		return len(b)
	}

	if len(a) != 0 && len(b) == 0 {
		return len(a)
	}

	if a == b {
		return 0
	}

	if a[0] == b[0] {
		return LevenshteinDistance(a[1:], b[1:])
	}

	return 1 + min(
		LevenshteinDistance(a[1:], b),
		LevenshteinDistance(a, b[1:]),
		LevenshteinDistance(a[1:], b[1:]),
	)
}

func min(a, b, c int) int {
	if a > b {
		a = b
	}
	if a > c {
		a = c
	}
	return a
}
