package backtracking

type Point struct {
	X int
	Y int
}

func (bt *BackTracking) NQueens(n int) [][]Point {
	if n <= 0 {
		return [][]Point{}
	}

	var (
		try               func(acc []Point, idx int)
		ss                [][]Point    = make([][]Point, 0)
		columnCheck       map[int]bool = make(map[int]bool, n)
		crossCheck        map[int]bool = make(map[int]bool, 2*n)
		reverseCrossCheck map[int]bool = make(map[int]bool, 2*n)
	)

	try = func(acc []Point, row int) {
		if row == n {
			ss = append(ss, acc)
		} else {
			for col := 0; col < n; col++ {
				if !columnCheck[col] && !crossCheck[row+col] && !reverseCrossCheck[row-col] {
					columnCheck[col] = true
					crossCheck[row+col] = true
					reverseCrossCheck[row-col] = true

					try(
						append(acc, Point{row + 1, col + 1}),
						row+1,
					)

					columnCheck[col] = false
					crossCheck[row+col] = false
					reverseCrossCheck[row-col] = false
				}
			}
		}
	}

	try([]Point{}, 0)

	return ss
}
