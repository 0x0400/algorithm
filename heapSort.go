package algorithm

type Heap struct {
	data     []int
	heapSize int
}

func maxHeapify(heap *Heap, index int) {
	left := index * 2
	right := left + 1
	largest := index
	if left <= heap.heapSize && heap.data[left] > heap.data[index] {
		largest = left
	}
	if right <= heap.heapSize && heap.data[right] > heap.data[largest] {
		largest = right
	}
	if largest != index {
		heap.data[index], heap.data[largest] = heap.data[largest], heap.data[index]
		maxHeapify(heap, largest)
	}
}

func buildMaxHeap(heap *Heap) {
	for i := heap.heapSize / 2; i >= 0; i-- {
		maxHeapify(heap, i)
	}
}

func HeapSort(data []int) {
	heap := &Heap{data, len(data) - 1}
	buildMaxHeap(heap)
	for i := heap.heapSize; i > 0; i-- {
		heap.data[0], heap.data[i] = heap.data[i], heap.data[0]
		heap.heapSize -= 1
		maxHeapify(heap, 0)
	}
}
