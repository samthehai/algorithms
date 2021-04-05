package generation

func (g *Generation) SubSetGenerate(n, k int) [][]int {
	ss := make([][]int, 0)

	if k > n || k <= 0 || n <= 0 {
		return ss
	}

	current := g.subSetGenerateFirst(n, k)
	ss = append(ss, current)

	for g.subSetGenerateHasNext(n, k, current) {
		next := g.subSetGenerateNext(n, k, current)
		ss = append(ss, next)

		current = next
	}

	return ss
}

func (g *Generation) subSetGenerateHasNext(n, k int, current []int) bool {
	for i := k - 1; i >= 0; i-- {
		if current[i] != n+i-k+1 {
			return true
		}
	}

	return false
}

func (g *Generation) subSetGenerateNext(n, k int, current []int) []int {
	next := make([]int, 0, k)
	next = append(next, current...)

	idx := k - 1
	for ; next[idx] == n+idx-k+1 && idx >= 0; idx-- {
	}

	next[idx] += 1
	for i := idx + 1; i < k; i++ {
		next[i] = next[i-1] + 1
	}

	return next
}

func (g *Generation) subSetGenerateFirst(n, k int) []int {
	s := make([]int, 0, k)

	for i := 0; i < k; i++ {
		s = append(s, i+1)
	}

	return s
}
