package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1(frequencies []int) int {
	var total int

	for _, freq := range frequencies {
		total += freq
	}

	return total
}

func part2(frequencies []int) int {
	var frequency int

	seens := make(map[int]bool)

	seens[0] = true

	for {
		for _, freq := range frequencies {
			frequency += freq

			if seens[frequency] {
				return frequency
			}
			seens[frequency] = true
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var frequencies []int

	for scanner.Scan() {
		freq, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}

		frequencies = append(frequencies, freq)
	}

	fmt.Println("part1: ", part1(frequencies))
	fmt.Println("part2: ", part2(frequencies))
}
