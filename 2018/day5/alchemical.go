package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func apply(rawLine string) string {
	line := []rune(rawLine)

Loop:
	for {
		var old rune
		for i, e := range line {
			if e != old && unicode.ToUpper(old) == unicode.ToUpper(e) {
				line = append(line[:i-1], line[i+1:]...)
				continue Loop
			}

			old = e
		}

		break Loop
	}

	return string(line)
}

func main() {
	rawLine := ""
	fmt.Scanln(&rawLine)

	fmt.Println("part1:", len(apply(rawLine)))

	min := math.MaxInt32
	for _, c := range "abcdefghijklmnopqrstuvwxyz" {
		replacedLine := strings.Replace(strings.Replace(rawLine, string(c), "", -1), string(unicode.ToUpper(c)), "", -1)

		l := len(apply(replacedLine))

		if l < min {
			min = l
		}
	}

	fmt.Println("part2:", min)
}
