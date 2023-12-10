package main

import (
	"aoc23/myLib"
	"fmt"
	"strings"
)

func main() {
	input := myLib.ErrHandledReadConv("input.txt")
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
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

func pipeTravel(input []string, i, j int) {
	/* valid pipe combos
	pipeNS -> pipeNS, pipeNE, pipeNW, pipeSW, pipeSE
	pipeEW -> pipeEW, pipeNE, pipeNW, pipeSW, pipeSE
	*/
	if i < len(input)-1 && input[i+1][j] == pipeNS {
		// valid to travel down
	}
	if i > 0 && input[i-1][j] == pipeNS {

	}
}

func partTwo(input []string) int {
	return 0
}

type position struct {
	symbol     byte
	horizontal int
	vertical   int
	stepsTaken int
}

func partOne(input []string) int {
	var initial position
	for i := 0; i < len(input); i++ {
		current := strings.IndexByte(input[i], start)
		if current != -1 {
			initial = position{start, current, i, 0}
		}
	}
	_ = initial
	return 0
}

/*
1) from s, find valid pipes, depends on which direcition the pipe faces
2) follow that pipe direction to the next: s-, -7
3) track steps

*/

func recursivePipeCheck(input []string, pos position) (byte, position) {
	if pos.vertical < 0 || pos.horizontal < 0 || pos.vertical >= len(input) || pos.vertical >= len(input[0]) {
		return ' ', position{}
	}
	sym, newPos := recursivePipeCheck(input, func(p position) position { p.vertical = pos.vertical - 1; return p }(pos))
	if sym != ' ' {
		switch input[newPos.vertical][pos.horizontal] {
		case pipeNS, pipeSE, pipeSW:

		}
	}
	sym, newPos = recursivePipeCheck(input, func(p position) position { p.vertical = pos.vertical + 1; return p }(pos))
	if sym != ' ' {
		switch input[newPos.vertical][pos.horizontal] {
		case pipeNS, pipeNE, pipeNW:
			// return input[newPos.vertical][newPos.horizontal], true
			validPipeCheck(input, newPos, position{})
		}
	}
	sym, newPos = recursivePipeCheck(input, func(p position) position { p.horizontal = pos.horizontal - 1; return p }(pos))
	if sym != ' ' {
		switch input[pos.vertical][newPos.horizontal] {
		case pipeEW, pipeNW, pipeSW:
			// return input[newPos.vertical][newPos.horizontal], true
			validPipeCheck(input, newPos, position{})
		}
	}
	sym, newPos = recursivePipeCheck(input, func(p position) position { p.horizontal = pos.horizontal - 1; return p }(pos))
	if sym != ' ' {
		switch input[pos.vertical][newPos.horizontal] {
		case pipeEW, pipeNE, pipeSE:
			// return input[newPos.vertical][newPos.horizontal], true
			validPipeCheck(input, newPos, position{})
		}
	}
	if input[newPos.vertical][newPos.horizontal] == start {
		// full loop
	}
	return ' ', position{}
}

func validPipeCheck(input []string, pos, newPos position) (byte, position) {
	if newPos.vertical < 0 || newPos.horizontal < 0 || newPos.vertical >= len(input) || newPos.vertical >= len(input[0]) {
		return ' ', position{}
	}

	if pos.vertical < newPos.vertical {
		switch input[newPos.vertical][pos.horizontal] {
		case pipeNS, pipeSE, pipeSW:

		}
	} else if pos.vertical > newPos.vertical {
		switch input[newPos.vertical][pos.horizontal] {
		case pipeNS, pipeNE, pipeNW:
			// return input[newPos.vertical][newPos.horizontal], true
			validPipeCheck(input, newPos, position{})
		}
	} else if pos.horizontal < newPos.horizontal {
		switch input[pos.vertical][newPos.horizontal] {
		case pipeEW, pipeNW, pipeSW:
			// return input[newPos.vertical][newPos.horizontal], true
			validPipeCheck(input, newPos, position{})
		}
	} else if pos.horizontal > newPos.horizontal {
		switch input[pos.vertical][newPos.horizontal] {
		case pipeEW, pipeNE, pipeSE:
			// return input[newPos.vertical][newPos.horizontal], true
			validPipeCheck(input, newPos, position{})
		}
	}
	if input[newPos.vertical][newPos.horizontal] == start {
		// full loop
	}
	return ' ', position{}
}
