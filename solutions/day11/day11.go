package day11

import (
	"math"
	"strings"
)

type galaxy struct {
	xPos int
	yPos int
	id   int
}

func SolveP1(input string) int {
	lines := strings.Split(input, "\n")
	emptyRows, emptyCols, galaxies := setup(lines)

	extraScale := 1

	return getDistances(galaxies, extraScale, emptyRows, emptyCols)
}

func SolveP2(input string) int {
	lines := strings.Split(input, "\n")
	emptyRows, emptyCols, galaxies := setup(lines)

	extraScale := int(math.Pow(10, 6) - 1)

	return getDistances(galaxies, extraScale, emptyRows, emptyCols)
}

func setup(lines []string) (emptyRows map[int]bool, emptyCols map[int]bool, galaxies []galaxy) {
	emptyRows = map[int]bool{}
	emptyCols = map[int]bool{}
	galaxies = []galaxy{}

	for row, line := range lines {
		rowEmpty := true
		for col, char := range strings.Split(line, "") {
			if _, ok := emptyCols[col]; !ok {
				emptyCols[col] = true
			}
			if char != "." {
				emptyCols[col] = false
				rowEmpty = false
			}
			if char == "#" {
				galaxy := galaxy{
					xPos: col,
					yPos: row,
					id:   len(galaxies) + 1,
				}

				galaxies = append(galaxies, galaxy)
			}
		}
		emptyRows[row] = rowEmpty
	}

	return emptyRows, emptyCols, galaxies

}

func getDistances(galaxies []galaxy, extraScale int, emptyRows map[int]bool, emptyCols map[int]bool) int {
	total := 0
	for i := 0; i < len(galaxies); i++ {
		gTotal := 0
		curGal := galaxies[i]
		//fmt.Printf("Checking galaxy: %+v\n\n", curGal)
		for j := i + 1; j < len(galaxies); j++ {
			nextGal := galaxies[j]
			//fmt.Printf("Against galaxy: %+v\n", nextGal)
			gTotal = calcDistance(curGal, nextGal, emptyRows, emptyCols, extraScale)
			//fmt.Printf("Distance: %v\n", gTotal)
			total += gTotal
		}
		//fmt.Println()
	}
	return total
}

func calcDistance(start galaxy, target galaxy, emptyRows map[int]bool, emptyCols map[int]bool, scale int) int {
	distance := 0
	distance += abs(target.xPos - start.xPos)
	distance += abs(target.yPos - start.yPos)

	minX := min(start.xPos, target.xPos)
	maxX := max(start.xPos, target.xPos)
	minY := min(start.yPos, target.yPos)
	maxY := max(start.yPos, target.yPos)

	for idx, empty := range emptyCols {
		if minX <= idx && idx <= maxX && empty {
			distance += scale
		}
	}

	for idx, empty := range emptyRows {
		if minY <= idx && idx <= maxY && empty {
			distance += scale
		}
	}

	return distance
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
