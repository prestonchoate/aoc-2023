package main

import (
	"flag"
	"fmt"

	"github.com/prestonchoate/aoc-2023/handlers"
)

func main() {
	daySelect := flag.Int("day", 0, "Select which day to execute")
	partSelect := flag.Int("part", 0, "Select which part to execute")
	flag.Parse()

	switch *daySelect {
	case 1:
		handlers.HandleDay1(*partSelect)
		break
	default:
		fmt.Printf("%v is an invalid day. Please try again\n", *daySelect)
		break
	}
}
