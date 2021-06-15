package sort

func MergeSort(arr []int) {
	helper := make([]int, len(arr))
	mergeSort(arr, helper, 0, len(arr)-1)
}

func mergeSort(arr []int, helper []int, low, high int) {
	if low < high {
		middle := low + (high-low)/2
		mergeSort(arr, helper, low, middle)
		mergeSort(arr, helper, middle+1, high)
		merge(arr, helper, low, middle, high)
	}
}

func merge(arr []int, helper []int, low, middle, high int) {
	for i := 0; i <= high; i++ {
		helper[i] = arr[i]
	}

	left, right, current := low, middle+1, low

	for left <= middle && right <= high {
		if helper[left] <= helper[right] {
			arr[current] = helper[left]
			left++
		} else {
			arr[current] = helper[right]
			right++
		}

		current++
	}

	remaining := middle - left
	for i := 0; i <= remaining; i++ {
		arr[current+i] = helper[left+i]
	}
}
