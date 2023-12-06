package day6

import "testing"

func TestPart1(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200`

	expected := 288
	got := SolveP1(input)

	if got != expected {
		t.Fatalf("Expected %v, but got %v\n", expected, got)
	}
}
