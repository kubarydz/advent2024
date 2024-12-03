package main

import (
	"fmt"
	"slices"

	"github.com/kubarydz/advent2024/utils"
)

const (
	respSample1 = 2
	respSample2 = 4
)

func main() {
	sample := utils.ReadInts("sample")
	resp := part1(sample)
	if resp != respSample1 {
		fmt.Printf("part 1 failed, got %d want %d\n", resp, respSample1)
		return
	}
	input := utils.ReadInts("input")
	resp = part1(input)
	fmt.Println("part1: ", resp)

	resp = part2(sample)
	if resp != respSample2 {
		fmt.Printf("part 2 failed, got %d want %d\n", resp, respSample2)
		return
	}
	resp = part2(input)
	fmt.Println("part2: ", resp)
}

func part1(input [][]int) int {
	counter := 0
	for _, row := range input {
		if checkSafety(row) {
			counter++
		}
	}
	return counter
}

func part2(input [][]int) int {
	counter := 0
	for _, row := range input {
		if checkSafety((row)) {
			counter++
			continue
		}
		for i := 0; i < len(row); i++ {
			if checkSafety(slices.Concat(row[:i], row[i+1:])) {
				counter++
				break
			}
		}
	}
	return counter
}

func checkSafety(input []int) bool {
	increasing := input[0] < input[1]

	fmt.Println(input)
	last := -1
	for _, nr := range input {
		if last == -1 {
			last = nr
			continue
		}
		if nr == last {
			return false
		}
		if (increasing && nr < last) || (!increasing && nr > last) {
			return false
		}
		if utils.Abs(nr-last) > 3 {
			return false
		}
		last = nr
	}
	return true
}
