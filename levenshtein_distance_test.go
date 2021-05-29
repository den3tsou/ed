package ed_test

import (
	"testing"

	"github.com/den3tsou/ed"
)

func TestLevenshteinDistance(t *testing.T) {
	for _, tc := range []struct {
		a        string
		b        string
		expected int
	}{
		{
			a:        "kitten",
			b:        "sitting",
			expected: 3,
		},
	} {
		result := ed.LevenshteinDistance(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("expected %d, got %d", tc.expected, result)
		}
	}
}
