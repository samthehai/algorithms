package generation_test

import (
	"testing"

	"github.com/samthehai/algorithms/cdvgt/generator/generation"
)

func (s *GenerationTestSuit) TestSubSetGenerator() {
	testcases := []struct {
		name     string
		N        int
		K        int
		expected [][]int
	}{
		{
			name:     "n,k = 0, 0",
			N:        0,
			K:        0,
			expected: [][]int{},
		},
		{
			name: "n,k = 1,1",
			N:    1,
			K:    1,
			expected: [][]int{
				{1},
			},
		},
		{
			name: "n,k = 2, 1",
			N:    2,
			K:    1,
			expected: [][]int{
				{1},
				{2},
			},
		},
		{
			name: "n,k = 5, 3",
			N:    5,
			K:    3,
			expected: [][]int{
				{1, 2, 3},
				{1, 2, 4},
				{1, 2, 5},
				{1, 3, 4},
				{1, 3, 5},
				{1, 4, 5},
				{2, 3, 4},
				{2, 3, 5},
				{2, 4, 5},
				{3, 4, 5},
			},
		},
	}

	for _, c := range testcases {
		s.T().Run(c.name, func(t *testing.T) {
			g := generation.Generation{}
			s.Assert().Equal(
				c.expected,
				g.SubSetGenerate(c.N, c.K),
			)
		})
	}
}

var benchmarkSubSetGeneratorResult [][]int

func BenchmarkSubSetGenerator(b *testing.B) {
	var r [][]int
	g := generation.Generation{}
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = g.SubSetGenerate(10, 5)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	benchmarkSubSetGeneratorResult = r
}
