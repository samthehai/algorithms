package backtracking_test

import (
	"testing"

	bt "github.com/samthehai/algorithms/cdvgt/generator/backtracking"
)

func (s *BackTrackingTestSuit) TestNQueens() {
	testcases := []struct {
		name     string
		n        int
		expected [][]bt.Point
	}{
		{
			name:     "n = 0",
			n:        0,
			expected: [][]bt.Point{},
		},
		{
			name: "n = 1",
			n:    1,
			expected: [][]bt.Point{
				{bt.Point{1, 1}},
			},
		},
		{
			name:     "n = 2",
			n:        2,
			expected: [][]bt.Point{},
		},
		{
			name:     "n = 3",
			n:        3,
			expected: [][]bt.Point{},
		},
		{
			name: "n = 5",
			n:    5,
			expected: [][]bt.Point{
				{{1, 1}, {2, 3}, {3, 5}, {4, 2}, {5, 4}},
				{{1, 1}, {2, 4}, {3, 2}, {4, 5}, {5, 3}},
				{{1, 2}, {2, 4}, {3, 1}, {4, 3}, {5, 5}},
				{{1, 2}, {2, 5}, {3, 3}, {4, 1}, {5, 4}},
				{{1, 3}, {2, 1}, {3, 4}, {4, 2}, {5, 5}},
				{{1, 3}, {2, 5}, {3, 2}, {4, 4}, {5, 1}},
				{{1, 4}, {2, 1}, {3, 3}, {4, 5}, {5, 2}},
				{{1, 4}, {2, 2}, {3, 5}, {4, 3}, {5, 1}},
				{{1, 5}, {2, 2}, {3, 4}, {4, 1}, {5, 3}},
				{{1, 5}, {2, 3}, {3, 1}, {4, 4}, {5, 2}},
			},
		},
	}

	for _, c := range testcases {
		s.T().Run(c.name, func(t *testing.T) {
			g := bt.BackTracking{}
			s.Assert().Equal(
				c.expected,
				g.NQueens(c.n),
			)
		})
	}
}

var benchmarkNQueensResult [][]bt.Point

func BenchmarkNQueens(b *testing.B) {
	var r [][]bt.Point
	g := bt.BackTracking{}
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = g.NQueens(30)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	benchmarkNQueensResult = r
}
