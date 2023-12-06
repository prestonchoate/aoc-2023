package day5

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type SeedRange struct {
	seedStart int
	seedEnd   int
}

func SolveP1(input string) int {
	chunks := strings.Split(input, "\n\n")
	seeds := createSeeds(chunks[0])
	if seeds == nil {
		fmt.Println("Operation failed")
		return -1
	}
	fmt.Printf("Seeds: %+v\n", seeds)

	for _, chunk := range chunks[1:] {
		seeds = processMap(chunk, seeds)
	}

	minVal := math.MaxInt
	for _, seed := range seeds {
		if seed < minVal {
			minVal = seed
		}
	}
	return minVal
}

func SolveP2(input string) int {
	chunks := strings.Split(input, "\n\n")

	seeds := createSeedsFromRanges(chunks[0])

	if seeds == nil {
		fmt.Println("Operation failed")
		return -1
	}
	fmt.Printf("Seeds: %+v\n", seeds)

	// Loop over every chunk to find transformations
	for _, chunk := range chunks[1:] {
		lines := strings.Split(chunk, "\n")
		for _, line := range lines[1:] {
			if len(line) <= 0 {
				continue
			}
			destNum, srcNum, rng := getTransforms(line)

			fmt.Printf("DestStart: %v, SrcStart: %v, Range: %v\n", destNum, srcNum, rng)

			for _, seed := range seeds {

				// TODO: None of these use cases are correct fix them

				// Handle case where begining is transformed but not all the way
				if seed.seedStart <= srcNum && seed.seedEnd > srcNum+rng {
					fmt.Printf("handling begining transform only\n")
					fmt.Printf("Need to create new seed staring at %v and ending at %v\n", srcNum+rng, seed.seedEnd)
					fmt.Printf("Need to change seed start num (%v) to %v\n", seed.seedStart, destNum+seed.seedStart-srcNum)
					fmt.Printf("Need to change seed end num (%v) to %v\n", seed.seedEnd, srcNum+rng)
				}

				// Handle case where end is transformed but not from begining
				if seed.seedStart < srcNum && seed.seedEnd >= srcNum+rng {
					fmt.Printf("handling end transform only\n")
					fmt.Printf("Need to create new seed starting at %v and ending at %v\n", seed.seedStart, srcNum)
					fmt.Printf("Need to change seed start num (%v) to %v\n", seed.seedStart, srcNum)
					fmt.Printf("Need to change seed end num (%v) to %v\n", seed.seedEnd, srcNum+rng)
				}

				// Handle case where middle is transformed but not begining or end

				// Handle case where whole seed is transformed
				if seed.seedStart >= srcNum && seed.seedEnd <= srcNum+rng {
					fmt.Printf("handling full transform\n")
					fmt.Printf("Need to change seed start num (%v) to %v\n", seed.seedStart, destNum+seed.seedStart-srcNum)
					fmt.Printf("Need to change seed end num (%v) to %v\n", seed.seedEnd, destNum+seed.seedEnd-srcNum)
				}
			}
		}
	}
	// Loop over seeds and if the range in the seed is within the transform range modify it
	// Handle situation where seed start range is less than the transform start
	// Handle situation where transform end range is less than seed end range

	/*
		for _, chunk := range chunks[1:] {
			seeds = processMap(chunk, seeds)
		}

		minVal := math.MaxInt
		for _, seed := range seeds {
			if seed < minVal {
				minVal = seed
			}
		}
	*/
	return 0
}

func createSeeds(input string) []int {
	seeds := []int{}

	parts := strings.Split(input, ": ")
	if len(parts) != 2 {
		fmt.Println("Failed to parse seed IDs")
		return nil
	}
	ids := strings.Split(parts[1], " ")
	for _, id := range ids {
		num, err := strconv.Atoi(id)
		if err != nil {
			fmt.Printf("Could not parse: %v Skipping\n", id)
			continue
		}
		seeds = append(seeds, num)

	}
	return seeds
}

func createSeedsFromRanges(input string) []SeedRange {
	seedRanges := []SeedRange{}
	parts := strings.Split(input, ": ")
	if len(parts) != 2 {
		fmt.Println("Failed to parse seed IDs")
		return nil
	}

	ids := strings.Split(parts[1], " ")
	for idx, id := range ids {
		num, err := strconv.Atoi(id)
		if err != nil {
			fmt.Printf("Could not parse: %v Skipping\n", id)
			continue
		}
		if idx%2 == 0 {
			rng, err := strconv.Atoi(ids[idx+1])
			if err != nil {
				fmt.Println("Something went wrong")
				return nil
			}
			seedRange := SeedRange{
				seedStart: num,
				seedEnd:   num + rng,
			}
			seedRanges = append(seedRanges, seedRange)
		}
	}
	return seedRanges
}

func processMap(input string, seeds []int) []int {
	lines := strings.Split(input, "\n")
	newSeeds := make([]int, len(seeds))
	for _, line := range lines[1:] {
		if len(line) <= 0 {
			continue
		}
		//fmt.Printf("Found line: %v\n", line)
		line = strings.TrimSpace(line)
		strs := strings.Split(line, " ")

		destNum, err := strconv.Atoi(strs[0])
		if err != nil {
			fmt.Printf("Failed to parse number: %v\n", strs[0])
			return nil
		}

		srcNum, err := strconv.Atoi(strs[1])
		if err != nil {
			fmt.Printf("Failed to parse number: %v\n", strs[1])
			return nil
		}

		rng, err := strconv.Atoi(strs[2])
		if err != nil {
			fmt.Printf("Failed to parse number: %v\n", strs[2])
			return nil
		}
		for idx, seed := range seeds {
			if seed == -1 {
				continue
			}
			newSeed := calculateNewSeed(srcNum, destNum, rng, seed)
			if newSeed != seed {
				newSeeds[idx] = newSeed
				seeds[idx] = -1
			}
		}
	}
	for idx, seed := range seeds {
		if seed != -1 {
			newSeeds[idx] = seed
		}
	}

	fmt.Printf("Updated seeds: %+v\n", newSeeds)
	return newSeeds
}

func calculateNewSeed(srcNum int, destNum int, rng int, seed int) int {
	if seed >= srcNum && seed < srcNum+rng {
		newSeed := destNum + (seed - srcNum)
		//fmt.Printf("Updating seed: %v to %v\n", seed, destNum+(seed-srcNum))
		return newSeed
	}
	return seed
}

func getTransforms(line string) (destStart int, srcStart int, rng int) {

	parts := strings.Split(line, " ")

	if len(parts) < 2 {
		fmt.Printf("Failed to get transform numbers")
		return -1, -1, -1
	}

	destStart, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Printf("Failed processing line. Error: %v\n", err.Error())
		return -1, -1, -1
	}

	srcStart, err = strconv.Atoi(parts[1])
	if err != nil {
		fmt.Printf("Failed processing line. Error: %v\n", err.Error())
		return -1, -1, -1
	}

	rng, err = strconv.Atoi(parts[2])
	if err != nil {
		fmt.Printf("Failed processing line. Error: %v\n", err.Error())
		return -1, -1, -1
	}

	return destStart, srcStart, rng
}
