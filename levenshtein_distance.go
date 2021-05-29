package ed

func LevenshteinDistance(a, b string) int {
	al := len(a) + 1
	bl := len(b) + 1
	d := make([][]int, al)
	for i := range d {
		d[i] = make([]int, bl)
	}

	for i := 1; i < al; i++ {
		d[i][0] = i
	}

	for j := 1; j < bl; j++ {
		d[0][j] = j
	}

	var cost int
	for j := 1; j < bl; j++ {
		for i := 1; i < al; i++ {
			if a[i-1] == b[j-1] {
				cost = 0
			} else {
				cost = 1
			}
			d[i][j] = min(
				d[i-1][j]+1,
				d[i][j-1]+1,
				d[i-1][j-1]+cost,
			)
		}
	}

	return d[al-1][bl-1]
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
