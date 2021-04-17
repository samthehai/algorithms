package dynamicprograming

func (d DynamicProgramming) LongestIncreaseSubSequence(sequence []int) []int {
	n := len(sequence)
	// the new input array that have -inf and +inf at 2 boundaries of array
	ss := make([]int, 0, n+2)

	// tt[i] contains the next index in the longest subsequence
	// after the index i in the sub sequence
	tt := make([]int, n+2)

	optimize := func() {
		const (
			uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
			maxInt   = 1<<(uintSize-1) - 1        // 1<<31 - 1 or 1<<63 - 1
			minInt   = -maxInt - 1                // -1 << 31 or -1 << 63
		)
		// ll[i] contains the length of longest increase sub sequence
		// start at index i of ss
		ll := make([]int, n+2)

		ss = append(ss, minInt)
		ss = append(ss, sequence...)
		ss = append(ss, maxInt)
		ll[n+1] = 1

		for i := n; i >= 0; i-- {
			jmax := n + 1
			for j := i + 1; j < n+2; j++ {
				if ss[j] > ss[i] && ll[j] > ll[jmax] {
					jmax = j
				}
			}
			ll[i] = ll[jmax] + 1
			tt[i] = jmax
		}
	}

	result := func() []int {
		res := make([]int, 0)
		i := tt[0]
		for ; i < n+1; i = tt[i] {
			res = append(res, ss[i])
		}

		return res
	}

	optimize()
	return result()
}
