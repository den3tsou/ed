package ed

func LevenshteinDistance(a, b string) int {
	al := len(a)
	bl := len(b)
	r0 := make([]int, bl+1)
	r1 := make([]int, bl+1)

	for i := 0; i < bl+1; i++ {
		r0[i] = i
	}

	for i := 0; i < al; i++ {
		r1[0] = i + 1

		for j := 0; j < bl; j++ {
			deletionCost := r0[j+1] + 1
			insertionCost := r1[j] + 1
			substitutionCost := r0[j]
			if a[i] != b[j] {
				substitutionCost += 1
			}

			r1[j+1] = min(deletionCost, insertionCost, substitutionCost)
		}

		for i := 0; i < bl+1; i++ {
			r0[i] = r1[i]
		}
	}

	return r0[bl]
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
