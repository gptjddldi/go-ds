package sort

func countingSort(arr []int) []int {
	maxi := 0
	for _, v := range arr {
		if v > maxi {
			maxi = v
		}
	}

	count := make([]int, maxi+1)

	for _, v := range arr {
		count[v]++
	}

	i := 0
	for j := 0; j <= maxi; j++ {
		for count[j] > 0 {
			arr[i] = j
			i++
			count[j]--
		}
	}

	return arr
}
