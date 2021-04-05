package generation_test

import (
	"testing"

	"github.com/samthehai/algorithms/cdvgt/generator/generation"
	"github.com/stretchr/testify/suite"
)

type testBinaryStringGenerator struct {
	suite.Suite
}

func TestBinaryStringGenerator(t *testing.T) {
	suite.Run(t, &testBinaryStringGenerator{})
}

func (s *testBinaryStringGenerator) TestGenerateBinaryString() {
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
			g := generation.BinaryStringGenerator{
				N: c.n,
			}
			s.Assert().Equal(
				c.expected,
				g.Generate(),
			)
		})
	}
}

var result []string

func BenchmarkGenerateBinaryString(b *testing.B) {
	var r []string
	g := generation.BinaryStringGenerator{
		N: 10,
	}
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = g.Generate()
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func (s *testBinaryStringGenerator) TestGenerateBinaryStringFP() {
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
			g := generation.BinaryStringGenerator{}
			s.Assert().Equal(
				c.expected,
				g.GenerateFP(c.n, []string{}),
			)
		})
	}
}

var resultFP []string

func BenchmarkGenerateBinaryStringFP(b *testing.B) {
	var r []string
	g := generation.BinaryStringGenerator{}
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = g.GenerateFP(10, []string{})
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	resultFP = r
}
