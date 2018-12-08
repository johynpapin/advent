package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var stepA, stepB int
	var numSteps int

	dependencies := make(map[int][]int)
	nexts := make(map[int][]int)

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "Step %c must be finished before step %c can begin.", &stepA, &stepB)

		stepA -= 65
		stepB -= 65

		if stepA > numSteps {
			numSteps = stepA
		}

		if stepB > numSteps {
			numSteps = stepB
		}

		dependencies[stepB] = append(dependencies[stepB], stepA)
		nexts[stepA] = append(nexts[stepA], stepB)
	}

	numSteps++

	var availables []int

	for step := 0; step < numSteps; step++ {
		if dependencies[step] == nil {
			availables = append(availables, step)
		}
	}

	var steps []int

	for len(steps) < numSteps {
		sort.Ints(availables)

		steps = append(steps, availables[0])

		for _, step := range nexts[availables[0]] {
			for i, val := range dependencies[step] {
				if val == availables[0] {
					dependencies[step] = remove(dependencies[step], i)
					break
				}
			}

			if len(dependencies[step]) == 0 {
				availables = append(availables, step)
			}
		}

		availables = availables[1:]
	}

	var part1 string

	for _, step := range steps {
		part1 += string(step + 65)
	}

	fmt.Println("part1:", part1)
	fmt.Println("part2:")
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
