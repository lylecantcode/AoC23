package main

import (
	"aoc23/myLib"
	"fmt"
	"log"
	"strings"
	"unicode"
)

func main() {
	input := myLib.ErrHandledReadConv("input.txt")
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func partTwo(input []string) int {
	return 0
}

func partOne(input []string) int {
	// for extra line
	body := input[2:]
	dirMap := make(map[string][]string)
	endPoints := []string{}
	for i := 0; i < len(body); i++ {

		if len(body[i]) == 0 {
			break
		}
		values := strings.FieldsFunc(body[i], func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r) // number added for the test data
		})
		dirMap[values[0]] = []string{values[1], values[2]}
		if values[0][2] == 'Z' {
			endPoints = append(endPoints, values[0])
		}
	}
	directions := input[0]
	// testing loop period, ZZZ loops back to AAAs children (but reversed)
	fmt.Println(endPoints)
	// run to figure out loop interval for each Z:
	var values []int

	for _, start := range endPoints {
		steps := 0
		loops := 0
		current := start
		for {
			loops++
			for i := 0; i < len(directions); i++ {
				steps++
				slice, exist := dirMap[current]
				if !exist {
					log.Fatal("invalid map key")
				}
				if directions[i] == 'L' {
					current = slice[0]
				} else if directions[i] == 'R' {
					current = slice[1]
				} else {
					fmt.Println(slice, directions[i])
					break
				}
			}
			fmt.Printf("steps: %v, direction loops: %v\n", steps, loops)
			if current == start {
				break
			}
		}
		values = append(values, steps)
	}

	return lcm(values)
}

func lcm(ints []int) int {
	result := ints[0] * ints[1] / gcd(ints[0], ints[1])

	for i := 2; i < len(ints); i++ {
		result = lcm([]int{result, ints[i]})
	}

	return result
}
func gcd(total, multi int) int {
	for multi != 0 {
		x := multi
		multi = total % multi
		total = x
	}
	return total
}
