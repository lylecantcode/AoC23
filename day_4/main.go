package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("failed to read input")
	}
	inputLines := strings.Split(string(input), "\n")
	partOne(inputLines)
	partTwo(inputLines)
}

func partTwo(input []string) {
	// use part 1 loop and store in an array and then multiply the array
	pointsSlice := [][]int{}
	for i := 0; i < len(input); i++ {
		if len(input[i]) == 0 {
			break
		}
		ticketValue := 0
		ticket, winning, _ := strings.Cut(input[i][9:], "|")

		ticketVals := strings.Fields(ticket)
		winningVals := strings.Fields(winning)
		for j := 0; j < len(ticketVals); j++ {
			if slices.Contains(winningVals, ticketVals[j]) {
				ticketValue++
			}
		}
		pointsSlice = append(pointsSlice, []int{ticketValue, 1})
	}
	fmt.Println(pointsSlice)

	total := 0
	for i := 0; i < len(pointsSlice); i++ {
		newcards := pointsSlice[i][0]
		for j := 1; j <= newcards; j++ {
			if i+j < len(pointsSlice) {
				pointsSlice[i+j][1] += pointsSlice[i][1]
			}
		}
		total += pointsSlice[i][1]

	}
	fmt.Println(total)
}

func partOne(input []string) {
	points := 0
	for i := 0; i < len(input); i++ {
		if len(input[i]) == 0 {
			break
		}
		ticketValue := 0
		ticket, winning, _ := strings.Cut(input[i][9:], "|")

		ticketVals := strings.Fields(ticket)
		winningVals := strings.Fields(winning)
		for j := 0; j < len(ticketVals); j++ {
			if slices.Contains(winningVals, ticketVals[j]) {
				if ticketValue == 0 {
					ticketValue = 1
				} else {
					ticketValue *= 2
				}
			}
		}
		points += ticketValue
	}
	fmt.Println(points)
}
