package main

import (
	mylib "aoc23/library"
	"fmt"
	"log"
	"os"
	"regexp"
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
	re := regexp.MustCompile(`[^0-9]+`)
	time := mylib.ErrHandledAtoi(re.ReplaceAllString(input[0], ""))
	dist := mylib.ErrHandledAtoi(re.ReplaceAllString(input[1], ""))
	minWin := 0
	maxWin := 0
	for j := 0; j < time; j++ {
		if j*(time-j) > dist {
			minWin = j
			break
		}
	}
	for i := time; i >= 0; i-- {
		if i*(time-i) > dist {
			maxWin = i
			break
		}
	}
	fmt.Println(maxWin - minWin + 1)
}

func partOne(input []string) {
	times := strings.Fields(input[0])[1:]
	dists := strings.Fields(input[1])[1:]
	winning := 1
	for i := 0; i < len(times); i++ {
		winRace := 0
		time := mylib.ErrHandledAtoi(times[i])
		for j := 0; j < time; j++ {
			if j*(time-j) > mylib.ErrHandledAtoi(dists[i]) {
				winRace++
			}
		}
		winning *= winRace
	}
	fmt.Println(winning)
}
