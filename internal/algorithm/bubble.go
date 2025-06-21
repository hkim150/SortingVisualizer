package algorithm

import "sortingvisualizer/internal/array"

func Bubble(arr *array.Array) error {
	for i:=arr.Len()-1; i>0; i-- {
		for j:=0; j<i; j++ {
			if arr.IsGT(j, j+1) {
				arr.Swap(j, j+1)
			}
		}
	}
	
	return nil
}
