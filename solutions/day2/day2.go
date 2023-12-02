package day2

import (
	"fmt"
	"strconv"
	"strings"
)

var gameConfig = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Solve(input string) (int, int) {
	lines := strings.Split(input, "\n")
	sum := 0
	totalPower := 0

	for _, line := range lines {
		if len(line) <= 0 {
			continue
		}

		fmt.Printf("Parsing line: %v\n", line)
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Printf("Invalid input\n")
			return -1, -1
		}

		gameNum := getGameNum(parts[0])
		fmt.Printf("Game Number: %v\n", gameNum)

		valid, power := validateRoundData(parts[1])
		if valid {
			sum += gameNum
		} else {
		}
		totalPower += power
	}

	return sum, totalPower
}

func getGameNum(input string) int {
	parts := strings.Split(input, " ")
	if len(parts) < 2 || parts[0] != "Game" {
		fmt.Printf("Could not parse game number from input\n")
		return -1
	}

	gameNum, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Printf("Failed to parse game number from input. Error: %v", err.Error())
	}
	return gameNum
}

func validateRoundData(input string) (bool, int) {
	parts := strings.Split(input, ";")
	mins := map[string]int{
		"red":   0,
		"blue":  0,
		"green": 0,
	}
	valid := true

	for _, part := range parts {
		pulls := strings.Split(part, ",")
		for _, pull := range pulls {
			pull = strings.TrimSpace(pull)
			vals := strings.Split(pull, " ")
			if len(vals) < 2 {
				fmt.Printf("Invalid input. Tried to parse: %v\n", pull)
				return false, -1
			}
			color := vals[1]
			validCount := gameConfig[color]
			count, err := strconv.Atoi(vals[0])
			if err != nil {
				fmt.Printf("Could not parse pull count: %v", err.Error())
				return false, -1
			}
			if count > validCount {
				valid = false
			}
			minCount := mins[color]
			if count > minCount {
				mins[color] = count
			}
		}
	}

	power := calculateRoundPower(&mins)
	return valid, power
}

func calculateRoundPower(mins *map[string]int) int {
	total := 1
	for _, count := range *mins {
		total *= count
	}

	return total
}
