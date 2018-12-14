package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	//	"time"
)

type Direction int

const (
	TOP Direction = iota
	RIGHT
	BOTTOM
	LEFT
)

type Cart struct {
	Direction        Direction
	X, Y             int
	LastIntersection int
}

func main() {
	var carts []*Cart
	var grid []string

	scanner := bufio.NewScanner(os.Stdin)
	for y := 0; scanner.Scan(); y++ {
		var line string

		for x, c := range scanner.Text() {
			switch c {
			case '^':
				carts = append(carts, &Cart{TOP, x, y, 0})
				line += "|"
			case '<':
				carts = append(carts, &Cart{LEFT, x, y, 0})
				line += "-"
			case '>':
				carts = append(carts, &Cart{RIGHT, x, y, 0})
				line += "-"
			case 'v':
				carts = append(carts, &Cart{BOTTOM, x, y, 0})
				line += "|"
			default:
				line += string(c)
			}
		}

		grid = append(grid, line)
	}

	for i := 0; ; i++ {
		fmt.Println("-->", i)

		sort.Slice(carts, func(i, j int) bool {
			if carts[i].Y < carts[j].Y {
				return true
			}

			return carts[i].X < carts[j].X
		})

		for _, cart := range carts {
			switch cart.Direction {
			case TOP:
				cart.Y--
			case BOTTOM:
				cart.Y++
			case RIGHT:
				cart.X++
			case LEFT:
				cart.X--
			}

			switch grid[cart.Y][cart.X] {
			case '+':
				if cart.LastIntersection == 0 {
					cart.Direction = (cart.Direction - 1) % 4
					if cart.Direction < 0 {
						cart.Direction += 4
					}
					cart.LastIntersection++
				} else if cart.LastIntersection == 1 {
					cart.LastIntersection++
				} else {
					cart.Direction = (cart.Direction + 1) % 4
					cart.LastIntersection = 0
				}
			case '/':
				if cart.Direction == TOP {
					cart.Direction = RIGHT
				} else if cart.Direction == LEFT {
					cart.Direction = BOTTOM
				} else if cart.Direction == RIGHT {
					cart.Direction = TOP
				} else {
					cart.Direction = LEFT
				}
			case '\\':
				if cart.Direction == TOP {
					cart.Direction = LEFT
				} else if cart.Direction == RIGHT {
					cart.Direction = BOTTOM
				} else if cart.Direction == BOTTOM {
					cart.Direction = RIGHT
				} else {
					cart.Direction = TOP
				}
			}
		}
		/*
			for y, line := range grid {
			Loop:
				for x, c := range line {
					for _, cart := range carts {
						if cart.Y == y && cart.X == x {
							fmt.Print("C")
							continue Loop
						}
					}

					fmt.Print(string(c))
				}
				fmt.Println()
			}
		*/
		for i, cart1 := range carts {
			for j, cart2 := range carts {
				if i != j && cart1.Y == cart2.Y && cart1.X == cart2.X {
					fmt.Printf("Collision => %d, %d\n", cart1.X, cart1.Y)
					return
				}
			}
		}

		fmt.Println()

		//time.Sleep(time.Second)
	}

	fmt.Println("part1:")
}
