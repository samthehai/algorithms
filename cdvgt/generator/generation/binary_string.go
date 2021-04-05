package generation

func (g *Generation) BinaryStringGenerate(n int) []string {
	ss := make([]string, 0)

	if n <= 0 {
		return ss
	}

	current := g.binaryStringGenerateFirst(n)
	for !g.binaryStringGenerateHasNext(n, current) {
		ss = append(ss, current)
		current = g.binaryStringGenerateNext(n, current)
	}

	ss = append(ss, current)

	return ss
}

func (g *Generation) binaryStringGenerateFirst(
	digitNum int,
) string {
	str := ""
	for i := 0; i < digitNum; i++ {
		str += "0"
	}

	return str
}

func (g *Generation) binaryStringGenerateHasNext(
	n int,
	current string,
) bool {
	for i := 0; i < n; i++ {
		if current[i] == '0' {
			return false
		}
	}

	return true
}

func (g *Generation) binaryStringGenerateNext(
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
		g.binaryStringGenerateFirst(n-zeroIndex-1)
}

func (g *Generation) BinaryStringGenerateFP(
	n int,
	acc []string,
) []string {
	if n <= 0 {
		return acc
	}

	if len(acc) == 0 {
		return g.BinaryStringGenerateFP(
			n,
			append(acc, g.binaryStringGenerateFirstFP(n, "")),
		)
	}

	lastItem := acc[len(acc)-1]
	if g.binaryStringGenerateHasNextFP(n, lastItem) {
		return acc
	}

	return g.BinaryStringGenerateFP(
		n,
		append(acc, g.binaryStringGenerateNextFP(n, lastItem)),
	)
}

func (g *Generation) binaryStringGenerateFirstFP(
	n int,
	acc string,
) string {
	if n <= 0 {
		return acc
	}

	return g.binaryStringGenerateFirstFP(n-1, acc+"0")
}

func (g *Generation) binaryStringGenerateHasNextFP(
	n int,
	str string,
) bool {
	if n <= 0 {
		return false
	}

	if n == 1 {
		return []byte(str)[n-1] == '1'
	}

	return []byte(str)[n-1] == '1' &&
		g.binaryStringGenerateHasNextFP(n-1, str[0:n-1])
}

func (g *Generation) binaryStringGenerateNextFP(
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
		g.binaryStringGenerateFirstFP(n-zeroIndex-1, "")
}
