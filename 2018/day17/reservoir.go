package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Tile struct {
	Clay    bool
	Water   bool
	Falling bool
}

type Input struct {
	XMin int
	XMax int
	YMin int
	YMax int
}

func main() {
	var inputs []*Input
	xMin, yMin, xMax, yMax := math.MaxInt32, math.MaxInt32, 0, 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ", ")

		x, y := s[0], s[1]

		if x[0] == 'y' {
			x, y = y, x
		}

		input := &Input{}

		x, y = x[2:], y[2:]

		xx := strings.Split(x, "..")
		input.XMin, _ = strconv.Atoi(xx[0])
		if len(xx) == 1 {
			input.XMax = input.XMin
		} else {
			input.XMax, _ = strconv.Atoi(xx[1])
		}

		yy := strings.Split(y, "..")
		input.YMin, _ = strconv.Atoi(yy[0])
		if len(yy) == 1 {
			input.YMax = input.YMin
		} else {
			input.YMax, _ = strconv.Atoi(yy[1])
		}

		if input.XMin < xMin {
			xMin = input.XMin
		}

		if input.YMin < yMin {
			yMin = input.YMin
		}

		if input.XMax > xMax {
			xMax = input.XMax
		}

		if input.YMax > yMax {
			yMax = input.YMax
		}

		inputs = append(inputs, input)
	}

	yMax++
	xMax = xMax - xMin + 3

	grid := make([][]*Tile, yMax)
	for y := range grid {
		grid[y] = make([]*Tile, xMax)
		for x := range grid[y] {
			grid[y][x] = &Tile{}
		}
	}

	for _, input := range inputs {
		input.XMin -= xMin - 1
		input.XMax -= xMin - 1

		for y := input.YMin; y <= input.YMax; y++ {
			for x := input.XMin; x <= input.XMax; x++ {
				grid[y][x].Clay = true
			}
		}
	}

	grid[0][500-xMin+1].Falling = true

	for {
		var changes int
		for y, line := range grid {
			var botLine []*Tile
			if y+1 < yMax {
				botLine = grid[y+1]
			}

			var found bool
			for x, tile := range line {
				if tile.Water || tile.Falling {
					found = true
				}

				if tile.Falling && botLine != nil {
					if !(botLine[x].Clay || botLine[x].Water || botLine[x].Falling) {
						if !botLine[x].Falling {
							botLine[x].Falling = true
							changes++
						}
					} else if botLine[x].Clay || botLine[x].Water {
						startClayIndex := -1
						startAirIndex := -1
						for x2 := x; x2 >= 0; x2-- {
							if line[x2].Clay {
								startClayIndex = x2
								break
							}

							if !(botLine[x2].Clay || botLine[x2].Water) {
								startAirIndex = x2
								break
							}
						}

						clayIndex := -1
						airIndex := -1
						for x2 := x; x2 < xMax; x2++ {
							if line[x2].Clay {
								clayIndex = x2
								break
							}

							if !(botLine[x2].Clay || botLine[x2].Water) {
								airIndex = x2
								break
							}
						}

						if startClayIndex > -1 && clayIndex > -1 {
							if !tile.Water {
								tile.Water = true
								changes++
							}
						} else {
							start := startClayIndex
							end := clayIndex

							if start == -1 {
								start = startAirIndex
							} else {
								start++
							}

							if end == -1 {
								end = airIndex
							} else {
								end--
							}

							for x2 := start; x2 <= end; x2++ {
								if !line[x2].Falling {
									line[x2].Falling = true
									changes++
								}
							}
						}
					}
				}

				if tile.Water {
					startX := -1
					for x2 := x; x2 >= 0; x2-- {
						if line[x2].Clay {
							startX = x2
							break
						}
					}

					if startX > -1 {
						clayIndex := -1
						for x2 := startX + 1; x2 < xMax; x2++ {
							if line[x2].Clay {
								clayIndex = x2
								break
							}

							if !(botLine[x2].Clay || botLine[x2].Water) {
								break
							}
						}

						if clayIndex > -1 {
							for x2 := startX + 1; x2 < clayIndex; x2++ {
								if line[x2].Falling || !line[x2].Water {
									line[x2].Falling = false
									line[x2].Water = true
									changes++
								}
							}
						}
					}
				}
			}

			if !found {
				fmt.Println(y)
				break
			}
		}

		//show(grid)

		if changes == 0 {
			break
		}
	}

	show(grid)

	fmt.Println("part1:", calc(grid, true)-yMin)
	fmt.Println("part2:", calc(grid, false))
}

func calc(grid [][]*Tile, olali bool) int {
	var olala int
	for _, line := range grid {
		for _, tile := range line {
			if tile.Water || (olali && tile.Falling) {
				olala++
			}
		}
	}

	return olala
}

func show(grid [][]*Tile) {
	for _, line := range grid {
		var rawLine string

		for _, tile := range line {
			if tile.Clay {
				rawLine += "#"
			} else if tile.Falling {
				rawLine += "|"
			} else if tile.Water {
				rawLine += "~"
			} else {
				rawLine += "."
			}
		}

		fmt.Println(rawLine)
	}

	fmt.Println()
}
