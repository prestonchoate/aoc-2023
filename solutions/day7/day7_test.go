package day7

import (
	"testing"

	"github.com/prestonchoate/aoc-2023/solutions"
)

func TestPart1(t *testing.T) {
	input := solutions.GetInputString("../../inputs/day7-sample.txt")

	expected := 6440

	got := SolveP1(input)

	if got != expected {
		t.Fatalf("Expected %v, Got %v\n", expected, got)
	}
}

func TestPart2(t *testing.T) {
	input := solutions.GetInputString("../../inputs/day7-sample.txt")

	expected := 5905

	got := SolveP2(input)

	if got != expected {
		t.Fatalf("Expected %v, Got %v\n", expected, got)
	}

}

func TestClassifyHand(t *testing.T) {
	hand := createHand()
	hand.cardStr = "AAAJ5"
	hand.addCard("A").addCard("A").addCard("A").addCard("J").addCard("5")
	hand.bet = 10

	cardvals["J"] = 0

	wilds := []string{"J"}
	hand.classifyHand(wilds)

	if hand.HandType != FourOfAKind {
		t.Fatalf("Expected hand: %v to be four of a kind, instead got %v\n", hand.cardStr, hand.HandType)
	}

	hand = createHand()
	hand.cardStr = "A32JJ"
	hand.addCard("A").addCard("3").addCard("2").addCard("J").addCard("J")
	hand.classifyHand(wilds)

	if hand.HandType != ThreeOfAKind {
		t.Fatalf("Expected hand: %v to be three of a kind instead got %v\n", hand.cardStr, hand.HandType)
	}

	hand = createHand()
	hand.cardStr = "J2279"
	hand.addCard("J").addCard("2").addCard("2").addCard("7").addCard("9")
	hand.classifyHand(wilds)

	if hand.HandType != ThreeOfAKind {
		t.Fatalf("Expected hand %v to be three of a kind, instead got %v\n", hand.cardStr, hand.HandType)
	}

	hand = createHand()
	hand.cardStr = "J3377"
	hand.addCard("J").addCard("3").addCard("3").addCard("7").addCard("7")
	hand.classifyHand(wilds)

	if hand.HandType != FullHouse {
		t.Fatalf("Expected hand %v to be full house, instead got %v\n", hand.cardStr, hand.HandType)
	}
}
