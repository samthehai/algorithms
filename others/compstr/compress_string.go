package compstr

import "fmt"

type CompressString struct{}

func NewCompressString() CompressString {
	return CompressString{}
}

func (comp *CompressString) Compress(str string) string {
	n := len(str)
	if n < 1 {
		return str
	}

	compStr, prev, count := "", rune(str[0]), 0

	for _, c := range str {
		if c == prev {
			count++
		} else {
			compStr += comp.calcPartialCompStr(prev, count)
			prev = c
			count = 1
		}
	}

	compStr += comp.calcPartialCompStr(prev, count)

	fmt.Println(compStr)

	if len(compStr) < len(str) {
		return compStr
	}

	return str
}

func (comp *CompressString) calcPartialCompStr(prev rune, count int) string {
	if count > 1 {
		return fmt.Sprint(string(prev), count)
	}

	return fmt.Sprint(string(prev))
}
