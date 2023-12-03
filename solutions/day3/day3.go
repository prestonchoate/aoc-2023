package day3

import (
	"fmt"
	"strconv"
	"strings"
)

type Nums struct {
	xPos  int
	yPos  int
	val   int64
	width int
	found bool
}

type Schars struct {
	xPos int
	yPos int
	val  string
}

var numbers []Nums
var specials []Schars

func Solve(input string) int64 {

	numbers = []Nums{}
	specials = []Schars{}

	sum := *new(int64)

	setupData(input)

	for _, num := range numbers {
		// check each special char to see if is adjacent to number
		for _, schar := range specials {
			if schar.yPos == num.yPos || schar.yPos == num.yPos-1 || schar.yPos == num.yPos+1 {

				if schar.xPos == num.xPos-1 {
					fmt.Printf("Found part number: %v\n", num.val)
					num.found = true
					sum += num.val
					break
				}

				counter := 0
				for counter < num.width {
					if schar.xPos == num.xPos+counter || schar.xPos == num.xPos+counter+1 {
						fmt.Printf("Found part number: %v\n", num.val)
						num.found = true
						sum += num.val
						break
					}
					counter++
				}
			}
		}
	}
	return sum
}

func SolveP2(input string) int64 {
	sum := *new(int64)

	numbers = []Nums{}
	specials = []Schars{}

	setupData(input)

	for _, char := range specials {
		if char.val == "*" {
			found := []int64{}
			for _, number := range numbers {
				if number.yPos == char.yPos || number.yPos == char.yPos-1 || number.yPos == char.yPos+1 {

					counter := 0
					for counter < number.width {
						if number.xPos+counter == char.xPos || number.xPos+counter == char.xPos-1 || number.xPos == char.xPos+1 {
							//fmt.Printf("%v is a potential gear part number\n", number.val)
							found = append(found, number.val)
							counter++
							break
						}
						counter++
					}
				}
				if len(found) >= 2 {
					//fmt.Println("We found some gears!")
					sum += found[0] * found[1]
					break
				}

			}
		}
	}

	return sum
}

func setupData(input string) {
	lines := strings.Split(input, "\n")

	for i := 0; i < len(lines); i++ {
		chars := strings.Split(lines[i], "")

		for j := 0; j < len(lines[i]); j++ {
			char := lines[i][j]

			val, err := strconv.ParseInt(string(char), 10, 64)
			if err == nil {
				// fmt.Printf("Creating a new number that starts with %v at position %v:%v\n", val, j, i)
				number := Nums{
					xPos:  j,
					yPos:  i,
					val:   val,
					width: 1,
				}
				numLen := findEndOfNumber(chars, j) - j
				// fmt.Printf("Number has a length of %v\n", numLen)
				count := 1
				for count <= numLen {
					if j+count >= len(chars) {
						count++
						continue
					}
					digit, err := strconv.ParseInt(chars[j+count], 10, 64)
					if err == nil {
						// fmt.Printf("Appending digit: %v to number: %v\n", digit, number.val)
						number.val = (number.val * 10) + digit
						number.width++
					}
					count++
				}

				numbers = append(numbers, number)
				j += numLen
				continue
			}

			if chars[j] != "." {
				schar := Schars{
					val:  chars[j],
					xPos: j,
					yPos: i,
				}
				specials = append(specials, schar)
			}
		}
	}

	// fmt.Printf("Found Numbers: %v\n", numbers)
	// fmt.Printf("Found special chars: %v\n", specials)
}

func findEndOfNumber(line []string, start int) int {
	start++
	for ; start < len(line); start++ {
		_, err := strconv.Atoi(line[start])
		if err == nil {
		} else {
			return start - 1
		}
	}
	return start
}
