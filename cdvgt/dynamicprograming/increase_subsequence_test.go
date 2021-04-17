package dynamicprograming_test

import (
	"testing"

	"github.com/samthehai/algorithms/cdvgt/dynamicprograming"
)

func (s *DynamicProgrammingTestSuit) TestLongestIncreaseSubSequence() {
	testcases := []struct {
		name     string
		sequence []int
		expected []int
	}{
		{
			name:     "test 1",
			sequence: []int{},
			expected: []int{},
		},
		{
			name:     "test 2",
			sequence: []int{5, 2, 3, 4, 9, 10, 5, 6, 7, 8},
			expected: []int{2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:     "test 3",
			sequence: []int{1, 2, 3, 8, 9, 4, 5, 6, 20, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 9, 10},
		},
	}

	for _, c := range testcases {
		s.T().Run(c.name, func(t *testing.T) {
			d := dynamicprograming.DynamicProgramming{}
			s.Assert().Equal(
				c.expected,
				d.LongestIncreaseSubSequence(c.sequence),
			)
		})
	}
}
