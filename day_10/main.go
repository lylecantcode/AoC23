package main

import (
	"aoc23/myLib"
	"log"
	"strings"
)

func main() {
	input := myLib.ErrHandledReadConv("input.txt")
	log.Printf("part one: %v\n\n", partOne(input))
	log.Printf("part two: %v\n\n", partTwo(input))
	test := myLib.ErrHandledReadConv("test_input.txt")
	testOne := partOne(test)
	testTwo := partTwo(test)
	if testOne != 8 {
		log.Fatal("incorrect response from part one test: ", testOne)
	}
	if testTwo != 2 {
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
	return 0
}

/*
1) from s, find valid pipes, depends on which direcition the pipe faces
2) follow that pipe direction to the next: s-, -7
3) track steps
*/

func partOne(input []string) int {
	var p position
	for i := 0; i < len(input); i++ {
		current := strings.IndexByte(input[i], start)
		if current != -1 {
			p = position{start, current, i, 0, start}
			break
		}
	}
	var n, s, e, w int
	n = p.north(input)
	s = p.south(input)
	e = p.east(input)
	w = p.west(input)
	log.Println(n, s, e, w)
	// total := p.move(input)
	return myLib.Biggest(n, s, e, w) / 2

}

func (p position) north(input []string) int {
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
		return np.north(input)
	case pipeSE:
		return np.east(input)
	case pipeSW:
		return np.west(input)
	case start:
		log.Println("start", np.stepsTaken)
		return np.stepsTaken
	default:
		log.Println("ended north start")
		return 0
	}
}

func (p position) south(input []string) int {
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
		return np.south(input)
	case pipeNE:
		return np.east(input)
	case pipeNW:
		return np.west(input)
	case start:
		log.Println("start", np.stepsTaken)
		return np.stepsTaken
	default:
		log.Println("ended south start")
		return 0
	}
}

func (p position) east(input []string) int {
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
		return np.east(input)
	case pipeNW:
		return np.north(input)
	case pipeSW:
		return np.south(input)
	case start:
		log.Println("start", np.stepsTaken)
		return np.stepsTaken
	default:
		log.Println("ended east start")
		return 0
	}
}

func (p position) west(input []string) int {
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
		return np.west(input)
	case pipeNE:
		return np.north(input)
	case pipeSE:
		return np.south(input)
	case start:
		log.Println("start", np.stepsTaken)
		return np.stepsTaken
	default:
		log.Println("ended west start")
		return 0
	}
}
