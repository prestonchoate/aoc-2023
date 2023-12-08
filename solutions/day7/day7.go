package day7

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var cardvals map[string]int = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

type handType int

func (h handType) String() string {
	switch h {
	case HighCard:
		return "High Card"
	case Pair:
		return "Pair"
	case TwoPair:
		return "Two Pair"
	case ThreeOfAKind:
		return "Three of a Kind"
	case FullHouse:
		return "Full House"
	case FourOfAKind:
		return "Four of a Kind"
	case FiveOfAKind:
		return "Five of a Kind"
	}
	return ""
}

const (
	HighCard     handType = 0
	Pair         handType = 1
	TwoPair      handType = 2
	ThreeOfAKind handType = 3
	FullHouse    handType = 4
	FourOfAKind  handType = 5
	FiveOfAKind  handType = 6
)

type Hand struct {
	cards    map[string]int
	bet      int
	lowCard  string
	highCard string
	HandType handType
	cardStr  string
}

func createHand() Hand {
	hand := Hand{
		cards: make(map[string]int),
	}

	return hand
}

func (h *Hand) addCard(card string) *Hand {
	if count, ok := h.cards[card]; ok {
		h.cards[card] = count + 1
	} else {
		h.cards[card] = 1
	}

	if h.lowCard == "" {
		h.lowCard = card
	}

	if h.lowCard != card {
		lcVal := cardvals[h.lowCard]
		cVal := cardvals[card]
		if cVal < lcVal {
			h.lowCard = card
		}
	}
	if h.highCard != card {
		hcVal := cardvals[h.highCard]
		cVal := cardvals[card]
		if cVal > hcVal {
			h.highCard = card
		}
	}

	return h
}

// I hate everything about this function. There is probably some clever recursive way to do this instead
func (h *Hand) classifyHand(wilds []string) *Hand {
	count := len(h.cards)
	wildCount := 0
	hasWilds := false
	for _, wild := range wilds {
		count, ok := h.cards[wild]
		if ok {
			hasWilds = true
			wildCount += count
		}
	}

	switch count {
	case 1:
		h.HandType = handType(FiveOfAKind)
		break
	case 2:
		// Could be 4 of a kind or full house
		for _, v := range h.cards {
			if v == 1 {
				h.HandType = FourOfAKind
				break
			}
			if v == 2 {
				h.HandType = FullHouse
				break
			}
		}
		break
	case 3:
		// Could be 3 of a kind or two pair
		for char, v := range h.cards {
			if v == 3 {
				if slices.Contains(wilds, char) {
					// Three wilds
					h.HandType = FourOfAKind
					hasWilds = false
					break
				}
				if wildCount == 2 {
					h.HandType = FiveOfAKind
					hasWilds = false
					break
				}
				if wildCount == 1 {
					h.HandType = FourOfAKind
					hasWilds = false
					break
				}
				h.HandType = ThreeOfAKind
				break
			}
		}
		if h.HandType == HighCard {
			if wildCount == 2 {
				h.HandType = FourOfAKind
				hasWilds = false
				break
			}
			if wildCount == 1 {
				h.HandType = FullHouse
				hasWilds = false
				break
			}
			h.HandType = TwoPair
		}
		break
	case 4:
		for _, v := range h.cards {
			if v == 2 {
				if hasWilds {
					h.HandType = ThreeOfAKind
					hasWilds = false
					break
				}
			}
		}
		if h.HandType == HighCard {
			h.HandType = Pair
		}
		break
	default:
		h.HandType = HighCard
	}

	if hasWilds {
		h.HandType = min(h.HandType+handType(wildCount), FiveOfAKind)
	}
	return h
}

func setupHands(input string, wilds []string) *[]Hand {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	hands := []Hand{}

	if wilds != nil {
		for _, wild := range wilds {
			cardvals[wild] = 0
		}
	}

	for _, line := range lines {
		if len(line) <= 0 {
			continue
		}
		line = strings.TrimSpace(line)

		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			fmt.Printf("%v got split into %v parts\n", line, len(parts))
			return nil
		}

		cardStr := parts[0]
		betStr := parts[1]

		bet, err := strconv.Atoi(betStr)
		if err != nil {
			fmt.Printf("Could not parse %v into number\n", betStr)
			return nil
		}

		hand := createHand()
		hand.bet = bet
		hand.cardStr = cardStr
		for _, card := range cardStr {
			hand.addCard(string(card))
		}

		hand.classifyHand(wilds)
		hands = append(hands, hand)
	}

	return &hands
}

func SolveP1(input string) int {

	hands := setupHands(input, nil)

	sorted := rankHands(*hands)

	total := 0
	for idx, hand := range sorted {
		total += hand.bet * (idx + 1)
	}
	return total
}

func SolveP2(input string) int {
	wilds := []string{
		"J",
	}
	hands := setupHands(input, wilds)

	sorted := rankHands(*hands)
	total := 0

	for idx, hand := range sorted {
		total += hand.bet * (idx + 1)
		//fmt.Printf("Sorted Hand#%v: %+v\n", idx+1, hand)
	}
	return total
}

func rankHands(hands []Hand) []Hand {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].HandType == hands[j].HandType {
			// Figure out the first high card
			for k := 0; k < len(hands[i].cardStr); k++ {
				iCard := string(hands[i].cardStr[k])
				jCard := string(hands[j].cardStr[k])
				if iCard == jCard {
					continue
				}
				iVal := cardvals[iCard]
				jVal := cardvals[jCard]
				return iVal < jVal
			}
		}
		return hands[i].HandType < hands[j].HandType
	})

	return hands
}
