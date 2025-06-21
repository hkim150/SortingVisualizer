package internal

import (
	"fmt"
	"sortingvisualizer/internal/algorithm"
	"sortingvisualizer/internal/array"
	"strings"
)

var sorters = map[string]sorterFunc{
	"bubble": algorithm.Bubble,
	"merge":  algorithm.Merge,
	"quick": algorithm.Quick,
}

type sorterFunc func(arr *array.Array) error

func (f sorterFunc) Sort(arr *array.Array) error {
	return f(arr)
}

func Sort(arr *array.Array, algorithm string) error {
	sortFunc, ok := sorters[algorithm]
	if !ok {
		keys := make([]string, 0, len(sorters))
		for key := range sorters {
			keys = append(keys, key)
		}

		algorithms := strings.Join(keys, ", ")

		return fmt.Errorf("Unknown sorting algorithm. Choose from: %v\n", algorithms)
	}

	return sortFunc.Sort(arr)
}
