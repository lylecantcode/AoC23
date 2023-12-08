package main

import (
	"aoc23/myLib"
	"fmt"
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
	starting := []string{}
	for i := 0; i < len(body); i++ {

		if len(body[i]) == 0 {
			break
		}
		values := strings.FieldsFunc(body[i], func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r) // number added for the test data
		})
		dirMap[values[0]] = []string{values[1], values[2]}
		if values[0][2] == 'A' {
			starting = append(starting, values[0])
		}
	}
	steps := 0
	directions := input[0]
	current := starting
	fmt.Println(starting)
mainLoop:
	for {
		for i := 0; i < len(directions); i++ {
			steps++
			for j := 0; j < len(current); j++ {
				slice := dirMap[current[j]]
				if directions[i] == 'L' {
					current[j] = slice[0]
				} else if directions[i] == 'R' {
					current[j] = slice[1]
				} else {
					fmt.Println(slice, directions[i])
					break
				}
			}
			allZ := true
			// fmt.Printf("%v->(%s)\n", current, string(directions[i]))
			for k := 0; k < len(starting) && allZ == true; k++ {
				if current[k][2] != 'Z' {
					allZ = false
				}
			}
			if steps%1000 == 0 {
				fmt.Println(steps, "still running")
			}
			if allZ {
				break mainLoop
			}
		}

	}
	return steps
}
