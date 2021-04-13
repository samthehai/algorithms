package dynamicprograming_test

import (
	"testing"

	"github.com/samthehai/algorithms/cdvgt/dynamicprograming"
)

func (s *DynamicProgrammingTestSuit) TestAnalyseNumber() {
	testcases := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "n = 0",
			n:        0,
			expected: 1,
		},
		{
			name:     "n = 1",
			n:        1,
			expected: 1,
		},
		{
			name:     "n = 2",
			n:        2,
			expected: 2,
		},
		{
			name:     "n = 5",
			n:        5,
			expected: 7,
		},
		{
			name:     "n = 6",
			n:        6,
			expected: 11,
		},
	}

	for _, c := range testcases {
		s.T().Run(c.name, func(t *testing.T) {
			d := dynamicprograming.DynamicProgramming{}
			s.Assert().Equal(
				c.expected,
				d.AnalyseNumber(c.n),
			)
		})
	}
}

func (s *DynamicProgrammingTestSuit) TestAnalyseNumberRecursive() {
	testcases := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "n = 0",
			n:        0,
			expected: 1,
		},
		{
			name:     "n = 1",
			n:        1,
			expected: 1,
		},
		{
			name:     "n = 2",
			n:        2,
			expected: 2,
		},
		{
			name:     "n = 5",
			n:        5,
			expected: 7,
		},
		{
			name:     "n = 6",
			n:        6,
			expected: 11,
		},
	}

	for _, c := range testcases {
		s.T().Run(c.name, func(t *testing.T) {
			d := dynamicprograming.DynamicProgramming{}
			s.Assert().Equal(
				c.expected,
				d.AnalyseNumberRecursive(c.n),
			)
		})
	}
}

func (s *DynamicProgrammingTestSuit) TestAnalyseNumberRecursiveFaster() {
	testcases := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "n = 0",
			n:        0,
			expected: 1,
		},
		{
			name:     "n = 1",
			n:        1,
			expected: 1,
		},
		{
			name:     "n = 2",
			n:        2,
			expected: 2,
		},
		{
			name:     "n = 5",
			n:        5,
			expected: 7,
		},
		{
			name:     "n = 6",
			n:        6,
			expected: 11,
		},
	}

	for _, c := range testcases {
		s.T().Run(c.name, func(t *testing.T) {
			d := dynamicprograming.DynamicProgramming{}
			s.Assert().Equal(
				c.expected,
				d.AnalyseNumberRecursiveFaster(c.n),
			)
		})
	}
}
