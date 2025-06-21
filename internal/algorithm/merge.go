package algorithm

import (
	"sortingvisualizer/internal/array"
)

func Merge(arr *array.Array) error {
	mergeSort(arr, 0, arr.Len()-1)
	return nil
}

func mergeSort(arr *array.Array, start, end int) {
	if start >= end {
		return
	}

	half := (end - start + 1) / 2
	s1, e1 := start, start+half-1
	s2, e2 := start+half, end
	mergeSort(arr, s1, e1)
	mergeSort(arr, s2, e2)

	temp := make([]int, 0, end-start+1)
	for s1 <= e1 || s2 <= e2 {
		if s1 > e1 {
			temp = append(temp, arr.GetValue(s2))
			s2++
		} else if s2 > e2 {
			temp = append(temp, arr.GetValue(s1))
			s1++
		} else if arr.IsLT(s1, s2) {
			temp = append(temp, arr.GetValue(s1))
			s1++
		} else {
			temp = append(temp, arr.GetValue(s2))
			s2++
		}
	}

	for i, v := range temp {
		arr.SetValue(start+i, v)
	}
}
