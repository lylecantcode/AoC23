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
	if partOne(myLib.ErrHandledReadConv("test_input.txt")) != 114 {
		log.Fatal("incorrect response from test input")
	}
}

func partTwo(input []string) int {
	return 0
}

func partOne(input []string) int {
	var rows [][][]int
	var result int
	for i := 0; i < len(input); i++ {
		rows = append(rows, [][]int{myLib.StringToIntArray(input[i])})
		_, add := rowInterval(rows[i], 0)
		result += add
	}
	return result
}

func rowInterval(input [][]int, steps int) ([]int, int) {
	if len(input) == 0 {
		return nil, 0
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
		return rowInterval(input, steps)
	} else {
		newVal := 0
		for i := 0; i < len(input); i++ {
			if len(input[i]) != 0 {
				newVal += input[i][len(input[i])-1]
				// fmt.Printf("%v\n", input[i])
			}
		}
		// fmt.Printf("%v\n", newVal)
		return nil, newVal
	}
}
