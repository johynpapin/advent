package main

import (
	"fmt"
)

func show(recipes []int, e0, e1 int) {
	var line string

	for i, score := range recipes {
		if i == e0 {
			line += fmt.Sprintf("(%d)", score)
		} else if i == e1 {
			line += fmt.Sprintf("[%d]", score)
		} else {
			line += fmt.Sprintf(" %d ", score)
		}
	}

	fmt.Println(line)
}

const MAGIC = 77201

//const MAGIC = 18

func main() {
	recipes := []int{3, 7}
	e0, e1 := 0, 1

	show(recipes, e0, e1)

	for len(recipes) < MAGIC+10 {
		newRecipe := recipes[e0] + recipes[e1]
		r0, r1 := newRecipe/10, newRecipe%10

		if newRecipe < 10 {
			recipes = append(recipes, r1)
		} else {
			recipes = append(recipes, r0, r1)
		}

		e0 = (e0 + 1 + recipes[e0]) % len(recipes)
		e1 = (e1 + 1 + recipes[e1]) % len(recipes)
	}

	fmt.Println("part1:", recipes[MAGIC:MAGIC+10])
}
