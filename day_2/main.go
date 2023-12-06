package main

import (
	"aoc23/myLib"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type colours struct {
	red   int
	blue  int
	green int
}

func main() {
	solving(myLib.ErrHandledReadConv("input.txt"))
}

func solving(input []string) {
	contained := colours{
		red:   12,
		green: 13,
		blue:  14,
	}

	possible := 0
	total := 0
	// split on : , and ;
	for i := 0; i < len(input); i++ {
		inputLine := input[i]
		if len(inputLine) == 0 {
			break
		}
		gameNumber, gameCount, _ := strings.Cut(inputLine[5:], ":")
		gameNum, err := strconv.Atoi(gameNumber)
		if err != nil {
			log.Fatal("parsing game number failed: ", err)
		}
		// for part 1:
		if colourCount(gameCount, " green", contained.green) && colourCount(gameCount, " red", contained.red) && colourCount(gameCount, " blue", contained.blue) {
			possible += gameNum
		}
		// for part 2:
		total += power(gameCount, " green") * power(gameCount, " red") * power(gameCount, " blue")
	}
	fmt.Println("part 1:", possible)
	fmt.Println("part 2:", total)
}
func power(input, colour string) int {
	max := 0
	games := strings.Split(input, colour)
	for i := 0; i < len(games); i++ {
		if len(games[i]) == 0 {
			break
		}
		game := games[i][len(games[i])-2:]

		values := strings.FieldsFunc(game, func(r rune) bool {
			return !unicode.IsNumber(r)
		})
		if len(values) == 0 {
			continue
		}

		val := myLib.ErrHandledAtoi(values[0])

		if val > max {
			max = val
		}
	}
	return max
}

// was used for part 1
func gameCount(input, colour string, max int) bool {
	games := strings.Split(input, colour)
	for i := 0; i < len(games); i++ {
		if len(games[i]) == 0 {
			break
		}
		game := games[i][len(games[i])-2:]

		values := strings.FieldsFunc(game, func(r rune) bool {
			return !unicode.IsNumber(r)
		})
		if len(values) == 0 {
			continue
		}
		ints := make(sort.IntSlice, len(values))
		for i, s := range values {
			ints[i] = myLib.ErrHandledAtoi(s)
		}
		ints.Sort()
		if ints[len(ints)-1] > max {
			return false
		}
	}
	return true
}

func colourCount(input, colour string, max int) bool {
	games := strings.Split(input, colour)
	for i := 0; i < len(games); i++ {
		if len(games[i]) == 0 {
			break
		}
		game := games[i][len(games[i])-2:]

		values := strings.FieldsFunc(game, func(r rune) bool {
			return !unicode.IsNumber(r)
		})
		if len(values) == 0 {
			continue
		}
		ints := make(sort.IntSlice, len(values))
		for i, s := range values {
			ints[i] = myLib.ErrHandledAtoi(s)
		}
		ints.Sort()
		if ints[len(ints)-1] > max {
			return false
		}
	}
	return true
}
