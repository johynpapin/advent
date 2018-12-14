package main

import (
	"bufio"
	"fmt"
	"os"
)

func show(state string) {
	fmt.Println(state)
	//	fmt.Println(strings.TrimRight(strings.TrimLeft(state, "."), "."))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var state string
	fmt.Sscanf(scanner.Text(), "initial state: %s", &state)

	patterns := make(map[string]string)
	scanner.Scan()
	var rawPattern string
	var rawPlant string
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s => %s", &rawPattern, &rawPlant)

		patterns[rawPattern] = rawPlant
	}

	offset := -5

	state = "....." + state + "....."

	var i int

	for ; ; i++ {
		tmpState := state

		for j := 2; j < len(state)-2; j++ {
			pattern := state[j-2 : j+3]
			tmpState = tmpState[:j] + patterns[pattern] + tmpState[j+1:]
		}

		for tmpState[0] != '#' {
			tmpState = tmpState[1:]
			offset++
		}

		for tmpState[len(tmpState)-1] != '#' {
			tmpState = tmpState[:len(tmpState)-1]
		}

		tmpState = "....." + tmpState + "....."
		offset -= 5

		if tmpState == state {
			break
		}

		state = tmpState
	}

	offset += 50000000000 - (i + 1)

	show(state)

	var sum int
	for i := 0; i < len(state); i++ {
		if state[i] == '#' {
			sum += offset + i
		}
	}

	fmt.Println("part1:", sum)
}
