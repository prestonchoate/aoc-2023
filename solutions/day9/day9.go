package day9

import (
	"fmt"
	"strconv"
	"strings"
)

func SolvePart1(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		nums := getNums(line)
		// fmt.Printf("Nums: %+v\n", nums)
		nextNum := getNextNum(nums)
		// fmt.Printf("Got next num: %v\n", nextNum)
		sum += nextNum
	}

	return sum
}

func SolvePart2(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		nums := getNums(line)
		// fmt.Printf("Nums: %+v\n", nums)
		prevNum := getPrevNum(nums)
		// fmt.Printf("Got prev num: %v\n", prevNum)
		sum += prevNum
	}

	return sum
}

func getNums(line string) []int {
	numStrs := strings.Split(line, " ")
	nums := []int{}
	for _, numStr := range numStrs {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		nums = append(nums, num)
	}
	return nums
}

func getPattern(nums []int) []int {
	pattern := []int{}
	for i := 0; i < len(nums)-1; i++ {
		pattern = append(pattern, nums[i+1]-nums[i])
	}

	return pattern
}

func getNextNum(nums []int) int {
	if allZero(nums) {
		return 0
	}

	pattern := getPattern(nums)
	diff := getNextNum(pattern)
	return nums[len(nums)-1] + diff
}

func getPrevNum(nums []int) int {
	if allZero(nums) {
		return 0
	}

	pattern := getPattern(nums)
	diff := getPrevNum(pattern)
	return nums[0] - diff
}

func allZero(pattern []int) bool {
	sum := 0
	for _, num := range pattern {
		sum += num
	}
	return sum == 0
}
