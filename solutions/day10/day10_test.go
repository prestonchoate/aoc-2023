package day10

import (
	"testing"

	"github.com/prestonchoate/aoc-2023/solutions"
)

func TestPart1(t *testing.T) {
	input := solutions.GetInputString("../../inputs/day10-sample.txt")

	expected := 2

	got := SolveP1(input)

	if got != expected {
		t.Fatalf("Expected %v, Got %v\n", expected, got)
	}
}

func TestPart2(t *testing.T) {
	input := solutions.GetInputString("../../inputs/day10-sample2.txt")

	expected := 10

	got := SolveP2(input)

	if got != expected {
		t.Fatalf("Expected %v, Got %v\n", expected, got)
	}
}
