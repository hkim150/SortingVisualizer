package internal

import (
	"fmt"
	"sortingvisualizer/internal/algorithm"
	"strings"
)

var sorters = map[string]sorterFunc{
	"bubble": algorithm.Bubble,
}

type sorterFunc func(data []int) error

func (f sorterFunc) Sort(data []int) error {
	return f(data)
}

func Sort(data []int, algorithm string) error {
	sortFunc, ok := sorters[algorithm]
	if !ok {
		keys := make([]string, 0, len(sorters))
		for key := range sorters {
			keys = append(keys, key)
		}

		algorithms := strings.Join(keys, ", ")

		return fmt.Errorf("Unknown sorting algorithm. Choose from: %v\n", algorithms)
	}

	return sortFunc.Sort(data)
}
