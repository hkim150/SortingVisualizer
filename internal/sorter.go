package internal

import (
	"fmt"
	"sortingvisualizer/internal/algorithm"
	"sortingvisualizer/internal/array"
	"strings"
)

var algorithms = map[string]sortFunc{
	"bubble":    algorithm.Bubble,
	"merge":     algorithm.Merge,
	"quick":     algorithm.Quick,
	"insertion": algorithm.Insertion,
	"selection": algorithm.Selection,
	"heap":      algorithm.Heap,
}

type sortFunc func(arr *array.Array)

func (f sortFunc) Sort(arr *array.Array) {
	f(arr)
}

func Sort(arr *array.Array, algorithm string) error {
	sortFunc, ok := algorithms[algorithm]
	if !ok {
		keys := make([]string, 0, len(algorithms))
		for key := range algorithms {
			keys = append(keys, key)
		}

		algorithms := strings.Join(keys, ", ")

		return fmt.Errorf("Unknown sorting algorithm. Choose from: %v\n", algorithms)
	}

	sortFunc.Sort(arr)
	return nil
}

func Algorithms() string {
	algos := make([]string, 0, len(algorithms))
	for key := range algorithms {
		algos = append(algos, key)
	}

	return strings.Join(algos, ", ")
}
