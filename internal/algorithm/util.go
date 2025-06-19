package algorithm

import "math/rand"

func CreateData(n int) []int {
	data := make([]int, n, n)
	for i := range data {
		data[i] = i + 1
	}

	rand.Shuffle(n, func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})

	return data
}