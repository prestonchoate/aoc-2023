package day1

import (
	"testing"
)

func TestD1P1(t *testing.T) {
	input := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`

	want := 142

	res := d1p1(input)

	if want != res {
		t.Fatalf("Expected: %v \t Got: %v\n", want, res)
	}
}

func TestD1P2(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

	want := 281
	res := d1p2(input)
	if want != res {
		t.Fatalf("Expected %v \t Got: %v\n", want, res)
	}
}

func TestD1P2Edge(t *testing.T) {
	input := "twolxzdhfourqjeightfour55zjvconeightnf"

	nums := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	want := 8
	_, res := findRightNum(input, &nums, true)

	if want != res {
		t.Fatalf("Expected %v \t Got: %v\n", want, res)
	}
}
