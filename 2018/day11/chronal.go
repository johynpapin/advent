package main

import (
	"fmt"
)

const SERIAL = 7857

func getPowerLevel(serial, x, y int) int {
	level := ((x+10)*y + serial) * (x + 10)
	level = level%1000 - level%100
	level /= 100
	return level - 5
}

func main() {
	var grid [300][300]int

	for y := 0; y < 300; y++ {
		for x := 0; x < 300; x++ {
			grid[y][x] = getPowerLevel(SERIAL, x+1, y+1)
		}
	}

	var maxX, maxY, max, level int

	for y := 0; y < 297; y++ {
		for x := 0; x < 297; x++ {
			level = grid[y][x] + grid[y+1][x] + grid[y+2][x] + grid[y][x+1] + grid[y+1][x+1] + grid[y+2][x+1] + grid[y][x+2] + grid[y+1][x+2] + grid[y+2][x+2]

			if level > max {
				max = level
				maxX = x
				maxY = y
			}
		}
	}

	fmt.Printf("part1: %d,%d %d\n", maxX+1, maxY+1, max)

	maxX, maxY, max, level = 0, 0, 0, 0

	var maxI int

	grid2 := grid

	for i := 1; i <= 300; i++ {
		fmt.Println("i =", i)

		for y := 0; y < 300-i; y++ {
			for x := 0; x < 300-i; x++ {
				level = grid2[y][x]

				for yy := y; yy <= y+i-1; yy++ {
					level += grid[yy][x+i]
				}

				for xx := x; xx <= x+i; xx++ {
					level += grid[y+i][xx]
				}

				grid2[y][x] = level

				if level > max {
					max = level
					maxX = x
					maxY = y
					maxI = i
				}
			}
		}
	}

	fmt.Printf("part2: %d,%d,%d %d\n", maxX+1, maxY+1, maxI+1, max)
}
