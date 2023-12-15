package main

import (
	"aoc23/myLib"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// normal method was struggling with the linebreaks
	file, err := os.Open("./test_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	patterns := [][]string{{}}

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			patterns = append(patterns, []string{})
		} else {
			patterns[len(patterns)-1] = append(patterns[len(patterns)-1], scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println(partOne(patterns))
}

func partTwo(patterns [][]string) int {
	return 0
}

func partOne(patterns [][]string) int {
	total := 0
	for _, pattern := range patterns {
		sym := vertSymCheck(pattern)
		if sym > 0 {
			fmt.Println("vert")
			total += sym
		} else {
			fmt.Println("hor")
			total += HorSymCheck(pattern)
		}
	}
	return total
}

func vertSymCheck(pattern []string) int {
	transPattern := [][]byte{}
	for i := 0; i < len(pattern); i++ {
		transPattern = append(transPattern, []byte(pattern[i]))
	}
	transPattern = myLib.Transpose(transPattern)
	stringTransPattern := []string{}
	for i := 0; i < len(transPattern); i++ {
		stringTransPattern = append(stringTransPattern, string(transPattern[i]))
	}
	return HorSymCheck(stringTransPattern) / 100
}

func HorSymCheck(pattern []string) int {
	smudged := -1
	symLines := []int{}
	for i := 0; i < len(pattern)-1; i++ {
		compare := len(sliceCompare([]byte(pattern[i]), []byte(pattern[i+1])))
		if compare == 1 && smudged == -1 {
			symLines = append(symLines, i)
			smudged = i
		} else if compare == 0 {
			symLines = append(symLines, i)
		}
	}
	log.Println("sym lines", symLines)
	for k := 0; k < len(symLines); k++ {
		sym := true
		line := symLines[k]
		var compLineOne, compLineTwo int
		for l := 1; l < len(pattern)/2 && sym; l++ {
			compLineOne, compLineTwo = line-l, line+1+l
			if compLineOne < 0 || compLineTwo >= len(pattern) {
				break
			}
			compare := len(sliceCompare([]byte(pattern[compLineOne]), []byte(pattern[compLineTwo])))
			if compare == 0 {
				fmt.Println("match")
			} else if compare == 1 && smudged == -1 {
				smudged = l
			} else {
				sym = false
			}
		}
		log.Println(smudged, compLineOne, compLineTwo, sym)
		if sym && smudged >= 0 && compLineOne <= smudged && compLineTwo >= smudged {
			return (line + 1) * 100
		}
	}
	fmt.Println(smudged)
	fmt.Println("no matches", pattern)
	return -1
}

func sliceCompare[S ~[]E, E comparable](s1, s2 S) S {
	var dif S
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			dif = append(dif, s1[i])
		}
	}
	fmt.Println(dif)
	return dif
}
