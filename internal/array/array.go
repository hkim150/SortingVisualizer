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
	data         []int
	previousData []int // Track previous state for differential updates
	minVal       int
	maxVal       int
	compCount    int
	moveCount    int
	printDelay   time.Duration
	changedCols  map[int]bool // Track which columns have changed
	initialized  bool         // Track if we've done the initial full print
}

func NewArray(size int) *Array {
	arr := &Array{
		data:         make([]int, size),
		previousData: make([]int, size),
		minVal:       minElemSize,
		maxVal:       minElemSize + size - 1,
		printDelay:   40 * time.Millisecond,
		changedCols:  make(map[int]bool),
		initialized:  false,
	}

	for i := 0; i < size; i++ {
		arr.data[i] = arr.minVal + i
	}

	rand.Shuffle(size, func(i, j int) {
		arr.data[i], arr.data[j] = arr.data[j], arr.data[i]
	})

	// Initialize previousData to match data
	copy(arr.previousData, arr.data)

	return arr
}

func (a *Array) Len() int {
	return len(a.data)
}

func (a *Array) Swap(i, j int) {
	a.data[i], a.data[j] = a.data[j], a.data[i]
	a.moveCount++
	// Mark both columns as changed
	a.changedCols[i] = true
	a.changedCols[j] = true
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
	// Mark column as changed
	a.changedCols[i] = true
	a.Print()
}

func (a *Array) Print() {
	if a.Len() == 0 {
		return
	}

	// If this is the first print or we haven't initialized, do a full print
	if !a.initialized {
		a.printFull()
		a.initialized = true
		copy(a.previousData, a.data)
		a.changedCols = make(map[int]bool) // Clear changes after full print
		// Update stats after initial print
		a.updateStats()
		time.Sleep(a.printDelay)
		return
	}

	// Only update changed columns
	if len(a.changedCols) > 0 {
		a.printDifferential()
		copy(a.previousData, a.data)
		a.changedCols = make(map[int]bool) // Clear changes after update
	}

	// Always update the stats at the bottom
	a.updateStats()

	time.Sleep(a.printDelay)
}

func (a *Array) printFull() {
	clearScreen()
	numColors := len(rainbowColors)

	// Print the visualization
	for r := a.maxVal; r > 0; r-- {
		for _, val := range a.data {
			if r <= val {
				a.printColoredBlock(val, numColors)
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	
	// Print empty lines for stats (they will be filled by updateStats)
	fmt.Println()
	fmt.Println()
}

func (a *Array) printDifferential() {
	numColors := len(rainbowColors)

	// For each changed column, update only that column
	for col := range a.changedCols {
		// Move cursor to the top of the column and redraw it
		for r := a.maxVal; r > 0; r-- {
			// Calculate cursor position: row (a.maxVal - r + 1), column (col * 2 + 1)
			row := a.maxVal - r + 1
			cursorCol := col*2 + 1

			// Move cursor to position
			fmt.Printf("\033[%d;%dH", row, cursorCol)

			// Print the appropriate content for this position
			if r <= a.data[col] {
				a.printColoredBlock(a.data[col], numColors)
			} else {
				fmt.Print("  ")
			}
		}
	}
}

func (a *Array) printColoredBlock(val, numColors int) {
	// Calculate color based on value's position in the range
	ratio := 0.0
	if a.maxVal > a.minVal { // Avoid division by zero
		ratio = float64(val-a.minVal) / float64(a.maxVal-a.minVal)
	}

	// Map ratio to color index
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
}

func (a *Array) updateStats() {
	// Move cursor to stats position (after the visualization)
	statsRow := a.maxVal + 2
	fmt.Printf("\033[%d;1H", statsRow)
	fmt.Printf("Comparison Count: %d", a.compCount)

	fmt.Printf("\033[%d;1H", statsRow+1)
	fmt.Printf("Move Count: %d\n", a.moveCount)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
