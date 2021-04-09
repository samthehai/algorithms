package generation_test

import (
	"testing"

	"github.com/samthehai/algorithms/cdvgt/generator/generation"
)

func (s *GenerationTestSuit) TestPermutationGenerator() {
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
			name:     "n = 1",
			n:        1,
			expected: [][]int{{1}},
		},
		{
			name: "n = 2",
			n:    2,
			expected: [][]int{
				{1, 2},
				{2, 1},
			},
		},
		{
			name: "n = 4",
			n:    4,
			expected: [][]int{
				{1, 2, 3, 4},
				{1, 2, 4, 3},
				{1, 3, 2, 4},
				{1, 3, 4, 2},
				{1, 4, 2, 3},
				{1, 4, 3, 2},
				{2, 1, 3, 4},
				{2, 1, 4, 3},
				{2, 3, 1, 4},
				{2, 3, 4, 1},
				{2, 4, 1, 3},
				{2, 4, 3, 1},
				{3, 1, 2, 4},
				{3, 1, 4, 2},
				{3, 2, 1, 4},
				{3, 2, 4, 1},
				{3, 4, 1, 2},
				{3, 4, 2, 1},
				{4, 1, 2, 3},
				{4, 1, 3, 2},
				{4, 2, 1, 3},
				{4, 2, 3, 1},
				{4, 3, 1, 2},
				{4, 3, 2, 1},
			},
		},
	}

	for _, c := range testcases {
		s.T().Run(c.name, func(t *testing.T) {
			g := generation.Generation{}
			s.Assert().Equal(
				c.expected,
				g.PermutationGenerate(c.n),
			)
		})
	}
}

var benchmarkPermutationGeneratorResult [][]int

func BenchmarkPermutationGenerator(b *testing.B) {
	var r [][]int
	g := generation.Generation{}
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = g.PermutationGenerate(10)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	benchmarkPermutationGeneratorResult = r
}
