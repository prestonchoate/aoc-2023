package day9

import (
	"testing"

	"github.com/prestonchoate/aoc-2023/solutions"
)

func TestPart1(t *testing.T) {
	input := solutions.GetInputString("../../inputs/day9-sample.txt")

	expected := 114
	got := SolvePart1(input)

	if got != expected {
		t.Fatalf("Expected %v, Got %v\n", expected, got)
	}
}

func TestPart2(t *testing.T) {

	input := solutions.GetInputString("../../inputs/day9-sample.txt")

	expected := 2
	got := SolvePart2(input)

	if got != expected {
		t.Fatalf("Expected %v, Got %v\n", expected, got)
	}
}

func TestGetNextNum(t *testing.T) {
	nums := []int{
		8, 4, 0, -4, -8, -12, -16, -20, -24, -28, -32, -36, -40, -44, -48, -52, -56, -60, -64, -68, -72,
	}

	next := getNextNum(nums)
	expected := -76

	if next != expected {
		t.Fatalf("Expected %v, got %v\n", expected, next)
	}
}
