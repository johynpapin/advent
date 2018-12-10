package main

import (
	"bufio"
	"fmt"
	"os"
)

type Marble struct {
	Next     *Marble
	Previous *Marble
	Value    int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var numPlayers, lastMarble int

	fmt.Sscanf(scanner.Text(), "%d players; last marble is worth %d points", &numPlayers, &lastMarble)

	players := make([]int, numPlayers)

	currentMarble := &Marble{}
	currentMarble.Next = currentMarble
	currentMarble.Previous = currentMarble

	actualValue := 1

	var currentPlayer int

	for actualValue != lastMarble {
		currentPlayer = (actualValue - 1) % numPlayers

		if actualValue%23 == 0 {
			players[currentPlayer] += actualValue

			removedMarble := currentMarble.Previous.Previous.Previous.Previous.Previous.Previous.Previous

			removedMarble.Previous.Next = removedMarble.Next
			removedMarble.Next.Previous = removedMarble.Previous

			players[currentPlayer] += removedMarble.Value

			currentMarble = removedMarble.Next
		} else {
			marble := &Marble{Value: actualValue}

			marble.Next = currentMarble.Next.Next
			marble.Previous = currentMarble.Next

			currentMarble.Next.Next.Previous = marble
			currentMarble.Next.Next = marble

			currentMarble = marble
		}

		actualValue++
	}

	var bestScore int
	for _, score := range players {
		if score > bestScore {
			bestScore = score
		}
	}

	fmt.Println("part1:", bestScore)
}

func showState(turn int, marble *Marble, currentMarble *Marble) {
	if marble == nil || turn == 0 {
		fmt.Println("[-] (0)")
	} else {
		r := fmt.Sprintf("[%d] ", turn)

		pointer := marble

		for {
			if pointer == currentMarble {
				r += fmt.Sprintf(" (%d) ", pointer.Value)
			} else {
				r += fmt.Sprintf("  %d  ", pointer.Value)
			}

			pointer = pointer.Next

			if pointer == marble {
				break
			}
		}

		fmt.Println(r)
	}
}
