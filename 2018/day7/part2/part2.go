package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Worker struct {
	Working   bool
	Producing int
	TimeLeft  int
}

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
	workers := []*Worker{&Worker{}, &Worker{}, &Worker{}, &Worker{}, &Worker{}}

	var time int

	for len(steps) < numSteps {
		for _, worker := range workers {
			if worker.Working {
				worker.TimeLeft--
			}
		}

		for _, worker := range workers {
			if worker.TimeLeft == 0 {
				worker.Working = false
				steps = append(steps, worker.Producing)

				for _, step := range nexts[worker.Producing] {
					for i, val := range dependencies[step] {
						if val == worker.Producing {
							dependencies[step] = remove(dependencies[step], i)
							break
						}
					}

					if len(dependencies[step]) == 0 {
						availables = append(availables, step)
					}
				}
			}
		}

		sort.Ints(availables)

		for _, worker := range workers {
			if len(availables) == 0 {
				break
			}

			if !worker.Working {
				worker.Working = true
				worker.Producing = availables[0]
				worker.TimeLeft = availables[0] + 61

				availables = availables[1:]
			}
		}

		output := fmt.Sprintf("%4d  ", time)
		for _, worker := range workers {
			if worker.Working {
				output += fmt.Sprintf("  %c  ", worker.Producing+65)
			} else {
				output += "  .  "
			}
		}
		fmt.Println(output)

		time++
	}

	fmt.Println("part2:", time-1)
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
