package main

import (
	"flag"
	"fmt"

	"github.com/prestonchoate/aoc-2023/solutions"
	"github.com/prestonchoate/aoc-2023/solutions/day1"
	"github.com/prestonchoate/aoc-2023/solutions/day2"
	"github.com/prestonchoate/aoc-2023/solutions/day3"
)

func main() {
	daySelect := flag.Int("day", 0, "Select which day to execute")
	partSelect := flag.Int("part", 0, "Select which part to execute")
	flag.Parse()

	switch *daySelect {
	case 1:
		handleDay1(*partSelect)
		break
	case 2:
		if *partSelect == 1 {
			total, _ := day2.Solve(solutions.GetInputString("inputs/day2-part1.txt"))
			fmt.Printf("Day Two Part One solution: %v\n", total)
		}
		if *partSelect == 2 {
			_, totalPower := day2.Solve(solutions.GetInputString("inputs/day2-part1.txt"))
			fmt.Printf("Day Two Part Two soltuion: %v\n", totalPower)
		}
		if *partSelect != 1 && *partSelect != 2 {
			handleInvalidInput(*daySelect, *partSelect)
		}
		break
	case 3:
		if *partSelect == 1 {
			total := day3.Solve(solutions.GetInputString("inputs/day3.txt"))
			fmt.Printf("Day Three Part One solution: %v\n", total)
		}
		if *partSelect == 2 {
			total := day3.SolveP2(solutions.GetInputString("inputs/day3.txt"))
			fmt.Printf("Day Three Part Two solution: %v\n", total)
		}
		if *partSelect != 1 && *partSelect != 2 {
			handleInvalidInput(*daySelect, *partSelect)
		}
	default:
		handleInvalidInput(*daySelect, *partSelect)
		break
	}
}

func handleInvalidInput(daySelect int, partSelect int) {
	fmt.Printf("%v - %v is an invalid selection. Please try again\n", daySelect, partSelect)
}

// TODO: This seems uneccesary

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
