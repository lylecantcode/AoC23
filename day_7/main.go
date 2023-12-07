package main

import (
	"aoc23/myLib"
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := myLib.ErrHandledReadConv("input.txt")
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func partTwo(input []string) int {
	return 0
}

// 7:5, 6:4, 5:full house, 4:3, 3:2*2, 2:2, 1:highest
// first numbers on tie
type rank struct {
	hand   string
	bid    int
	values int
}

func partOne(input []string) int {
	cards := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	// order input by strength, keep track of "bid", bid*strength (1 based index)
	var order []rank
	// 5 cards
	// -1 because of extra line at end of input
	strength := len(input) - 1
	for i := 0; i < strength; i++ {
		handVal := 1
		hand := input[i][:5]
		for j := 0; j < len(cards); j++ {
			val := strings.Count(hand, cards[j])
			switch val {
			case 5:
				handVal = 7
			case 4:
				handVal = 6
			case 3:
				if handVal == 2 {
					handVal = 5
				} else {
					handVal = 4
				}
			case 2:
				if handVal == 3 {
					handVal = 5
				} else if handVal == 2 {
					handVal = 3
				} else {
					handVal = 2
				}
			}
		}
		order = append(order, rank{hand: hand, values: handVal, bid: myLib.ErrHandledAtoi(strings.TrimSpace(input[i][5:]))})
	}
	for i := 0; i < len(order); i++ {
		order[i].hand = strings.ReplaceAll(order[i].hand, "A", "Z")
		order[i].hand = strings.ReplaceAll(order[i].hand, "K", "Y")
		order[i].hand = strings.ReplaceAll(order[i].hand, "Q", "X")
		order[i].hand = strings.ReplaceAll(order[i].hand, "J", "W")
	}
	sort.SliceStable(order, func(i, j int) bool { return strings.Compare(order[i].hand, order[j].hand) < 1 })
	sort.SliceStable(order, func(i, j int) bool { return order[i].values < order[j].values })
	total := 0
	for i := 0; i < strength; i++ {
		total += order[i].bid * (i + 1)
	}
	fmt.Println(order)
	return total
}
