package day10

import (
	"fmt"
	"strings"

	"os"
)

type position struct {
	xpos int
	ypos int
}

func push(q []position, v position) []position {
	return append(q, v)
}

func pop(q []position) ([]position, position) {
	v := q[0]
	return q[1:], v
}

func SolveP1(input string) int {
	lines := strings.Split(input, "\n")
	queue := []position{}
	seen := []position{}
	pipeMap := [][]string{}
	curPos := position{}
	for y, line := range lines {
		chars := []string{}
		for x, char := range strings.Split(line, "") {
			chars = append(chars, char)
			if char == "S" {
				curPos.xpos = x
				curPos.ypos = y
				seen = append(seen, curPos)
				queue = push(queue, curPos)
			}
		}
		pipeMap = append(pipeMap, chars)
	}

	for len(queue) > 0 {
		pos := position{}
		queue, pos = pop(queue)
		char := pipeMap[pos.ypos][pos.xpos]

		// check if there is a valid pipe to move through above pos if so add to queue
		if pos.ypos > 0 && (char == "S" || char == "|" || char == "J" || char == "L") {
			moveChar := pipeMap[pos.ypos-1][pos.xpos]
			if !haveSeen(seen, pos.xpos, pos.ypos-1) && (moveChar == "|" || moveChar == "7" || moveChar == "F") {
				newPos := position{xpos: pos.xpos, ypos: pos.ypos - 1}
				seen = append(seen, newPos)
				queue = push(queue, newPos)
			}
		}

		// check if there is a valid pipe to move through right of pos if so add to queue
		if pos.xpos < len(pipeMap[pos.ypos])-1 && (char == "S" || char == "-" || char == "L" || char == "F") {
			moveChar := pipeMap[pos.ypos][pos.xpos+1]
			if !haveSeen(seen, pos.xpos+1, pos.ypos) && (moveChar == "-" || moveChar == "7" || moveChar == "J") {
				newPos := position{xpos: pos.xpos + 1, ypos: pos.ypos}
				seen = append(seen, newPos)
				queue = append(queue, newPos)
			}
		}

		// check if there is a valid pipe to move through left of pos if so add to queue
		if pos.xpos > 0 && (char == "S" || char == "-" || char == "7" || char == "J") {
			moveChar := pipeMap[pos.ypos][pos.xpos-1]
			if !haveSeen(seen, pos.xpos-1, pos.ypos) && (moveChar == "-" || moveChar == "L" || moveChar == "F") {
				newPos := position{xpos: pos.xpos - 1, ypos: pos.ypos}
				seen = append(seen, newPos)
				queue = push(queue, newPos)
			}
		}

		// check if there is a valid pipe to move through below pos if so add to queue
		if pos.ypos < len(pipeMap) && (char == "S" || char == "|" || char == "7" || char == "F") {
			moveChar := pipeMap[pos.ypos+1][pos.xpos]
			if !haveSeen(seen, pos.xpos, pos.ypos+1) && (moveChar == "S" || moveChar == "|" || moveChar == "J" || moveChar == "L") {
				newPos := position{xpos: pos.xpos, ypos: pos.ypos + 1}
				seen = append(seen, newPos)
				queue = push(queue, newPos)
			}
		}
	}

	fmt.Fprintf(os.Stdout, "Seen nodes count: %v\n", len(seen))

	return len(seen) / 2
}

