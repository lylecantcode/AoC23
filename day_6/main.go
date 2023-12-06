package main

import (
	"aoc23/myLib"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	input := myLib.ErrHandledReadConv("input.txt")
	partOne(input)
	partTwo(input)
}

func partTwo(input []string) {
	re := regexp.MustCompile(`[^0-9]+`)
	time := myLib.ErrHandledAtoi(re.ReplaceAllString(input[0], ""))
	dist := myLib.ErrHandledAtoi(re.ReplaceAllString(input[1], ""))
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
		time := myLib.ErrHandledAtoi(times[i])
		for j := 0; j < time; j++ {
			if j*(time-j) > myLib.ErrHandledAtoi(dists[i]) {
				winRace++
			}
		}
		winning *= winRace
	}
	fmt.Println(winning)
}
