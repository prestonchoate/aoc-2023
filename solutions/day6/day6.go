package day6

import (
	"fmt"
	"strconv"
	"strings"
)

type Race struct {
	raceTime int
	distance int
}

func SolveP1(input string) int {
	input = strings.TrimSpace(input)
	parts := strings.Split(input, "\n")
	totalWins := 0
	races := parseRaces(parts)
	if races == nil {
		return -1
	}

	for _, race := range *races {
		//fmt.Printf("Found race: Time - %v, Distance - %v\n", race.raceTime, race.distance)
		wins := calculateWinCount(race.raceTime, race.distance)
		if totalWins == 0 {
			totalWins += wins
		} else {
			totalWins *= wins
		}
	}
	return totalWins
}

func SolveP2(input string) int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	timeStrs := strings.Split(lines[0], ":")
	times := strings.TrimSpace(timeStrs[1])
	times = strings.Join(strings.Fields(times), "")
	raceTime, err := strconv.Atoi(times)
	if err != nil {
		fmt.Printf("Could not convert %v to an integer\n", times)
		return -1
	}

	distStrs := strings.Split(lines[1], ":")
	dists := strings.TrimSpace(distStrs[1])
	dists = strings.Join(strings.Fields(dists), "")
	raceDist, err := strconv.Atoi(dists)
	if err != nil {
		fmt.Printf("Could not convert %v to an integer\n", dists)
		return -1
	}

	return calculateWinCount(raceTime, raceDist)
}

func parseRaces(parts []string) *[]Race {
	races := []Race{}

	if len(parts) != 2 {
		for idx, part := range parts {
			fmt.Printf("%v - \"%v\"\n", idx, part)
		}
		fmt.Printf("Something went wrong parsing %+v\n", parts)
		return nil
	}

	timeStrs := strings.Split(parts[0], ":")
	times := strings.TrimSpace(timeStrs[1])
	times = strings.Join(strings.Fields(times), " ")
	timeVals := strings.Split(times, " ")
	distanceStrs := strings.Split(parts[1], ":")
	distances := strings.TrimSpace(distanceStrs[1])
	distances = strings.Join(strings.Fields(distances), " ")
	distVals := strings.Split(distances, " ")

	if len(timeVals) != len(distVals) {
		fmt.Printf("Invalid input: \ntimes: %+v len: %v\n dists: %+v len: %v\n", timeVals, len(timeVals), distVals, len(distVals))
		return nil
	}

	for i := 0; i < len(timeVals); i++ {
		raceTime, err := strconv.Atoi(timeVals[i])
		if err != nil {
			fmt.Printf("Failed to parse race duration from %v\n", timeVals[i])
			return nil
		}

		distance, err := strconv.Atoi(distVals[i])
		if err != nil {
			fmt.Printf("Failed to parse race distance from %v\n", distVals[i])
		}

		race := Race{
			raceTime: raceTime,
			distance: distance,
		}
		races = append(races, race)
	}

	return &races
}

func calculateWinCount(raceTime int, distance int) int {
	wins := 0
	counter := 0
	raceDist := 0
	for counter <= raceTime {
		raceDist = counter * (raceTime - counter)
		//fmt.Printf("Button held for %v ms, total distance %v\n", counter, raceDist)
		if raceDist > distance {
			wins++
		}
		counter++
	}

	return wins
}
