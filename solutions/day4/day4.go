package day4

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	id           int
	winningNums  []int
	revealedNums []int
	score        int
	matches      int
}

func Solve(input string) int {
	totalPoints := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		cardId := parseCardId(line)
		if cardId == -1 {
			continue
		}
		fmt.Printf("Found Card ID: %v\n", cardId)

		winningNums, revealedNums := parseNumbers(line)
		fmt.Printf("Winning Numbers: %v\n", winningNums)
		fmt.Printf("Revealed Numbers: %v\n", revealedNums)

		cardPoints, _ := calculateCardPoints(winningNums, revealedNums)
		fmt.Printf("Total points for card: %v is %v\n", cardId, cardPoints)
		totalPoints += cardPoints
	}

	return totalPoints
}

func SolveP2(input string) int {
	totalCards := 0
	cards := []Card{}

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		cardId := parseCardId(line)
		winningNums, revealedNums := parseNumbers(line)

		cardPoints, matches := calculateCardPoints(winningNums, revealedNums)

		card := Card{
			id:           cardId,
			winningNums:  winningNums,
			revealedNums: revealedNums,
			score:        cardPoints,
			matches:      matches,
		}

		cards = append(cards, card)
	}

	//fmt.Printf("Cards: %+v\n", cards)

	totalCards = calculateCardCopies(cards)

	return totalCards
}

func parseCardId(line string) int {
	parts := strings.Split(line, ":")
	if len(parts) < 2 {
		fmt.Printf("Could not parse the card number from supplied input: %v\n", line)
		return -1
	}

	stripped := strings.Join(strings.Fields(parts[0]), " ")
	idStr := strings.Split(stripped, " ")

	if len(idStr) != 2 {
		fmt.Printf("Could not parse the card number from supplied input: %v\n", parts[0])
		return -1
	}

	str := strings.TrimSpace(idStr[1])
	cardId, err := strconv.Atoi(str)

	if err != nil {
		fmt.Printf("Could not parse the card number from the supplied input: %v\n", str)
		return -1
	}

	return cardId
}

// Returns a slice of winning numbers, and a slice of revealed numbers
func parseNumbers(line string) ([]int, []int) {
	winningNums := []int{}
	revealedNums := []int{}

	parts := strings.Split(line, ":")
	if len(parts) < 2 {
		fmt.Printf("Could not parse numbers from the supplied input: %v\n", line)
		return nil, nil
	}

	numStrs := strings.Split(parts[1], " | ")
	if len(numStrs) < 2 {
		fmt.Printf("Could not parse numbers from the supplied input: %v\n", line)
		return nil, nil
	}

	numStr := strings.TrimSpace(numStrs[0])
	numbers := strings.Split(numStr, " ")

	for idx, str := range numbers {
		if len(str) == 0 {
			continue
		}
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Failed to parse number at index %v: %v\n", idx, numbers)
			continue
		}
		winningNums = append(winningNums, num)
	}

	numStr = strings.TrimSpace(numStrs[1])
	numbers = strings.Split(numStr, " ")

	for idx, str := range numbers {
		if len(str) == 0 {
			continue
		}
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Failed to parse number it index %v: %v\n", idx, numbers)
			continue
		}
		revealedNums = append(revealedNums, num)
	}

	return winningNums, revealedNums
}

func calculateCardPoints(winningNums []int, revealedNums []int) (int, int) {
	points := 0
	matches := 0

	for _, num := range revealedNums {
		if slices.Contains(winningNums, num) {
			if points == 0 {
				matches++
				points++
				continue
			}
			matches++
			points *= 2
		}
	}

	return points, matches
}

func calculateCardCopies(cards []Card) int {
	total := 0
	copies := make(map[int]int)

	for _, card := range cards {
		extraCount := copies[card.id]
		fmt.Printf("Running game ID%v X%v with %v matches\n", card.id, extraCount+1, card.matches)

		for extraCount >= 0 {
			for i := 1; i <= card.matches; i++ {
				if card.id+i <= len(cards) {
					copies[card.id+i]++
				}
			}
			total++
			extraCount--
		}
	}

	fmt.Printf("Copies: %+v\n", copies)

	return total
}
