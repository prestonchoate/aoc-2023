package main

import (
	"flag"
	"fmt"

	"github.com/prestonchoate/aoc-2023/solutions/day1"
)

func main() {
	daySelect := flag.Int("day", 0, "Select which day to execute")
	partSelect := flag.Int("part", 0, "Select which part to execute")
	flag.Parse()

	switch *daySelect {
	case 1:
		handleDay1(*partSelect)
		break
	default:
		fmt.Printf("%v is an invalid day. Please try again\n", *daySelect)
		break
	}
}

func handleDay1(part int) {
	switch part {
	case 1:
		answer := day1.D1P1_Solve("inputs/d1p1-in.txt")
		fmt.Printf("Day One Part One solution: %v\n", answer)
		break
	case 2:
		answer := day1.D1P2_Solve("inputs/d1p1-in.txt")
		fmt.Printf("Day One Part Two solution: %v\n", answer)
		break
	default:
		fmt.Printf("%v is not a valid part for this day\n", part)
		break
	}
}
