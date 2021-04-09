package backtracking_test

import (
	"testing"

	"github.com/samthehai/algorithms/cdvgt/generator/backtracking"
)

func (s *BackTrackingTestSuit) TestAnalyseNumber() {
	testcases := []struct {
		name     string
		n        int
		expected [][]int
	}{
		{
			name:     "n = 0",
			n:        0,
			expected: [][]int{},
		},
		{
			name: "n = 1",
			n:    1,
			expected: [][]int{
				{1},
			},
		},
		{
			name: "n = 2",
			n:    2,
			expected: [][]int{
				{1, 1},
				{2},
			},
		},
		{
			name: "n = 5",
			n:    5,
			expected: [][]int{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 2},
				{1, 1, 3},
				{1, 2, 2},
				{1, 4},
				{2, 3},
				{5},
			},
		},
		{
			name: "n = 6",
			n:    6,
			expected: [][]int{
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 2},
				{1, 1, 1, 3},
				{1, 1, 2, 2},
				{1, 1, 4},
				{1, 2, 3},
				{1, 5},
				{2, 2, 2},
				{2, 4},
				{3, 3},
				{6},
			},
		},
	}

	for _, c := range testcases {
		s.T().Run(c.name, func(t *testing.T) {
			g := backtracking.BackTracking{}
			s.Assert().Equal(
				c.expected,
				g.AnalyseNumber(c.n),
			)
		})
	}
}

var benchmarkAnalyseNumberResult [][]int

func BenchmarkAnalyseNumber(b *testing.B) {
	var r [][]int
	g := backtracking.BackTracking{}
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = g.AnalyseNumber(30)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	benchmarkAnalyseNumberResult = r
}
