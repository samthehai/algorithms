package generation

func (g *BinaryStringGenerator) Generate(
	n int,
) []string {
	ss := make([]string, 0)

	if n <= 0 {
		return ss
	}

	current := g.createInitialString(n)
	for !g.isEndString(n, current) {
		ss = append(ss, current)
		current = g.generateNextString(n, current)
	}

	ss = append(ss, current)

	return ss
}

func (g *BinaryStringGenerator) createInitialString(
	n int,
) string {
	str := ""
	for i := 0; i < n; i++ {
		str += "0"
	}

	return str
}

func (g *BinaryStringGenerator) isEndString(
	n int,
	str string,
) bool {
	for i := 0; i < n; i++ {
		if str[i] == '0' {
			return false
		}
	}

	return true
}

func (g *BinaryStringGenerator) generateNextString(
	n int,
	currentString string,
) string {
	zeroIndex := n - 1
	for ; zeroIndex >= 0; zeroIndex-- {
		if []byte(currentString)[zeroIndex] == '0' {
			break
		}
	}

	return currentString[0:zeroIndex] +
		"1" +
		g.createInitialString(n-zeroIndex-1)
}
