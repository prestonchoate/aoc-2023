package day2

import (
	"testing"
)

func TestD2P1(t *testing.T) {
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	expect := 8

	got, _ := Solve(input)
	if expect != got {
		t.Fatalf("Expected %v, got %v\n", expect, got)
	}
}

func TestD2P2(t *testing.T) {
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	expectSum := 8
	expectPower := 2286

	gotSum, gotPower := Solve(input)
	if expectSum != gotSum {
		t.Fatalf("Expected valid game sum: %v, got game sum: %v\n", expectSum, gotSum)
	}
	if expectPower != gotPower {
		t.Fatalf("Expected valid total game power: %v, got valid total game power: %v\n", expectPower, gotPower)
	}
}
