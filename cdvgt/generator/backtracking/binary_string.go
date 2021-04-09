package backtracking

func (bt *BackTracking) BinaryStringGenerate(n int) []string {
	if n <= 0 {
		return []string{}
	}

	var (
		try func(acc string, idx int)
		ss  []string = make([]string, 0)
	)

	try = func(acc string, idx int) {
		if idx == n {
			ss = append(ss, acc)
		} else {
			binaryDigits := []byte{'0', '1'}
			for _, d := range binaryDigits {
				try(acc+string(d), idx+1)
			}
		}
	}

	try("", 0)

	return ss
}
