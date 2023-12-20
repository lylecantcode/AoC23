package main

import (
	"aoc23/myLib"
	"flag"
	"fmt"
	"log"
)

func main() {
	testFlag := flag.Bool("t", false, "controls which input to use")
	flag.Parse()
	inputFile := "input.txt"
	if *testFlag {
		inputFile = "test_input.txt"
	}
	input := myLib.ErrHandledReadConv(inputFile)
	input = input[:len(input)-1]
	log.Println(partOne(input))
	log.Println(partTwo(input))
}

func partTwo(input []string) int {
	return 0
}

type beam struct {
	dir  byte
	i, j int
}

func partOne(input []string) int {
	energised := make([][]map[byte]interface{}, len(input))
	width := len(input[0])
	for i := range input {
		energised[i] = make([]map[byte]interface{}, width)
	}
	b := &beam{'e', 0, 0}

	// bool array to track energised squares?
	// need to track beam position to know which way mirrors reflect
	b.travel(input, energised)

	count := 0
	for i := 0; i < len(energised); i++ {
		fmt.Println(energised[i], "\n", input[i])
		for j := 0; j < len(energised[i]); j++ {
			if len(energised[i][j]) != 0 {
				count++
			}
		}
	}
	return count
}

func (b *beam) travel(input []string, tracker [][]map[byte]interface{}) {
	if b.i < 0 || b.j < 0 || b.i >= len(input) || b.j >= len(input[0]) {
		return
	}
	_, exists := tracker[b.i][b.j][b.dir]
	if exists {
		return
	}

	tracker[b.i][b.j] = map[byte]interface{}{
		b.dir: nil,
	}

	current := input[b.i][b.j]
	if current == '/' || current == '\\' {
		b.reflect(current)
	}
	if (b.dir == 'e' || b.dir == 'w') && current == '|' || (b.dir == 'n' || b.dir == 's') && current == '-' {
		bNew := b.split(current)
		bNew.move()
		bNew.travel(input, tracker)
	}
	b.move()
	b.travel(input, tracker)
}

func (b *beam) split(splitter byte) beam {
	if splitter == '-' {
		b.dir = 'w'
		return beam{'e', b.i, b.j}
	}
	// if splitter == '|' {
	b.dir = 'n'
	return beam{'s', b.i, b.j}
	// }
}

func (b *beam) reflect(mirror byte) {
	if mirror == '/' {
		switch b.dir {
		case 'n':
			b.dir = 'e'
		case 'e':
			b.dir = 'n'
		case 's':
			b.dir = 'w'
		case 'w':
			b.dir = 's'
		}
	}
	if mirror == '\\' {
		switch b.dir {
		case 'n':
			b.dir = 'w'
		case 'w':
			b.dir = 'n'
		case 's':
			b.dir = 'e'
		case 'e':
			b.dir = 's'
		}
	}
}

func (b *beam) move() {
	switch b.dir {
	case 'n':
		b.i--
	case 'e':
		b.j++
	case 's':
		b.i++
	case 'w':
		b.j--
	}
}
