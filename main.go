package main

import (
	"flag"
	"fmt"

	"github.com/prestonchoate/aoc-2023/solutions"
	"github.com/prestonchoate/aoc-2023/solutions/day1"
	"github.com/prestonchoate/aoc-2023/solutions/day2"
	"github.com/prestonchoate/aoc-2023/solutions/day3"
	"github.com/prestonchoate/aoc-2023/solutions/day4"
)

func main() {
	// TODO: Add a global debug flag that prints info when true
	daySelect := flag.Int("day", 0, "Select which day to execute")
	partSelect := flag.Int("part", 0, "Select which part to execute")
	flag.Parse()
	if *daySelect <= 0 || *daySelect > 25 || *partSelect != 1 && *partSelect != 2 {
		handleInvalidInput(*daySelect, *partSelect)
		return
	}

	// TODO: How can I set this up to dynamically choose a solution?
	switch *daySelect {
	case 1:
		input := solutions.GetInputString("inputs/day1.txt")
		answer := 0
		if *partSelect == 1 {

			answer = day1.D1P1_Solve(input)
		}
		if *partSelect == 2 {
			answer = day1.D1P2_Solve(input)
		}

		printSolutionString(*daySelect, *partSelect, answer)
		break
	case 2:
		total := 0
		input := solutions.GetInputString("inputs/day2.txt")

		if *partSelect == 1 {
			total, _ = day2.Solve(input)
		}
		if *partSelect == 2 {
			_, total = day2.Solve(input)
		}

		printSolutionString(*daySelect, *partSelect, total)
		break
	case 3:
		total := *new(int64)
		total = 0
		input := solutions.GetInputString("inputs/day3.txt")
		if *partSelect == 1 {
			total = day3.Solve(input)
		}
		if *partSelect == 2 {
			total = day3.SolveP2(input)
		}

		printSolutionString(*daySelect, *partSelect, total)
	case 4:
		total := 0
		input := solutions.GetInputString("inputs/day4.txt")
		if *partSelect == 1 {
			total = day4.Solve(input)
		}
		if *partSelect == 2 {
			total = day4.SolveP2(input)
		}

		printSolutionString(*daySelect, *partSelect, total)
	default:
		handleInvalidInput(*daySelect, *partSelect)
		break
	}
}

func handleInvalidInput(daySelect int, partSelect int) {
	fmt.Printf("%v - %v is an invalid selection. Please try again\n", daySelect, partSelect)
}

func printSolutionString(daySelect int, partSelect int, answer any) {
	fmt.Printf("Day %v Part %v solution: %v\n", daySelect, partSelect, answer)
}
