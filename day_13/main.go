package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// normal method was struggling with the linebreaks
	file, err := os.Open("./input.txt")
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
			total += sym
		} else {
			total += HorSymCheck(pattern)
		}
	}
	return total
}

func vertSymCheck(pattern []string) int {
	return HorSymCheck(transpose(pattern)) / 100
}

func transpose(input []string) []string {
	byteSlice := make([][]byte, len(input[0]))
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			byteSlice[j] = append(byteSlice[j], input[i][j])
		}
	}
	output := []string{}
	for i := 0; i < len(byteSlice); i++ {
		output = append(output, string(byteSlice[i]))
	}
	return output
}

func HorSymCheck(pattern []string) int {
	symLines := []int{}
	for i := 0; i < len(pattern)-1; i++ {
		if strings.Compare(pattern[i], pattern[i+1]) == 0 {
			symLines = append(symLines, i)
		}
	}
	for k := 0; k < len(symLines); k++ {
		line := symLines[k]
		sym := true
		for l := 0; l < len(pattern)/2 && sym; l++ {
			if line-l < 0 || line+1+l >= len(pattern) {
				break
			}
			if strings.Compare(pattern[line-l], pattern[line+l+1]) == 0 {
				// fmt.Println("match")
			} else {
				sym = false
			}
		}
		if sym {
			return (line + 1) * 100
		}
	}
	fmt.Println("no H matches", pattern)
	return -1
}
