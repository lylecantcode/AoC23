package main

import (
	"aoc23/myLib"
	"fmt"
	"strings"
)

func main() {
	input := myLib.ErrHandledRead("input.txt")
	partOne(input)
	partTwo(input)
}

func partTwo(input []byte) {
	total := 0
	inputLines := strings.Split(string(input), "\n")
	inputMap := map[string]int{
		// "zero":  0,
		"0":     0,
		"one":   1,
		"1":     1,
		"two":   2,
		"2":     2,
		"three": 3,
		"3":     3,
		"four":  4,
		"4":     4,
		"five":  5,
		"5":     5,
		"six":   6,
		"6":     6,
		"seven": 7,
		"7":     7,
		"eight": 8,
		"8":     8,
		"nine":  9,
		"9":     9,
	}
	for i := 0; i < len(inputLines); i++ {
		line := inputLines[i]
		var min, firstVal, max, lastVal int = len(line), 0, 0, 0
		for k, v := range inputMap {
			firstIndex := strings.Index(line, k)
			if firstIndex != -1 {
				if firstIndex <= min {
					min = firstIndex
					firstVal = v
				}
				lastIndex := strings.LastIndex(line, k)
				if lastIndex >= max {
					max = lastIndex
					lastVal = v
				}

			}
		}
		total += 10*firstVal + lastVal
	}
	fmt.Println(total)
}

func partOne(input []byte) {
	total := 0
	var lineIntSlice []int
	for i := 0; i < len(input); i++ {
		val := input[i]
		// []byte(\n) == 10
		if val == 10 || i == len(input)-1 {
			if len(lineIntSlice) == 1 {
				total += lineIntSlice[0]*10 + lineIntSlice[0]
			} else if len(lineIntSlice) > 0 {
				total += lineIntSlice[0]*10 + lineIntSlice[len(lineIntSlice)-1]
			}
			lineIntSlice = nil
			// []byte("0","9") 48 -> 57 = 0 -> 9
		} else if val >= 48 && val <= 57 {
			lineIntSlice = append(lineIntSlice, int(val)-48)
		}
	}
	fmt.Println(total)
}
