package backtracking_test

import (
	"testing"

	"github.com/samthehai/algorithms/cdvgt/generator/backtracking"
)

func (s *BackTrackingTestSuit) TestGenerateBinaryString() {
	testcases := []struct {
		name     string
		n        int
		expected []string
	}{
		{
			name:     "n equal 0",
			n:        0,
			expected: []string{},
		},
		{
			name:     "n qual 1",
			n:        1,
			expected: []string{"0", "1"},
		},
		{
			name:     "n qual 2",
			n:        2,
			expected: []string{"00", "01", "10", "11"},
		},
		{
			name:     "n qual 3",
			n:        3,
			expected: []string{"000", "001", "010", "011", "100", "101", "110", "111"},
		},
	}

	for _, c := range testcases {
		s.T().Run(c.name, func(t *testing.T) {
			bt := backtracking.BackTracking{}
			s.Assert().Equal(
				c.expected,
				bt.BinaryStringGenerate(c.n),
			)
		})
	}
}

var benchmarkGenerateBinaryStringResult []string

func BenchmarkGenerateBinaryString(b *testing.B) {
	var r []string
	g := backtracking.BackTracking{}
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = g.BinaryStringGenerate(10)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	benchmarkGenerateBinaryStringResult = r
}
