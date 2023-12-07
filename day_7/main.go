package main

import (
	"aoc23/myLib"
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := myLib.ErrHandledReadConv("input.txt")
	fmt.Println(partTwo(input))
}

type HandType int

// first numbers on tie
type rank struct {
	hand   string
	bid    int
	values HandType
}

const (
	single HandType = iota + 1
	pair
	twoPair
	threes
	fullHouse
	fours
	fives
)

func partTwo(input []string) int {
	cards := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}
	// order input by strength, keep track of "bid", bid*strength (1 based index)
	var order []rank
	// -1 because of extra line at end of input
	strength := len(input) - 1
	for i := 0; i < strength; i++ {
		handVal := single
		// 5 cards
		hand := input[i][:5]
		jokers := strings.Count(hand, "J")
		for j := 0; j < len(cards); j++ {
			val := strings.Count(hand, cards[j])
			switch val {
			case 5:
				handVal = fives
			case 4:
				handVal = fours
			case 3:
				if handVal == pair {
					handVal = fullHouse
				} else {
					handVal = threes
				}
			case 2:
				if handVal == threes {
					handVal = fullHouse
				} else if handVal == pair {
					handVal = twoPair
				} else {
					handVal = pair
				}
			}
		}
		if jokers > 0 {
			switch handVal {
			case fours: // 4 of a kind -> 5
				handVal = fives
			case threes: // 3 of a kind -> 4, 5
				if jokers == 2 {
					handVal = fives
				} else {
					handVal = fours
				}
			case twoPair: // 2 pair -> full house
				handVal = fullHouse
			case pair: // 1 pair  -> 3, 4, 5
				if jokers == 3 {
					handVal = fives
				} else if jokers == 2 {
					handVal = fours
				} else {
					handVal = threes
				}
			case single:
				if jokers >= 4 {
					handVal = fives
				} else if jokers == 3 {
					handVal = fours
				} else if jokers == 2 {
					handVal = threes
				} else {
					handVal = pair
				}
			}
		}
		order = append(order, rank{hand: hand, values: handVal, bid: myLib.ErrHandledAtoi(strings.TrimSpace(input[i][5:]))})
	}
	for i := 0; i < len(order); i++ {
		order[i].hand = strings.ReplaceAll(order[i].hand, "A", "Z")
		order[i].hand = strings.ReplaceAll(order[i].hand, "K", "Y")
		order[i].hand = strings.ReplaceAll(order[i].hand, "Q", "X")
		order[i].hand = strings.ReplaceAll(order[i].hand, "J", "1")
	}

	sort.SliceStable(order, func(i, j int) bool { return strings.Compare(order[i].hand, order[j].hand) < 1 })
	sort.SliceStable(order, func(i, j int) bool { return order[i].values < order[j].values })
	total := 0
	for i := 0; i < strength; i++ {
		total += order[i].bid * (i + 1)
	}
	return total
}
