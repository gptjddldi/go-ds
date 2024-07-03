package sort

func bubbleSort(arr []int) []int {
	done := false
	for !done {
		done = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				done = false
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	return arr
}

// optimized
func bubbleSort2(arr []int) []int {
	done := false
	for i := 0; i < len(arr); i++ {
		done = true
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				done = false
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

		if done {
			break
		}
	}
	return arr
}
