package algorithm

import "sortingvisualizer/internal/array"

func Selection(arr *array.Array) {
	for i := range arr.Len() {
		minIdx := i
		for j:=i+1; j<arr.Len(); j++ {
			if arr.IsLT(j, minIdx) {
				minIdx = j
			}
		}

		if i != minIdx {
			arr.Swap(i, minIdx)
		}
	}
}