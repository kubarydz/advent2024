package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(filename string) []byte {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("cannot open %v, error: %v\n", filename, err)
		panic("cannot open file")
	}
	return bytes
}

func ChunkInput(input []byte) [][]byte {
	var chunks [][]byte
	lastStart := 0
	for i, b := range input {
		if b == '\n' {
			chunks = append(chunks, input[lastStart:i])
			lastStart = i + 1

		}
	}
	return chunks
}

func ReadInts(filename string) [][]int {
	bytes := ReadChunks(filename)
	result := [][]int{}
	for _, row := range bytes {
		rowStr := strings.Split(string(row), " ")
		rowInt := []int{}
		for _, rstr := range rowStr {
			i, err := strconv.Atoi(rstr)
			if err != nil {
				panic(err)
			}
			rowInt = append(rowInt, i)
		}
		result = append(result, rowInt)
	}
	return result
}

func ReadChunks(filename string) [][]byte {
	return ChunkInput(ReadInput(filename))
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
