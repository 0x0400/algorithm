package algorithm

import (
	"sort"
)

func partition(array sort.Interface, start int, end int) int {
	if start >= end {
		return start
	}
	pivotIdx := start
	for i := start; i < end; i++ {
		if array.Less(i, end) {
			array.Swap(pivotIdx, i)
			pivotIdx++
		}
	}
	array.Swap(pivotIdx, end)
	return pivotIdx
}

func QuickSort(array sort.Interface, start int, end int) {
	partitionIdx := partition(array, start, end)
	if partitionIdx > start {
		QuickSort(array, start, partitionIdx-1)
	}
	if partitionIdx < end {
		QuickSort(array, partitionIdx+1, end)
	}
}
