package generation_test

import (
	"testing"

	"github.com/samthehai/algorithms/cdvgt/generator/generation"
)

func (s *GenerationTestSuit) TestGenerateBinaryString() {
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
			g := generation.Generation{}
			s.Assert().Equal(
				c.expected,
				g.BinaryStringGenerate(c.n),
			)
		})
	}
}

var benchmarkGenerateBinaryStringResult []string

func BenchmarkGenerateBinaryString(b *testing.B) {
	var r []string
	g := generation.Generation{}
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = g.BinaryStringGenerate(10)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	benchmarkGenerateBinaryStringResult = r
}

func (s *GenerationTestSuit) TestGenerateBinaryStringFP() {
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
			g := generation.Generation{}
			s.Assert().Equal(
				c.expected,
				g.BinaryStringGenerateFP(c.n, []string{}),
			)
		})
	}
}

var benchmarkGenerateBinaryStringFPResult []string

func BenchmarkGenerateBinaryStringFP(b *testing.B) {
	var r []string
	g := generation.Generation{}
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = g.BinaryStringGenerateFP(10, []string{})
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	benchmarkGenerateBinaryStringFPResult = r
}
