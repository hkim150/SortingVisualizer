package algorithm

import "sortingvisualizer/internal/array"

// Use max heapify then extract the max element to the end of the array
func Heap(arr *array.Array) {
	Heapify(arr)

	for i := arr.Len()-1; i > 0; i-- {
		arr.Swap(0, i)
		siftDown(arr, i, 0)
	}
}

func Heapify(arr *array.Array) {
	for i := arr.Len()/2 - 1; i > -1; i-- {
		siftDown(arr, arr.Len(), i)
	}
}

func leftChildIdx(i int) int {
	return 2*i + 1
}

func rightChildIdx(i int) int {
	return 2*i + 2
}

func siftDown(arr *array.Array, n, i int) {
	for {
		l := leftChildIdx(i)
		r := rightChildIdx(i)
		largest := i

		if l < n && arr.IsGT(l, largest) {
			largest = l
		}

		if r < n && arr.IsGT(r, largest) {
			largest = r
		}

		if largest == i {
			break
		}

		arr.Swap(i, largest)
		i = largest
	}
}
