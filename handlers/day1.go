package handlers

import (
	"fmt"

	"github.com/prestonchoate/aoc-2023/solutions"
)

func HandleDay1(part int) {
	switch part {
	case 1:
		answer := solutions.D1P1_Solve("inputs/d1p1-in.txt")
		fmt.Printf("Day One Part One solution: %v\n", answer)
		break
	case 2:
		answer := solutions.D1P2_Solve("inputs/d1p1-in.txt")
		fmt.Printf("Day One Part Two solution: %v\n", answer)
		break
	default:
		fmt.Printf("%v is not a valid part for this day\n", part)
		break
	}
}
