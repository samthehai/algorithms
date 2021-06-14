package sort

import "fmt"

func QuickSort(arr []int, low, high int) {
	if low > high {
		panic(fmt.Errorf("can not process with low greater than high"))
	}

	index := partition(arr, low, high)

	if low < index-1 {
		QuickSort(arr, low, index-1)
	}

	if high > index {
		QuickSort(arr, index, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[low+(high-low)/2]

	left, right := low, high
	for left <= right {
		for arr[left] < pivot {
			left++
		}

		for arr[right] > pivot {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	return left
}
