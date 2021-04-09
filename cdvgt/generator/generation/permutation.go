package generation

func (g *Generation) PermutationGenerate(n int) [][]int {
	pp := make([][]int, 0)

	if n <= 0 {
		return pp
	}

	if n == 1 {
		return [][]int{{1}}
	}

	for c := g.permutationGenerateFirst(n); c != nil; c = g.permutationGenerateNext(n, c) {
		pp = append(pp, c)
	}

	return pp
}

func (g *Generation) permutationGenerateFirst(n int) []int {
	a := make([]int, 0, n)

	for i := 0; i < n; i++ {
		a = append(a, i+1)
	}

	return a
}

func (g *Generation) permutationGenerateNext(n int, current []int) []int {
	swapIdx := n - 1
	for ; swapIdx > 0 && current[swapIdx-1] > current[swapIdx]; swapIdx-- {
	}

	if swapIdx == 0 {
		return nil
	} else {
		swapIdx--

		next := make([]int, 0, n)
		next = append(next, current...)

		idx := n - 1
		for ; idx > swapIdx && next[idx] < next[swapIdx]; idx-- {
		}

		// swap
		next[idx], next[swapIdx] = next[swapIdx], next[idx]

		// reverse
		for left, right := swapIdx+1, n-1; left < right; left, right = left+1, right-1 {
			next[left], next[right] = next[right], next[left]
		}

		return next
	}
}
