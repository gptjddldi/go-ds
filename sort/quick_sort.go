package sort

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	l := len(arr)
	pivot := arr[l-1]
	idx := 0
	for i := 0; i < l; i++ {
		if arr[i] < pivot {
			arr[idx], arr[i] = arr[i], arr[idx]
			idx++
		}
	}

	arr[idx], arr[l-1] = arr[l-1], arr[idx]

	quickSort(arr[:idx])
	quickSort(arr[idx+1:])
	return arr
}
