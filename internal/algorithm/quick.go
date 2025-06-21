package algorithm

import (
	"math/rand"
	"sortingvisualizer/internal/array"
)

func Quick(arr *array.Array) error {
	quickSort(arr, 0, arr.Len()-1)
	return nil
}

func quickSort(arr *array.Array, start, end int) {
	if start >= end {
		return
	}

	pivot := start + rand.Intn(end-start+1)

	// partition
	arr.Swap(pivot, end)
	pivot = end

	l := start
	r := end - 1

	for l <= r {
		if arr.IsLTE(l, pivot) {
			l++
		} else {
			arr.Swap(l, r)
			r--
		}
	}

	arr.Swap(l, pivot)
	pivot = l

	quickSort(arr, start, pivot-1)
	quickSort(arr, pivot+1, end)
}
