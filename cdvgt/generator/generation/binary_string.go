package generation

type BinaryStringGenerator struct {
	N int
}

func (g *BinaryStringGenerator) Generate() []string {
	ss := make([]string, 0)

	if g.N <= 0 {
		return ss
	}

	current := g.first(g.N)
	for !g.hasNext(current) {
		ss = append(ss, current)
		current = g.next(current)
	}

	ss = append(ss, current)

	return ss
}

func (g *BinaryStringGenerator) first(
	digitNum int,
) string {
	str := ""
	for i := 0; i < digitNum; i++ {
		str += "0"
	}

	return str
}

func (g *BinaryStringGenerator) hasNext(
	current string,
) bool {
	for i := 0; i < g.N; i++ {
		if current[i] == '0' {
			return false
		}
	}

	return true
}

func (g *BinaryStringGenerator) next(
	current string,
) string {
	zeroIndex := g.N - 1
	for ; zeroIndex >= 0; zeroIndex-- {
		if []byte(current)[zeroIndex] == '0' {
			break
		}
	}

	return current[0:zeroIndex] +
		"1" +
		g.first(g.N-zeroIndex-1)
}

func (g *BinaryStringGenerator) GenerateFP(
	n int,
	acc []string,
) []string {
	if n <= 0 {
		return acc
	}

	if len(acc) == 0 {
		return g.GenerateFP(
			n,
			append(acc, g.firstFP(n, "")),
		)
	}

	lastItem := acc[len(acc)-1]
	if g.hasNextFP(n, lastItem) {
		return acc
	}

	return g.GenerateFP(
		n,
		append(acc, g.nextFP(n, lastItem)),
	)
}

func (g *BinaryStringGenerator) firstFP(
	n int,
	acc string,
) string {
	if n <= 0 {
		return acc
	}

	return g.firstFP(n-1, acc+"0")
}

func (g *BinaryStringGenerator) hasNextFP(
	n int,
	str string,
) bool {
	if n <= 0 {
		return false
	}

	if n == 1 {
		return []byte(str)[n-1] == '1'
	}

	return []byte(str)[n-1] == '1' && g.hasNextFP(n-1, str[0:n-1])
}

func (g *BinaryStringGenerator) nextFP(
	n int,
	current string,
) string {
	zeroIndex := n - 1
	for ; zeroIndex >= 0; zeroIndex-- {
		if []byte(current)[zeroIndex] == '0' {
			break
		}
	}

	return current[0:zeroIndex] +
		"1" +
		g.firstFP(n-zeroIndex-1, "")
}
