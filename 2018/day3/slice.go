package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(areas [][]int) int {
	var total int

	for _, area := range areas {
		for _, claim := range area {
			if claim > 1 {
				total++
			}
		}
	}

	return total
}

func part2(areas [][]int, claims [][]int) int {
	for _, claim := range claims {
		id := claim[0]
		leftLeft := claim[1]
		topTop := claim[2]
		width := claim[3]
		height := claim[4]

		var broke bool

	Loop:
		for x := leftLeft; x < leftLeft+width; x++ {
			for y := topTop; y < topTop+height; y++ {
				if areas[x][y] > 1 {
					broke = true
					break Loop
				}
			}
		}

		if !broke {
			return id
		}
	}

	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	areas := make([][]int, 1000)
	for i := range areas {
		areas[i] = make([]int, 1000)
	}

	var claims [][]int

	var id, leftLeft, topTop, width, height int

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d", &id, &leftLeft, &topTop, &width, &height)

		claims = append(claims, []int{id, leftLeft, topTop, width, height})

		for x := leftLeft; x < leftLeft+width; x++ {
			for y := topTop; y < topTop+height; y++ {
				areas[x][y] += 1
			}
		}
	}

	fmt.Println("part1: ", part1(areas))
	fmt.Println("part2: ", part2(areas, claims))
}
