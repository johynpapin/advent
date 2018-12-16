package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	X, Y int
}

func (p0 Point) IsEqual(p1 Point) bool {
	return p0.X == p1.X && p0.Y == p1.Y
}

func (p Point) Neighbours(gridX, gridY int) []Point {
	var neighbours []Point

	if p.Y-1 >= 0 {
		neighbours = append(neighbours, Point{p.X, p.Y - 1})
	}

	if p.X-1 >= 0 {
		neighbours = append(neighbours, Point{p.X - 1, p.Y})
	}

	if p.X+1 < gridX {
		neighbours = append(neighbours, Point{p.X + 1, p.Y})
	}

	if p.Y+1 < gridY {
		neighbours = append(neighbours, Point{p.X, p.Y + 1})
	}

	return neighbours
}

type Unit struct {
	Pos       Point
	Goblin    bool
	HitPoints int
}

type Node struct {
	Pos  Point
	Wall bool
	Unit *Unit
}

type Grid [][]*Node

func show(grid Grid) {
	for _, line := range grid {
		var rawLine string
		var hps string

		for _, node := range line {
			if node.Wall {
				rawLine += "#"
			} else if node.Unit != nil {
				if hps != "" {
					hps += ", "
				}

				if node.Unit.Goblin {
					rawLine += "G"
					hps += fmt.Sprintf("G(%d)", node.Unit.HitPoints)
				} else {
					rawLine += "E"
					hps += fmt.Sprintf("E(%d)", node.Unit.HitPoints)
				}
			} else {
				rawLine += "."
			}
		}

		fmt.Println(rawLine + "   " + hps)
	}
}

var gridX, gridY int

func main() {
	var inputs []string

	scanner := bufio.NewScanner(os.Stdin)
	for y := 0; scanner.Scan(); y++ {
		inputs = append(inputs, scanner.Text())
	}

PowerLoop:
	for power := 4; ; power++ {
		var grid Grid
		var units []*Unit

		for y := 0; y < len(inputs); y++ {
			var line []*Node

			for x, c := range inputs[y] {
				switch c {
				case '#':
					line = append(line, &Node{Pos: Point{x, y}, Wall: true})
				case '.':
					line = append(line, &Node{Pos: Point{x, y}})
				case 'G':
					unit := &Unit{Point{x, y}, true, 200}
					line = append(line, &Node{Pos: Point{x, y}, Unit: unit})
					units = append(units, unit)
				case 'E':
					unit := &Unit{Point{x, y}, false, 200}
					line = append(line, &Node{Pos: Point{x, y}, Unit: unit})
					units = append(units, unit)
				}
			}

			grid = append(grid, line)
		}

		gridX, gridY = len(grid[0]), len(grid)

		fullRound := 0
	MainLoop:
		for i := 1; ; i++ {
			sort.Slice(units, func(i, j int) bool {
				return units[i].Pos.Y < units[j].Pos.Y || units[i].Pos.Y == units[j].Pos.Y && units[i].Pos.X < units[j].Pos.X
			})

		Olali:
			for id, unit := range units {
				if unit == nil {
					continue
				}

				var targets []*Unit
				for targetID, target := range units {
					if target == nil {
						continue
					}

					if id != targetID && target.Goblin != unit.Goblin {
						targets = append(targets, target)
					}
				}

				if len(targets) == 0 {
					break MainLoop
				}

				var inRangeNodes []*Node
				for _, target := range targets {
					for _, neighbour := range target.Pos.Neighbours(gridX, gridY) {
						node := grid[neighbour.Y][neighbour.X]
						if node.Unit == unit || (node.Unit == nil && !node.Wall) {
							inRangeNodes = append(inRangeNodes, node)
						}
					}
				}

				if len(inRangeNodes) == 0 {
					continue
				}

				for _, node := range inRangeNodes {
					if node.Pos.IsEqual(unit.Pos) {
						if attack(unit, grid, units, power) {
							continue PowerLoop
						}

						continue Olali
					}
				}

				move(unit, grid, inRangeNodes)
				if attack(unit, grid, units, power) {
					continue PowerLoop
				}
			}

			var units2 []*Unit
			for _, unit := range units {
				if unit != nil {
					units2 = append(units2, unit)
				}
			}
			units = units2

			fullRound++
		}

		var totalHP int
		for _, unit := range units {
			if unit != nil {
				totalHP += unit.HitPoints
			}
		}

		fmt.Println("part2:", fullRound*totalHP, power)
		return
	}
}

