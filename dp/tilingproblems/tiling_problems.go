package tilingproblems

// How many ways can you tile a board length of n with red tiles length of 1
// and blue tiles length of 2
type FindNumberOfWayTileBoard struct {
}

// Recursive is function using recursive approach
func (p FindNumberOfWayTileBoard) Recursive(n int) (numberOfWay int) {
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	return p.Recursive(n-1) + p.Recursive(n-2)
}

// TopDown is function using topdown approach
func (p FindNumberOfWayTileBoard) TopDown(n int) (numberOfWay int) {
	var (
		solver      func(n int) (result int)
		dp          []int = make([]int, n+1)
		initializer func()
	)

	initializer = func() {
		for i := range dp {
			dp[i] = -1
		}
	}

	solver = func(n int) (result int) {
		if n < 0 {
			return 0
		}

		if n == 0 {
			return 1
		}

		if dp[n] != -1 {
			return dp[n]
		}

		dp[n] = dp[n-1] + dp[n-2]
		return dp[n]
	}

	initializer()

	return solver(n)
}

// BottomUp is function using bottom-up approach
func (p FindNumberOfWayTileBoard) BottomUp(n int) (numberOfWay int) {
	solver := func(n int) (result int) {
		if n == 0 || n == 1 {
			return 1
		}

		dp := make([]int, n+1)
		dp[1] = 1

		for i := 2; i <= n; i++ {
			dp[i] = dp[i-1] + dp[i-2]
		}

		return dp[n]
	}

	return solver(n)
}
