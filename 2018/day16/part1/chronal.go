package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type OpcodeHandler func(int, int, int)

var opcodes map[int]OpcodeHandler

var registers [4]int

func init() {
	opcodes = make(map[int]OpcodeHandler)

	// addr
	opcodes[0] = func(a, b, c int) {
		registers[c] = registers[a] + registers[b]
	}

	// addi
	opcodes[1] = func(a, b, c int) {
		registers[c] = registers[a] + b
	}

	// mulr
	opcodes[2] = func(a, b, c int) {
		registers[c] = registers[a] * registers[b]
	}

	// muli
	opcodes[3] = func(a, b, c int) {
		registers[c] = registers[a] * b
	}

	// banr
	opcodes[4] = func(a, b, c int) {
		registers[c] = registers[a] & registers[b]
	}

	// bani
	opcodes[5] = func(a, b, c int) {
		registers[c] = registers[a] & b
	}

	// borr
	opcodes[6] = func(a, b, c int) {
		registers[c] = registers[a] | registers[b]
	}

	// bori
	opcodes[7] = func(a, b, c int) {
		registers[c] = registers[a] | b
	}

	// setr
	opcodes[8] = func(a, b, c int) {
		registers[c] = registers[a]
	}

	// seti
	opcodes[9] = func(a, b, c int) {
		registers[c] = a
	}

	// gtir
	opcodes[10] = func(a, b, c int) {
		if a > registers[b] {
			registers[c] = 1
		} else {
			registers[c] = 0
		}
	}

	// gtri
	opcodes[11] = func(a, b, c int) {
		if registers[a] > b {
			registers[c] = 1
		} else {
			registers[c] = 0
		}
	}

	// gtrr
	opcodes[12] = func(a, b, c int) {
		if registers[a] > registers[b] {
			registers[c] = 1
		} else {
			registers[c] = 0
		}
	}

	// eqir
	opcodes[13] = func(a, b, c int) {
		if a == registers[b] {
			registers[c] = 1
		} else {
			registers[c] = 0
		}
	}

	// eqri
	opcodes[14] = func(a, b, c int) {
		if registers[a] == b {
			registers[c] = 1
		} else {
			registers[c] = 0
		}
	}

	// eqrr
	opcodes[15] = func(a, b, c int) {
		if registers[a] == registers[b] {
			registers[c] = 1
		} else {
			registers[c] = 0
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var currentRegisters [4]int
	var currentOpcode [4]int
	var total int

	for i := 0; scanner.Scan(); i++ {
		if i%4 == 0 && !strings.HasPrefix(scanner.Text(), "Before") {
			break
		}

		if scanner.Text() == "" {
			continue
		}

		var a, b, c, d int
		if i%4 == 0 {
			fmt.Println("before", scanner.Text())
			fmt.Sscanf(scanner.Text(), "Before: [%d, %d, %d, %d]", &a, &b, &c, &d)
			currentRegisters[0] = a
			currentRegisters[1] = b
			currentRegisters[2] = c
			currentRegisters[3] = d
		} else if i%4 == 1 {
			fmt.Sscanf(scanner.Text(), "%d %d %d %d", &a, &b, &c, &d)
			currentOpcode[0] = a
			currentOpcode[1] = b
			currentOpcode[2] = c
			currentOpcode[3] = d
		} else if i%4 == 2 {
			fmt.Println("after", scanner.Text())
			fmt.Sscanf(scanner.Text(), "After:  [%d, %d, %d, %d]", &a, &b, &c, &d)

			var olala int
			for _, handler := range opcodes {
				registers = currentRegisters
				handler(currentOpcode[1], currentOpcode[2], currentOpcode[3])

				if registers[0] == a && registers[1] == b && registers[2] == c && registers[3] == d {
					olala++
				}

				if olala >= 3 {
					total++
					break
				}
			}
		}
	}

	fmt.Println("part1:", total)
}
