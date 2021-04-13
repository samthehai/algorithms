package dynamicprograming

func (d DynamicProgramming) AnalyseNumber(n int) int {
	if n <= 0 {
		return 1
	}

	var f []int = make([]int, n+1)

	f[0] = 1

	for m := 1; m <= n; m++ {
		for v := m; v <= n; v++ {
			f[v] = f[v] + f[v-m]
		}
	}

	return f[n]
}

func (d DynamicProgramming) AnalyseNumberRecursive(n int) int {
	if n < 0 {
		return 0
	}

	var getF func(m, v int) int

	getF = func(m, v int) int {
		if m == 0 {
			switch {
			case v == 0:
				return 1
			default:
				return 0
			}
		}

		switch {
		case m > v:
			return getF(m-1, v)
		default:
			return getF(m-1, v) + getF(m, v-m)
		}
	}

	return getF(n, n)
}

func (d DynamicProgramming) AnalyseNumberRecursiveFaster(n int) int {
	if n < 0 {
		return 0
	}

	var (
		getF func(m, v int) int
		f    [][]int = make([][]int, n+1)
	)

	getF = func(m, v int) int {
		if f[m] == nil || f[m][v] == -1 {
			if f[m] == nil {
				f[m] = make([]int, n+1)
				for i := 0; i < n+1; i++ {
					f[m][i] = -1
				}
			}

			var value int

			if m == 0 {
				switch {
				case v == 0:
					value = 1
				default:
				}
			} else {
				switch {
				case m > v:
					value = getF(m-1, v)
				default:
					value = getF(m-1, v) + getF(m, v-m)
				}
			}

			f[m][v] = value
		}

		return f[m][v]
	}

	return getF(n, n)
}
