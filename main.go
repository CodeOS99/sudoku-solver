package main

import "fmt"

var sudokuPuzzle = [9][9]int{
	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	{0, 0, 0, 0, 8, 0, 0, 7, 9},
}

func backTrack(grid *[9][9]int) {
	if recursivelyBackTrack(grid, 0, 0) {
		print("Solved :)\n")
	} else {
		print("No solution ):\n")
	}
}

func isSafe(grid *[9][9]int, i int, j int, k int) bool {
	for x := 0; x < 9; x++ {
		if grid[x][j] == k { // j remains constant to check for rows
			return false
		}
	}

	for y := 0; y < 9; y++ {
		if grid[i][y] == k {
			return false
		}
	}

	// Which quad?
	quadi := i / 3
	quadj := j / 3
	quadiStart := quadi * 3
	quadjStart := quadj * 3

	for x := quadiStart; x < quadiStart+3; x++ {
		for y := quadjStart; y < quadjStart+3; y++ {
			if grid[x][y] == k {
				return false
			}
		}
	}

	return true
}

func recursivelyBackTrack(grid *[9][9]int, i int, j int) bool {
	if i == 9 { // solved
		return true
	}

	if j == 9 { // use a recursive loop, if j is 9, then that means we want to advance a row
		j = 0
		i++
		if i == 9 {
			return true
		}
	}

	if grid[i][j] == 0 {
		for k := 1; k <= 9; k++ {
			if isSafe(grid, i, j, k) { // only try a number if it is safe
				grid[i][j] = k
				if recursivelyBackTrack(grid, i, j+1) {
					return true
				}
				grid[i][j] = 0
			}
		}
	} else { // otherwise its not empty
		return recursivelyBackTrack(grid, i, j+1) // so just continue :D
	}
	return false
}

func main() {
	backTrack(&sudokuPuzzle)
	for _, row := range sudokuPuzzle {
		for _, cell := range row {
			fmt.Printf("%d ", cell)
		}
		fmt.Println() // This will print a newline after each row
	}
}
