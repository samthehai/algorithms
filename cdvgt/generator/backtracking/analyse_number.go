package backtracking

func (bt *BackTracking) AnalyseNumber(n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}

	var (
		try    func(acc []int, sum int)
		result [][]int = make([][]int, 0)
	)

	try = func(acc []int, sum int) {
		start := 1
		if len(acc) > 0 {
			start = acc[len(acc)-1]
		}

		for i := start; i <= (n-sum)/2; i++ {
			try(append(acc, i), sum+i)
		}

		result = append(result, append(acc, n-sum))
	}

	try([]int{}, 0)

	return result
}
