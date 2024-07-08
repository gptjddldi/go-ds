package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sortTestCases = []struct {
	name string
	arr  []int
	want []int
}{
	{
		name: "Empty array",
		arr:  []int{},
		want: []int{},
	},
	{
		name: "Single element array",
		arr:  []int{1},
		want: []int{1},
	},
	{
		name: "Reverse order",
		arr:  []int{5, 4, 3, 2, 1},
		want: []int{1, 2, 3, 4, 5},
	},
	{
		name: "Already sorted",
		arr:  []int{1, 2, 3, 4, 5},
		want: []int{1, 2, 3, 4, 5},
	},
	{
		name: "Mixed order",
		arr:  []int{1, 3, 2, 5, 4},
		want: []int{1, 2, 3, 4, 5},
	},
	{
		name: "All elements same",
		arr:  []int{1, 1, 1, 1, 1},
		want: []int{1, 1, 1, 1, 1},
	},
	{
		name: "Mostly same elements with one different",
		arr:  []int{3, 3, 3, 3, 3, 3, 1},
		want: []int{1, 3, 3, 3, 3, 3, 3},
	},
	{
		name: "Larger array",
		arr:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		name: "Array with negative numbers",
		arr:  []int{-3, 4, 1, -1, 0, 5, -2},
		want: []int{-3, -2, -1, 0, 1, 4, 5},
	},
	{
		name: "Array with very large and small numbers",
		arr:  []int{1000000, -1000000, 0, 1, -1},
		want: []int{-1000000, -1, 0, 1, 1000000},
	},
	{
		name: "Array with many duplicates",
		arr:  []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
		want: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
	},
	{
		name: "Reverse of already sorted array",
		arr:  []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
}

func Test_insertionSort(t *testing.T) {
	for _, tt := range sortTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, insertionSort(tt.arr), tt.name)
		})
	}
}

func Test_selectionSort(t *testing.T) {
	for _, tt := range sortTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, selectionSort(tt.arr), tt.name)
		})
	}
}

func Test_bubbleSort(t *testing.T) {
	for _, tt := range sortTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bubbleSort(tt.arr), tt.name)
		})
	}
}

func Test_bubbleSort2(t *testing.T) {
	for _, tt := range sortTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bubbleSort2(tt.arr), tt.name)
		})
	}
}

func Test_mergeSort(t *testing.T) {
	for _, tt := range sortTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, mergeSort(tt.arr), tt.name)
		})
	}
}
