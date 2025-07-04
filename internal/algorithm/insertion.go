package algorithm

import "sortingvisualizer/internal/array"

func Insertion(arr *array.Array)  {
	for r:=1; r<arr.Len(); r++ {
		for l := range r {
			if arr.IsGT(l, r) {
				arr.Swap(l, r)
			}
		}
	}
}