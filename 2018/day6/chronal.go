package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

type GridEntry struct {
	Point    bool
	Distance int
	ID       int
	Sandwich bool
}

func showGrid(grid [][]*GridEntry) {
	for _, line := range grid {
		r := ""

		for _, entry := range line {
			if entry.Point {
				r += string(65 + entry.ID)
			} else if entry.Sandwich {
				r += "."
			} else if entry.ID == -1 || entry.Distance == 0 {
				r += " "
			} else {
				r += string(97 + entry.ID)
			}
		}

		fmt.Println(r)
	}

	fmt.Println()
	fmt.Println()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var points []Point

	var gridX, gridY int
	var point Point
	var id int

	for scanner.Scan() {
		id++

		fmt.Sscanf(scanner.Text(), "%d, %d", &point.X, &point.Y)
		points = append(points, point)

		if point.X >= gridX {
			gridX = point.X + 1
		}

		if point.Y >= gridY {
			gridY = point.Y + 1
		}
	}

	grid := make([][]*GridEntry, gridY)

	for i := range grid {
		grid[i] = make([]*GridEntry, gridX)

		for j := range grid[i] {
			grid[i][j] = &GridEntry{
				ID: -1,
			}
		}
	}

	for id, point := range points {
		grid[point.Y][point.X].Point = true
		grid[point.Y][point.X].ID = id
	}

	for id, point := range points {
		x, y := point.X, point.Y
		var stop bool

		for offset := 1; !stop; offset++ {
			stop = true

			for ox := 0; ox <= offset; ox++ {
				coordinates := []Point{{x - ox, y + offset - ox}, {x + ox, y + offset - ox}, {x - ox, y - offset + ox}, {x + ox, y - offset + ox}}

				if ox == 0 {
					coordinates = []Point{{x, y + offset}, {x, y - offset}}
				} else if ox == offset {
					coordinates = []Point{{x - ox, y}, {x + ox, y}}
				}

				for _, c := range coordinates {
					if c.X < 0 || c.X >= gridX || c.Y < 0 || c.Y >= gridY {
						continue
					}

					entry := grid[c.Y][c.X]
					if entry.Point {
						continue
					}

					if entry.Distance == 0 && entry.ID == -1 {
						stop = false

						entry.ID = id
						entry.Distance = offset

						continue
					}

					if entry.ID != id && entry.Distance == offset {
						entry.Sandwich = true
					} else if offset < entry.Distance {
						stop = false

						entry.Sandwich = false
						entry.Distance = offset
						entry.ID = id
					}
				}
			}
		}
	}

	showGrid(grid)

	areas := make(map[int]int, len(points))

	for y := 0; y < gridY; y++ {
		if !grid[y][0].Sandwich {
			areas[grid[y][0].ID] = -1
		}

		if !grid[y][gridX-1].Sandwich {
			areas[grid[y][gridX-1].ID] = -1
		}
	}

	for x := 0; x < gridX; x++ {
		if !grid[0][x].Sandwich {
			areas[grid[0][x].ID] = -1
		}

		if !grid[gridY-1][x].Sandwich {
			areas[grid[gridY-1][x].ID] = -1
		}
	}
	for _, line := range grid {
		for _, entry := range line {
			if !entry.Sandwich && areas[entry.ID] != -1 && !(!entry.Point && entry.Distance == 0) {
				areas[entry.ID]++
			}
		}
	}

	max := 0

	for _, value := range areas {
		if value > max {
			max = value
		}
	}

	fmt.Println("part1:", max)

	size := 0

	for y := 0; y < gridY; y++ {
		for x := 0; x < gridX; x++ {
			totalDistance := 0

			for _, point := range points {
				totalDistance += abs(x-point.X) + abs(y-point.Y)
			}

			if totalDistance < 10000 {
				size++
			}
		}
	}

	fmt.Println("part2:", size)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
