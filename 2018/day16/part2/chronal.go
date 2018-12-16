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
	var maybe [16][]int

	for i := range maybe {
		maybe[i] = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	}

	for i := 0; scanner.Scan(); i++ {
		if i%4 == 0 && !strings.HasPrefix(scanner.Text(), "Before") {
			break
		}

		if scanner.Text() == "" {
			continue
		}

		var a, b, c, d int
		if i%4 == 0 {
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
			fmt.Sscanf(scanner.Text(), "After:  [%d, %d, %d, %d]", &a, &b, &c, &d)

			op := currentOpcode[0]

			var olala []int
			for _, id := range maybe[op] {
				olala = append(olala, id)
			}

			for _, id := range olala {
				registers = currentRegisters
				opcodes[id](currentOpcode[1], currentOpcode[2], currentOpcode[3])

				if registers[0] != a || registers[1] != b || registers[2] != c || registers[3] != d {
					doko := find(maybe[op], id)
					if doko > -1 {
						maybe[op] = remove(maybe[op], doko)
					}
				}
			}
		}
	}

	var queue []int
	for _, ops := range maybe {
		if len(ops) == 1 {
			queue = append(queue, ops[0])
		}
	}

	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]

		for i := range maybe {
			if len(maybe[i]) > 1 {
				doko := find(maybe[i], top)
				if doko != -1 {
					maybe[i] = remove(maybe[i], doko)

					if len(maybe[i]) == 1 {
						queue = append(queue, maybe[i][0])
					}
				}
			}
		}
	}

	registers[0] = 0
	registers[1] = 0
	registers[2] = 0
	registers[3] = 0

	var a, b, c, d int
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		fmt.Sscanf(scanner.Text(), "%d %d %d %d", &a, &b, &c, &d)

		opcodes[maybe[a][0]](b, c, d)
	}

	fmt.Println("part2:", registers[0])
}

func remove(slice []int, index int) []int {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func find(slice []int, nani int) int {
	for i, kore := range slice {
		if kore == nani {
			return i
		}
	}

	return -1
}