func SolveP2(input string) int {
	lines := strings.Split(input, "\n")
	queue := []position{}
	seen := []position{}
	pipeMap := [][]string{}
	sPos := position{}
	possibleSVals := []string{"|", "J", "L", "7", "F", "-"}

	for y, line := range lines {
		chars := []string{}
		for x, char := range strings.Split(line, "") {
			chars = append(chars, char)
			if char == "S" {
				sPos.xpos = x
				sPos.ypos = y
				seen = append(seen, sPos)
				queue = push(queue, sPos)
			}
		}
		pipeMap = append(pipeMap, chars)
	}

	for len(queue) > 0 {
		pos := position{}
		queue, pos = pop(queue)
		char := pipeMap[pos.ypos][pos.xpos]

		// check if there is a valid pipe to move through above pos if so add to queue
		if pos.ypos > 0 && (char == "S" || char == "|" || char == "J" || char == "L") {
			moveChar := pipeMap[pos.ypos-1][pos.xpos]
			if !haveSeen(seen, pos.xpos, pos.ypos-1) && (moveChar == "|" || moveChar == "7" || moveChar == "F") {
				newPos := position{xpos: pos.xpos, ypos: pos.ypos - 1}
				seen = append(seen, newPos)
				queue = push(queue, newPos)
				if char == "S" {
					possibleSVals = intersection(possibleSVals, []string{"|", "J", "L"})
				}
			}
		}

		// check if there is a valid pipe to move through right of pos if so add to queue
		if pos.xpos < len(pipeMap[pos.ypos])-1 && (char == "S" || char == "-" || char == "L" || char == "F") {
			moveChar := pipeMap[pos.ypos][pos.xpos+1]
			if !haveSeen(seen, pos.xpos+1, pos.ypos) && (moveChar == "-" || moveChar == "7" || moveChar == "J") {
				newPos := position{xpos: pos.xpos + 1, ypos: pos.ypos}
				seen = append(seen, newPos)
				queue = append(queue, newPos)
				if char == "S" {
					possibleSVals = intersection(possibleSVals, []string{"-", "L", "F"})
				}

			}
		}

		// check if there is a valid pipe to move through left of pos if so add to queue
		if pos.xpos > 0 && (char == "S" || char == "-" || char == "7" || char == "J") {
			moveChar := pipeMap[pos.ypos][pos.xpos-1]
			if !haveSeen(seen, pos.xpos-1, pos.ypos) && (moveChar == "-" || moveChar == "L" || moveChar == "F") {
				newPos := position{xpos: pos.xpos - 1, ypos: pos.ypos}
				seen = append(seen, newPos)
				queue = push(queue, newPos)
				if char == "S" {
					possibleSVals = intersection(possibleSVals, []string{"-", "7", "J"})
				}
			}
		}

		// check if there is a valid pipe to move through below pos if so add to queue
		if pos.ypos < len(pipeMap) && (char == "S" || char == "|" || char == "7" || char == "F") {
			moveChar := pipeMap[pos.ypos+1][pos.xpos]
			if !haveSeen(seen, pos.xpos, pos.ypos+1) && (moveChar == "S" || moveChar == "|" || moveChar == "J" || moveChar == "L") {
				newPos := position{xpos: pos.xpos, ypos: pos.ypos + 1}
				seen = append(seen, newPos)
				queue = push(queue, newPos)
				if char == "S" {
					possibleSVals = intersection(possibleSVals, []string{"|", "7", "F"})
				}
			}
		}
	}

	if len(possibleSVals) > 1 {
		fmt.Printf("More than one possible value for S: %+v\n", possibleSVals)
		return -1
	}

	pipeMap[sPos.ypos][sPos.xpos] = possibleSVals[0]

	for yIdx, yVal := range pipeMap {
		for xIdx := range yVal {
			if !haveSeen(seen, xIdx, yIdx) {
				pipeMap[yIdx][xIdx] = "."
			}
		}
	}

	// This isn't working correctly
	// I'm trying to scan horizontally through the grid and mark each spot that could be considered outside
	// But the tricky part is accounting for the pipe directions and whether or not you can step around it
	outside := []position{}
	for row, yVal := range pipeMap {
		for column, xVal := range yVal {
			withinLoop := false
			var up bool
			if xVal == "|" {
				withinLoop = !withinLoop
			} else if xVal == "L" || xVal == "F" {
				up = xVal == "L"
			} else if xVal == "7" || xVal == "J" {
				if (up && xVal == "J") || (!up && xVal == "7") {
					withinLoop = !withinLoop
				}
			}
			if withinLoop == false {
				out := position{xpos: column, ypos: row}
				outside = append(outside, out)
			}
		}
	}

	return (len(pipeMap)*len(pipeMap[0]) - (len(seen) + len(outside)))
}

func haveSeen(seen []position, xpos int, ypos int) bool {

	for _, v := range seen {
		if v.xpos == xpos && v.ypos == ypos {
			return true
		}
	}

	return false
}

func intersection(a, b []string) []string {

	// uses empty struct (0 bytes) for map values.
	m := make(map[string]struct{}, len(b))

	// cached
	for _, v := range b {
		m[v] = struct{}{}
	}

	var s []string
	for _, v := range a {
		if _, ok := m[v]; ok {
			s = append(s, v)
		}
	}

	return s
}
