package day1

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func D1P1_Solve(input string) int {
	return d1p1(input)
}

func D1P2_Solve(input string) int {
	return d1p2(input)
}

func d1p1_old(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	processed := 0

	for i := 0; i < len(lines); i++ {
		if len(lines[i]) <= 0 {
			continue
		}
		processed++
		// convert line into rune slice
		runes := []rune(lines[i])
		k := len(runes) - 1
		first, last := -1, -1
		for j := 0; j <= len(runes)-1; j++ {
			if first > -1 && last > -1 {
				break
			}

			if first == -1 {
				if unicode.IsDigit(runes[j]) {
					first = int(runes[j] - '0')
				}
			}

			if last == -1 {
				if unicode.IsDigit(runes[k]) {
					last = int(runes[k] - '0')
				}
			}
			k--
		}

		if first == -1 || last == -1 {
			fmt.Printf("No number found in %v\n", lines[i])
			continue
		}
		number := (first * 10) + last
		sum += number
	}

	return sum
}

// TODO: Lots of duplication in the two parts. It could be condensed to a single function/handler

func d1p1(input string) int {
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

	lines := strings.Split(input, "\n")
	sum := 0

	for i := 0; i < len(lines); i++ {
		if len(lines[i]) <= 0 {
			continue
		}

		_, left_num := findLeftNum(lines[i], &nums, false)
		_, right_num := findRightNum(lines[i], &nums, false)

		val := (left_num * 10) + right_num
		sum += val
	}

	return sum
}

func d1p2(input string) int {
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

	lines := strings.Split(input, "\n")
	sum := 0

	for i := 0; i < len(lines); i++ {
		if len(lines[i]) <= 0 {
			continue
		}

		_, left_num := findLeftNum(lines[i], &nums, true)
		_, right_num := findRightNum(lines[i], &nums, true)

		val := (left_num * 10) + right_num
		sum += val
	}

	return sum
}

func findLeftNum(line string, nums *map[string]int, checkStr bool) (int, int) {
	leftIndex := math.MaxInt
	val := -1
	for word, num := range *nums {
		if checkStr {

			idx := strings.Index(line, word)

			if idx > -1 && leftIndex > idx {
				leftIndex = idx
				val = num
			}
		}

		idx := strings.Index(line, fmt.Sprintf("%v", num))

		if idx > -1 && leftIndex > idx {
			leftIndex = idx
			val = num
		}
	}
	return leftIndex, val
}

func findRightNum(line string, nums *map[string]int, checkStr bool) (int, int) {
	rightIndex := -1
	val := -1
	for word, num := range *nums {
		if checkStr {

			idx := strings.LastIndex(line, word)

			if rightIndex < idx {
				rightIndex = idx
				val = num
			}
		}

		idx := strings.LastIndex(line, fmt.Sprintf("%v", num))

		if rightIndex < idx {
			rightIndex = idx
			val = num
		}
	}
	return rightIndex, val
}
