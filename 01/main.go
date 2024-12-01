package main

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/kubarydz/advent2024/utils"
)

func main() {
	sample := utils.ChunkInput(utils.ReadInput("sample"))
	resp := part1(sample)
	if resp != 11 {
		fmt.Printf("part 1 failed, got %d want %d\n", resp, 11)
		return
	}
	input := utils.ChunkInput(utils.ReadInput("input"))
	resp = part1(input)
	fmt.Println("part1: ", resp)

	resp = part2(sample)
	if resp != 31 {
		fmt.Printf("part 2 failed, got %d want %d\n", resp, 31)
		return
	}
	resp = part2(input)
	fmt.Println("part2: ", resp)
}

func part1(input [][]byte) int {
	first, second := getSortedLists(input)
	sum := 0
	for i := 0; i < len(first); i++ {
		sum += utils.Abs(first[i] - second[i])
	}

	return sum
}

func getSortedLists(input [][]byte) ([]int, []int) {
	first := []int{}
	second := []int{}

	for _, line := range input {
		start := 0
		end := 0
		for i, char := range line {
			if char == ' ' {
				end = i
				break
			}
		}
		number := string(line[start:end])
		nr, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		first = append(first, nr)
		number = string(line[end+3:])
		nr, err = strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		second = append(second, nr)
	}

	slices.Sort(first)
	slices.Sort(second)

	return first, second
}

func part2(input [][]byte) int {
	first, second := getSortedLists(input)
	score := 0
	for _, firstNr := range first {
		counter := 0
		for _, secondNr := range second {
			if firstNr == secondNr {
				counter++
			} else if firstNr < secondNr {
				break
			}
		}
		score += counter * firstNr
	}
	return score
}
