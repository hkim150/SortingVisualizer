package algorithm

import "sortingvisualizer/internal/array"

func Shell(arr *array.Array) {
	gap := arr.Len() / 2
	for gap > 0 {
		for i := gap; i < arr.Len(); i++ {
			j := i
			for j-gap >= 0 && arr.IsGT(j-gap, j) {
				arr.Swap(j-gap, j)
				j -= gap
			}
		}

		gap /= 2
	}
}
