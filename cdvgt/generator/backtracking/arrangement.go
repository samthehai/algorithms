package backtracking

func (bt *BackTracking) ArrangementGenerate(n, k int) [][]int {
	if n <= 0 || k <= 0 {
		return [][]int{}
	}

	var (
		try func(acc []int, included []bool, idx int)
		ss  [][]int = make([][]int, 0)
	)

	try = func(acc []int, included []bool, idx int) {
		if idx == k {
			ss = append(ss, acc)
		} else {
			for i := 0; i < n; i++ {
				if included[i] {
					continue
				}

				try(
					append(acc, i+1),
					append(append(append([]bool{}, included[:i]...), true), included[i+1:]...),
					idx+1,
				)
			}
		}
	}

	try([]int{}, make([]bool, n), 0)

	return ss
}
