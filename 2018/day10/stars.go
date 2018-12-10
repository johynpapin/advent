package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Velocity struct {
	X, Y int
}

type Point struct {
	X, Y     int
	Velocity Velocity
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var positionX, positionY, velocityX, velocityY int

	var points []*Point

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "position=<%d, %d> velocity=<%d, %d>", &positionX, &positionY, &velocityX, &velocityY)

		points = append(points, &Point{
			X: positionX,
			Y: positionY,
			Velocity: Velocity{
				X: velocityX,
				Y: velocityY,
			},
		})
	}

	var seconds int

	for {
		for _, point := range points {
			point.X += point.Velocity.X
			point.Y += point.Velocity.Y
		}

		minX, maxX, minY, maxY := getGridSize(points)

		if minX < 0 {
			minX *= -1
		}

		if minY < 0 {
			minY *= -1
		}

		seconds++

		gridX := minX + maxX + 1
		gridY := minY + maxY + 1

		if gridY < 300 && gridX < 200 {
			showGrid(points, minX, minY, gridX, gridY)
			fmt.Println("after", seconds, "seconds")

			time.Sleep(time.Second)
		}
	}
}

func getGridSize(points []*Point) (minX, maxX, minY, maxY int) {
	for _, point := range points {
		if point.X > maxX {
			maxX = point.X
		}

		if point.X < minX {
			minX = point.X
		}

		if point.Y > maxY {
			maxY = point.Y
		}

		if point.Y < minY {
			minY = point.Y
		}
	}

	return
}

func showGrid(points []*Point, minX, minY, gridX, gridY int) {
	grid := make([][]bool, gridY)

	for i := range grid {
		grid[i] = make([]bool, gridX)
	}

	for _, point := range points {
		grid[point.Y+minY][point.X+minX] = true
	}

	for _, line := range grid {
		var s string

		for _, point := range line {
			if point {
				s += "#"
			} else {
				s += "."
			}
		}

		fmt.Println(s)
	}
}