func move(unit *Unit, grid [][]*Node, inRangeNodes []*Node) {
	var bestMove *Point
	var bestNode *Node
	bestDist := -1

	for _, rangeNode := range inRangeNodes {
		dists := make([][]int, gridY)
		for i := range dists {
			dists[i] = make([]int, gridX)

			for j := range dists[i] {
				dists[i][j] = -1
			}
		}

		dists[rangeNode.Pos.Y][rangeNode.Pos.X] = 0

		points := []Point{rangeNode.Pos}
		var node *Node
	Olala:
		for i := 1; len(points) != 0; i++ {
			var nextPoints []Point
			for _, point := range points {
				for _, neighbour := range point.Neighbours(gridX, gridY) {
					if neighbour.Y == unit.Pos.Y && neighbour.X == unit.Pos.X {
						dists[unit.Pos.Y][unit.Pos.X] = i
						break Olala
					}

					node = grid[neighbour.Y][neighbour.X]
					if dists[neighbour.Y][neighbour.X] == -1 && !node.Wall && node.Unit == nil {
						dists[neighbour.Y][neighbour.X] = i
						nextPoints = append(nextPoints, neighbour)
					}
				}
			}
			points = nextPoints
		}

		if dists[unit.Pos.Y][unit.Pos.X] != -1 {
			minMoveDist := -1
			var minMovePoint Point
			for _, neighbour := range unit.Pos.Neighbours(gridX, gridY) {
				if minMoveDist == -1 || dists[neighbour.Y][neighbour.X] != -1 && dists[neighbour.Y][neighbour.X] < minMoveDist {
					minMovePoint = neighbour
					minMoveDist = dists[neighbour.Y][neighbour.X]
				}
			}

			distToUnit := dists[unit.Pos.Y][unit.Pos.X]
			move := minMovePoint

			if bestDist == -1 || distToUnit < bestDist || (distToUnit == bestDist && (rangeNode.Pos.Y < bestNode.Pos.Y || rangeNode.Pos.Y == bestNode.Pos.Y && rangeNode.Pos.X < bestNode.Pos.X)) {
				bestNode = rangeNode
				bestMove = &move
				bestDist = distToUnit
			}
		}
	}

	if bestMove != nil {
		grid[unit.Pos.Y][unit.Pos.X].Unit = nil
		unit.Pos = *bestMove
		grid[unit.Pos.Y][unit.Pos.X].Unit = unit
	}
}

func attack(unit *Unit, grid [][]*Node, units []*Unit, power int) bool {
	target := &Unit{HitPoints: 4242}
	var node *Node

	for _, neighbour := range unit.Pos.Neighbours(gridX, gridY) {
		node = grid[neighbour.Y][neighbour.X]
		if node.Unit != nil && node.Unit.Goblin != unit.Goblin && node.Unit.HitPoints < target.HitPoints {
			target = node.Unit
		}
	}

	if target.HitPoints == 4242 {
		return false
	}

	if unit.Goblin {
		target.HitPoints -= 3
	} else {
		target.HitPoints -= power
	}

	if target.HitPoints <= 0 {
		grid[target.Pos.Y][target.Pos.X].Unit = nil

		for i, unit := range units {
			if unit != nil {
				if unit.Pos.X == target.Pos.X && unit.Pos.Y == target.Pos.Y {
					units[i] = nil
					break
				}
			}
		}

		return !target.Goblin
	}

	return false
}
