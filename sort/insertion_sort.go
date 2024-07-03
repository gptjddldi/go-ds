package sort

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		val := arr[i]
		j := i
		for j > 0 && arr[j-1] > val {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = val
	}

	return arr
}
