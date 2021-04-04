package generation

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
			append(acc, g.createInitialStringFP(n, "")),
		)
	}

	lastItem := acc[len(acc)-1]
	if g.isEndStringFP(n, lastItem) {
		return acc
	}

	return g.GenerateFP(
		n,
		append(acc, g.generateNextStringFP(n, lastItem)),
	)
}

func (g *BinaryStringGenerator) createInitialStringFP(
	n int,
	acc string,
) string {
	if n <= 0 {
		return acc
	}

	return g.createInitialStringFP(n-1, acc+"0")
}

func (g *BinaryStringGenerator) isEndStringFP(
	n int,
	str string,
) bool {
	if n <= 0 {
		return false
	}

	if n == 1 {
		return []byte(str)[n-1] == '1'
	}

	return []byte(str)[n-1] == '1' && g.isEndStringFP(n-1, str[0:n-1])
}

func (g *BinaryStringGenerator) generateNextStringFP(
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
		g.createInitialStringFP(n-zeroIndex-1, "")
}
