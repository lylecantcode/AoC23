package main

import (
	"aoc23/myLib"
	"log"
)

func main() {
	input := myLib.ErrHandledReadConv("input.txt")
	log.Println(partOne(input))
	log.Println(partTwo(input))
	test := myLib.ErrHandledReadConv("test_input.txt")
	if partOne(test) != 114 {
		log.Fatal("incorrect response from test input")
	}
	if partTwo(test) != 2 {
		log.Fatal("incorrect response from test input", partTwo(test))
	}

}

func partTwo(input []string) int {
	return 0
}

func partOne(input []string) int {
	return 0
}
