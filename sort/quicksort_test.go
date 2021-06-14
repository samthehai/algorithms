package sort_test

import (
	"testing"

	mysort "github.com/samthehai/algorithms/sort"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "one item input",
			input:  []int{1},
			output: []int{1},
		},
		{
			name:   "two items sorted input",
			input:  []int{1, 2},
			output: []int{1, 2},
		},
		{
			name:   "three items sorted input",
			input:  []int{1, 2, 3},
			output: []int{1, 2, 3},
		},
		{
			name:   "three items input",
			input:  []int{2, 1, 3},
			output: []int{1, 2, 3},
		},
		{
			name:   "three items reverse input",
			input:  []int{3, 2, 1},
			output: []int{1, 2, 3},
		},
		{
			name:   "normal input",
			input:  []int{8, 3, 7, 9, 6, 1, 9, 10},
			output: []int{1, 3, 6, 7, 8, 9, 9, 10},
		},
		{
			name:   "different number digit input",
			input:  []int{8, 2, 78, 892, 11, 0, 34},
			output: []int{0, 2, 8, 11, 34, 78, 892},
		},
		{
			name:   "input with different number digit longer",
			input:  []int{9, 3, 83, 9, 2, 0, 1, 65, 2, 822, 9, 11, 22, 3, 3, 3, 47},
			output: []int{0, 1, 2, 2, 3, 3, 3, 3, 9, 9, 9, 11, 22, 47, 65, 83, 822},
		},
		{
			name:   "mixed negative positive and zero input",
			input:  []int{-6, 9, 0, 1, 17, 91, 0, 178},
			output: []int{-6, 0, 0, 1, 9, 17, 91, 178},
		},
		{
			name:   "reverse input",
			input:  []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			output: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:   "negative input",
			input:  []int{-3, -2, -1, -9, -5, -1, -19, -33},
			output: []int{-33, -19, -9, -5, -3, -2, -1, -1},
		},
		{
			name:   "duplicate zero input",
			input:  []int{-5, -6, -7, 0, 0, 0, 0, -8, 1, 2, 3},
			output: []int{-8, -7, -6, -5, 0, 0, 0, 0, 1, 2, 3},
		},
	}

	for _, c := range cases {
		arr := c.input
		mysort.QuickSort(arr, 0, len(arr)-1)
		assert.Equal(t, c.output, arr)
	}
}
