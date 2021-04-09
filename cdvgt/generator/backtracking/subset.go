package backtracking

func (bt *BackTracking) SubSetGenerate(n, k int) [][]int {
	if n <= 0 {
		return [][]int{}
	}

	var (
		try func(acc []int, idx int)
		ss  [][]int = make([][]int, 0)
	)

	try = func(acc []int, idx int) {
		if idx == k {
			ss = append(ss, acc)
		} else {
			start := 0
			if len(acc) > 0 {
				start = acc[len(acc)-1]
			}

			for i := start + 1; i <= n-k+idx+1; i++ {
				try(append(acc, i), idx+1)
			}
		}
	}

	try([]int{}, 0)

	return ss
}
