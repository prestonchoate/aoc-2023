package day11

import (
	"testing"

	"github.com/prestonchoate/aoc-2023/solutions"
)

func TestP1(t *testing.T) {
	input := solutions.GetInputString("../../inputs/day11-sample.txt")

	expected := 374

	got := SolveP1(input)

	if got != expected {
		t.Fatalf("Expected %v, Got %v\n", expected, got)
	}
}

func TestP2(t *testing.T) {

}

func TestCalcDist(t *testing.T) {
	/*
		Empty Rows: map[0:false 1:false 2:false 3:true 4:false 5:false 6:false 7:true 8:false 9:false]

		Empty Cols: map[0:false 1:false 2:true 3:false 4:false 5:true 6:false 7:false 8:true 9:false]

		Galaxies: [{xPos:3 yPos:0 id:1} {xPos:7 yPos:1 id:2} {xPos:0 yPos:2 id:3} {xPos:6 yPos:4 id:4} {xPos:1 yPos:5 id:5} {xPos:9 yPos:6 id:6} {xPos:7 yPos:8 id:7} {xPos:0 yPos:9 id:8} {xPos:4 yPos:9 id:9}]
	*/

	eRows := map[int]bool{
		0: false, 1: false, 2: false, 3: true, 4: false, 5: false, 6: false, 7: true, 8: false, 9: false,
	}

	eCols := map[int]bool{
		0: false, 1: false, 2: true, 3: false, 4: false, 5: true, 6: false, 7: false, 8: true, 9: false,
	}

	start := galaxy{
		xPos: 3,
		yPos: 0,
		id:   1,
	}

	target := galaxy{
		xPos: 0,
		yPos: 2,
		id:   3,
	}

	expected := 6

	got := calcDistance(start, target, eRows, eCols, 1)

	if got != expected {
		t.Fatalf("Testing distance between galaxy %v and %v: Expected %v, Got %v\n", start.id, target.id, expected, got)
	}

	target.xPos = 7
	target.yPos = 1
	target.id = 2

	expected = 6
	got = calcDistance(start, target, eRows, eCols, 1)

	if got != expected {
		t.Fatalf("Testing distance between galaxy %v and %v: Expected %v, Got %v\n", start.id, target.id, expected, got)
	}

	target.id = 7
	target.xPos = 7
	target.yPos = 8

	expected = 15
	got = calcDistance(start, target, eRows, eCols, 1)

	if got != expected {
		t.Fatalf("Testing distance between galaxy %v and %v: Expected %v, Got %v\n", start.id, target.id, expected, got)
	}

	start.id = 3
	start.xPos = 0
	start.yPos = 2

	target.id = 6
	target.xPos = 9
	target.yPos = 6

	expected = 17
	got = calcDistance(start, target, eRows, eCols, 1)

	if got != expected {
		t.Fatalf("Testing distance between galaxy %v and %v: Expected %v, Got %v\n", start.id, target.id, expected, got)
	}

	start.id = 8
	start.xPos = 0
	start.yPos = 9

	target.id = 9
	target.xPos = 4
	target.yPos = 9

	expected = 5
	got = calcDistance(start, target, eRows, eCols, 1)

	if got != expected {
		t.Fatalf("Testing distance between galaxy %v and %v: Expected %v, Got %v\n", start.id, target.id, expected, got)
	}
}
