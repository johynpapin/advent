package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(ids []string) int {
	var twoCount, threeCount int

	for _, id := range ids {
		letters := make(map[rune]int)

		for _, letter := range id {
			letters[letter] += 1
		}

		var twoCounted, threeCounted bool

		for _, count := range letters {
			if !twoCounted && count == 2 {
				twoCount++
				twoCounted = true
			} else if !threeCounted && count == 3 {
				threeCount++
				threeCounted = true
			}
		}
	}

	return twoCount * threeCount
}

func areCorrectBoxes(a string, b string) (bool, string) {
	var common string
	var differed bool

	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			common += string(a[i])
		} else if differed {
			return false, ""
		} else {
			differed = true
		}
	}

	return true, common
}

func part2(ids []string) string {
	for i, id := range ids {
		for j, id2 := range ids {
			if i != j {
				if correct, common := areCorrectBoxes(id, id2); correct {
					return common
				}
			}
		}
	}

	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var ids []string

	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}

	fmt.Println("part1: ", part1(ids))
	fmt.Println("part2: ", part2(ids))
}
