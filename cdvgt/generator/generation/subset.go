package generation

type SubSetGenerator struct {
	N int
	K int
}

func (g *SubSetGenerator) Generate() [][]int {
	ss := make([][]int, 0)

	if g.K > g.N || g.K <= 0 || g.N <= 0 {
		return ss
	}

	current := g.first()
	ss = append(ss, current)

	for g.hasNext(current) {
		next := g.generateNext(current)
		ss = append(ss, next)

		current = next
	}

	return ss
}

func (g *SubSetGenerator) hasNext(current []int) bool {
	for i := g.K - 1; i >= 0; i-- {
		if current[i] != g.N+i-g.K+1 {
			return true
		}
	}

	return false
}

func (g *SubSetGenerator) generateNext(current []int) []int {
	next := make([]int, 0, g.K)
	next = append(next, current...)

	idx := g.K - 1
	for ; next[idx] == g.N+idx-g.K+1 && idx >= 0; idx-- {
	}

	next[idx] += 1
	for i := idx + 1; i < g.K; i++ {
		next[i] = next[i-1] + 1
	}

	return next
}

func (g *SubSetGenerator) first() []int {
	s := make([]int, 0, g.K)

	for i := 0; i < g.K; i++ {
		s = append(s, i+1)
	}

	return s
}
