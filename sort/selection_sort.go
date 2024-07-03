package sort

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i; j < len(arr); j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
	return arr
}
