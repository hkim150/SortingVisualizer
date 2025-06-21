package array

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	minElemSize = 1 // minimum height of the element
)

var (
	rainbowColors = []struct{ r, g, b int }{
		{255, 0, 0},   // Red
		{255, 127, 0}, // Orange
		{255, 255, 0}, // Yellow
		{0, 255, 0},   // Green
		{0, 0, 255},   // Blue
		{75, 0, 130},  // Indigo
		{128, 0, 128}, // Purple
	}
)

type Array struct {
	data       []int
	minVal     int
	maxVal     int
	compCount  int
	moveCount  int
	printDelay time.Duration
}

func NewArray(size int) *Array {
	arr := &Array{
		data:       make([]int, size),
		minVal:     minElemSize,
		maxVal:     minElemSize + size - 1,
		printDelay: 40 * time.Millisecond,
	}

	for i := 0; i < size; i++ {
		arr.data[i] = arr.minVal + i
	}

	rand.Shuffle(size, func(i, j int) {
		arr.data[i], arr.data[j] = arr.data[j], arr.data[i]
	})

	return arr
}

func (a *Array) Len() int {
	return len(a.data)
}

func (a *Array) Swap(i, j int) {
	a.data[i], a.data[j] = a.data[j], a.data[i]
	a.moveCount++
	a.Print()
}

func (a *Array) IsLT(i, j int) bool {
	a.compCount++
	return a.data[i] < a.data[j]
}

func (a *Array) IsLTE(i, j int) bool {
	a.compCount++
	return a.data[i] <= a.data[j]
}

func (a *Array) IsGT(i, j int) bool {
	a.compCount++
	return a.data[i] > a.data[j]
}

func (a *Array) GetValue(i int) int {
	return a.data[i]
}

func (a *Array) SetValue(i, v int) {
	a.moveCount++
	a.data[i] = v
	a.Print()
}

func (a *Array) Print() {
	if a.Len() == 0 {
		return
	}

	clearScreen()

	numColors := len(rainbowColors)

	// Print the visualization
	for r := a.maxVal; r > 0; r-- {
		for _, val := range a.data {
			if r <= val {
				// Calculate color based on value's position in the range
				ratio := 0.0
				if a.maxVal > a.minVal { // Avoid division by zero
					ratio = float64(val-a.minVal) / float64(a.maxVal-a.minVal)
				}

				// Map ratio to color index
				// This gives us which two colors to interpolate between
				colorIndex := ratio * float64(numColors-1)
				lowerIndex := int(colorIndex)
				upperIndex := lowerIndex + 1
				if upperIndex >= numColors {
					upperIndex = numColors - 1
				}

				// Calculate how far we are between the two colors (0.0 to 1.0)
				colorRatio := colorIndex - float64(lowerIndex)

				// Get the two colors to interpolate between
				c1 := rainbowColors[lowerIndex]
				c2 := rainbowColors[upperIndex]

				// Interpolate between the two colors
				red := int(float64(c1.r) + colorRatio*float64(c2.r-c1.r))
				green := int(float64(c1.g) + colorRatio*float64(c2.g-c1.g))
				blue := int(float64(c1.b) + colorRatio*float64(c2.b-c1.b))

				// Print colored block
				fmt.Printf("\033[38;2;%d;%d;%dm██\033[0m", red, green, blue)
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	fmt.Println("Comparison Count: ", a.compCount)
	fmt.Println("Move Count: ", a.moveCount)

	time.Sleep(a.printDelay)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
