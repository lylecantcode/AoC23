package main

import (
	"aoc23/myLib"
	"fmt"
	"log"
)

func main() {
	input := myLib.ErrHandledReadConv("input.txt")
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
	test := myLib.ErrHandledReadConv("test_input.txt")
	if partOne(test) != 114 {
		log.Fatal("incorrect response from test input")
	}
	if partTwo(test) != 2 {
		log.Fatal("incorrect response from test input", partTwo(test))
	}
}

func partTwo(input []string) int {
	var rows [][][]int
	var result int
	for i := 0; i < len(input); i++ {
		rows = append(rows, [][]int{myLib.StringToIntArray(input[i])})
		_, add, _ := rowInterval(rows[i], 0, true)
		result += add
	}
	return result
}
func partOne(input []string) int {
	var rows [][][]int
	var result int
	for i := 0; i < len(input); i++ {
		rows = append(rows, [][]int{myLib.StringToIntArray(input[i])})
		_, add, _ := rowInterval(rows[i], 0, false)
		result += add
	}
	return result
}

func rowInterval(input [][]int, steps int, first bool) ([]int, int, bool) {
	if len(input) == 0 {
		return nil, 0, first
	}
	interval := 0
	finished := true
	input = append(input, []int{})
	for i := 0; i < len(input[steps])-1; i++ {
		length := len(input) - 1
		interval = input[steps][i+1] - input[steps][i]
		input[length] = append(input[length], interval)
		if interval != 0 {
			finished = false
		}
	}
	steps++
	if !finished {
		return rowInterval(input, steps, first)
	} else {
		newVal := 0
		for i := 0; i < len(input); i++ {
			if len(input[i]) != 0 {
				if first {
					newVal = input[len(input)-i-1][0] - newVal
				} else {
					newVal += input[i][len(input[i])-1]
				}
			}
		}
		return nil, newVal, first
	}
}
