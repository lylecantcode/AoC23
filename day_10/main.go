package main

import (
	"aoc23/myLib"
	"fmt"
	"log"
	"strings"
)

func main() {
	input := myLib.ErrHandledReadConv("input.txt")
	input = input[:len(input)-1]
	log.Printf("part two: %v\n\n", partTwo(input))

	test := myLib.ErrHandledReadConv("test_input.txt")
	testTwo := partTwo(test)
	if testTwo != 8 {
		log.Fatal("incorrect response from part two test: ", testTwo)
	}

}

const (
	pipeNS = '|'
	pipeEW = '-'
	pipeNE = 'L'
	pipeNW = 'J'
	pipeSW = '7'
	pipeSE = 'F'
	ground = '.'
	start  = 'S'
)

type position struct {
	symbol     byte
	horizontal int
	vertical   int
	stepsTaken int
	from       byte
}

func partTwo(input []string) int {
	// establish bounaries, iterate through and check
	tracker := [][]*bool{}
	for i := 0; i < len(input); i++ {
		tracker = append(tracker, make([]*bool, len(input[i])))
	}

	var p position
	for i := 0; i < len(input); i++ {
		current := strings.IndexByte(input[i], start)
		if current != -1 {
			p = position{start, current, i, 0, start}
			break
		}
	}
	var n, s, e, w int
	n = p.north(input, tracker)
	s = p.south(input, tracker)
	e = p.east(input, tracker)
	w = p.west(input, tracker)
	log.Println(n, s, e, w)

	for i := 0; i < len(tracker); i++ {
		for j := 0; j < len(tracker[i]); j++ {
			if tracker[i][j] == nil {
				propogate(input, tracker, i, j)
			}
		}
	}

	// forgot about squeezing between pipes
	// check around each "false"

	count := 0
	for i := 0; i < len(tracker); i++ {
		for j := 0; j < len(tracker[i]); j++ {
			if tracker[i][j] == nil {
				count++
			}
			// 	fmt.Printf(".")
			// } else if *tracker[i][j] {
			// 	fmt.Printf("#")
			// } else {
			// 	fmt.Printf(" ")
			// }
		}
		fmt.Printf("\n")
	}

	return count

}

func propogate(input []string, tracker [][]*bool, i, j int) {
	if i < 0 || j < 0 || i > len(tracker)-1 || j > len(tracker[i])-1 {
		return
	}
	if tracker[i][j] == nil && (i == 0 || j == 0 || i == len(tracker)-1 || j == len(tracker[i])-1 ||
		(tracker[i-1][j] != nil && *tracker[i-1][j] == false) || (tracker[i+1][j] != nil && *tracker[i+1][j] == false) ||
		(tracker[i][j-1] != nil && *tracker[i][j-1] == false) || (tracker[i][j+1] != nil && *tracker[i][j+1] == false)) {

		tracker[i][j] = myLib.BoolPtr(false)

		propogate(input, tracker, i+1, j)
		propogate(input, tracker, i-1, j)
		propogate(input, tracker, i, j+1)
		propogate(input, tracker, i, j-1)
	}
}

func (p position) north(input []string, tracker [][]*bool) int {
	if p.vertical < 1 {
		return 0
	}
	np := position{
		from:       'n',
		symbol:     input[p.vertical-1][p.horizontal],
		horizontal: p.horizontal,
		vertical:   p.vertical - 1,
		stepsTaken: p.stepsTaken + 1,
	}
	// log.Println(string(np.from), string(np.symbol), np.stepsTaken)
	switch np.symbol {
	case pipeNS:
		tracker[np.vertical][np.horizontal] = myLib.BoolPtr(true)
		return np.north(input, tracker)
	case pipeSE:
		tracker[np.vertical][np.horizontal] = myLib.BoolPtr(true)
		return np.east(input, tracker)
	case pipeSW:
		tracker[np.vertical][np.horizontal] = myLib.BoolPtr(true)
		return np.west(input, tracker)
	case start:
		tracker[np.vertical][np.horizontal] = myLib.BoolPtr(true)
		return np.stepsTaken
	default:
		log.Println("ended north start")
		return 0
	}
}

func (p position) south(input []string, tracker [][]*bool) int {
	if p.vertical >= len(input)-1 {
		return 0
	}
	np := position{
		from:       's',
		symbol:     input[p.vertical+1][p.horizontal],
		horizontal: p.horizontal,
		vertical:   p.vertical + 1,
		stepsTaken: p.stepsTaken + 1,
	}
	// log.Println(string(np.from), string(np.symbol), np.stepsTaken)
	switch np.symbol {
	case pipeNS:
		tracker[np.vertical][np.horizontal] = myLib.BoolPtr(true)
		return np.south(input, tracker)
	case pipeNE:
		tracker[np.vertical][np.horizontal] = myLib.BoolPtr(true)
		return np.east(input, tracker)
	case pipeNW:
		tracker[np.vertical][np.horizontal] = myLib.BoolPtr(true)
		return np.west(input, tracker)
	case start:
		tracker[np.vertical][np.horizontal] = myLib.BoolPtr(true)
		return np.stepsTaken
	default:
		log.Println("ended south start")
		return 0
	}
}

func (p position) east(input []string, tracker [][]*bool) int {
	if p.horizontal >= len(input[0])-1 {
		return 0
	}
	np := position{
		from:       'e',
		symbol:     input[p.vertical][p.horizontal+1],
		horizontal: p.horizontal + 1,
		vertical:   p.vertical,
		stepsTaken: p.stepsTaken + 1,
	}
	// log.Println(string(np.from), string(np.symbol), np.stepsTaken)
	switch np.symbol {
	case pipeEW:
		tracker[np.vertical][np.horizontal] = myLib.BoolPtr(true)
		return np.east(input, tracker)
	case pipeNW:
		tracker[np.vertical][np.horizontal] = myLib.BoolPtr(true)
		return np.north(input, tracker)
	case pipeSW:
		tracker[np.vertical][np.horizontal] = myLib.BoolPtr(true)
		return np.south(input, tracker)
	case start:
		tracker[np.vertical][np.horizontal] = myLib.BoolPtr(true)
		return np.stepsTaken
	default:
		log.Println("ended east start")
		return 0
	}
}

func (p position) west(input []string, tracker [][]*bool) int {
	if p.horizontal < 1 {
		return 0
	}
	np := position{
		from:       'w',
		symbol:     input[p.vertical][p.horizontal-1],
		horizontal: p.horizontal - 1,
		vertical:   p.vertical,
		stepsTaken: p.stepsTaken + 1,
	}
	// log.Println(string(np.from), string(np.symbol), np.stepsTaken)
	switch np.symbol {
	case pipeEW:
		tracker[np.vertical][np.horizontal] = myLib.BoolPtr(true)
		return np.west(input, tracker)
	case pipeNE:
		tracker[np.vertical][np.horizontal] = myLib.BoolPtr(true)
		return np.north(input, tracker)
	case pipeSE:
		tracker[np.vertical][np.horizontal] = myLib.BoolPtr(true)
		return np.south(input, tracker)
	case start:
		tracker[np.vertical][np.horizontal] = myLib.BoolPtr(true)
		return np.stepsTaken
	default:
		log.Println("ended west start")
		return 0
	}
}

// byteInput := []byte(input[np.vertical])
// byteInput[p.horizontal] = 'x'
// input[np.vertical] = string(byteInput)
