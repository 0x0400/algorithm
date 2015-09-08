package algorithm

func partition(array []int, start int, end int) int {
    if start >= end {
        return start
    }
    pivot := array[end]
    pivotIdx := start - 1
    for i := start; i < end; i++ {
        if array[i] <= pivot {
            pivotIdx += 1
            array[pivotIdx], array[i] = array[i], array[pivotIdx]
        }
    }
    pivotIdx += 1
    array[pivotIdx], array[end] = array[end], array[pivotIdx]
    return pivotIdx
}

func QuickSort(array []int, start int, end int) {
    partitionIdx := partition(array, start, end)
    if partitionIdx > start {
        QuickSort(array, start, partitionIdx - 1)
    }
    if partitionIdx < end {
        QuickSort(array, partitionIdx + 1, end)
    }
}
