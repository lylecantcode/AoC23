package main

import (
	"aoc23/myLib"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	input := myLib.ErrHandledReadConv("input.txt")
	input = input[:len(input)-1]
	log.Printf("part two: %v\n\n", partTwo(input))

	test := myLib.ErrHandledReadConv("input.txt")
	test = test[:len(test)-1]
	testTwo := partTwo(test)
	if testTwo != 353 {
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
	stepsTaken float64
}

func partTwo(input []string) int {
	// establish bounaries, iterate through and check
	tracker := [][]float64{}
	for i := 0; i < len(input); i++ {
		tracker = append(tracker, make([]float64, len(input[i])))
	}

	var p position
	for i := 0; i < len(input); i++ {
		current := strings.IndexByte(input[i], start)
		if current != -1 {
			p = position{start, current, i, 0}
			break
		}
	}
	var n, s, e, w float64
	n = p.north(input, tracker)
	s = p.south(input, tracker)
	e = p.east(input, tracker)
	w = p.west(input, tracker)
	log.Println(n, s, e, w)

	for i := 0; i < len(tracker); i++ {
		for j := 0; j < len(tracker[i]); j++ {
			if tracker[i][j] == 0 {
				propogate(input, tracker, i, j)
			}
		}
	}

	/* forgot about squeezing between pipes
	check around each "false", can check for above if east+west == 0
	*/
	count := 0
	for i := 0; i < len(tracker); i++ {
		for j := 0; j < len(tracker[i]); j++ {
			if tracker[i][j] == 0 {
				if !isGapNS(tracker, i, j, -1, 1) && !isGapNS(tracker, i, j, -1, -1) && !isGapNS(tracker, i, j, 1, 1) && !isGapNS(tracker, i, j, 1, -1) &&
					!isGapEW(tracker, i, j, -1, 1) && !isGapEW(tracker, i, j, -1, -1) && !isGapEW(tracker, i, j, 1, 1) && !isGapEW(tracker, i, j, 1, -1) {
					// 			count++
					// 			fmt.Printf("[.%3v.]", count)
				} else {
					tracker[i][j] = -1
					// fmt.Printf("[     ]")
					// }
					// 		if tracker[i][j] == 0 {
					// 		}

					// 	} else if tracker[i][j] > 0 {
					// 		fmt.Printf("[%5v]", tracker[i][j])
					// 	} else {
					// 		fmt.Printf("%7v", " ")
					// 	}
					// }
					// fmt.Printf("\n")
				}
			}

		}
	}
	for i := 0; i < len(tracker); i++ {
		for j := 0; j < len(tracker[i]); j++ {
			if tracker[i][j] == 0 {
				if !isGapNS(tracker, i, j, -1, 1) && !isGapNS(tracker, i, j, -1, -1) && !isGapNS(tracker, i, j, 1, 1) && !isGapNS(tracker, i, j, 1, -1) &&
					!isGapEW(tracker, i, j, -1, 1) && !isGapEW(tracker, i, j, -1, -1) && !isGapEW(tracker, i, j, 1, 1) && !isGapEW(tracker, i, j, 1, -1) {
					count++
					fmt.Printf("[.....]")

				} else if tracker[i][j] > 0 {
					fmt.Printf("[%5v]", tracker[i][j])
				} else {
					fmt.Printf("%7v", " ")

				}
				fmt.Printf("\n")

			}

		}
	}

	return count

}

func isGapNS(tracker [][]float64, i, j, iMod, jMod int) bool {
	if tracker[i][j] < 0 {
		return true
	}
	if i > 0 && j > 0 && i < len(tracker)-1 && j < len(tracker[i])-1 && math.Abs(float64(tracker[i+iMod][j]-tracker[i+iMod][j+jMod])) != 1 {
		return isGapNS(tracker, i+iMod, j, iMod, jMod)
	}

	return false
}

func isGapEW(tracker [][]float64, i, j, iMod, jMod int) bool {
	if tracker[i][j] < 0 {
		return true
	}
	x := float64(tracker[i][j+jMod] - tracker[i+iMod][j+jMod])
	if i > 0 && j > 0 && i < len(tracker)-1 && j < len(tracker[i])-1 && math.Abs(x) != 1 {
		log.Println("%v = %v - %v", x, tracker[i][j+jMod], tracker[i+iMod][j+jMod])
		return isGapNS(tracker, i, j+jMod, iMod, jMod)
	}

	return false
}

func propogate(input []string, tracker [][]float64, i, j int) {
	if i < 0 || j < 0 || i > len(tracker)-1 || j > len(tracker[i])-1 {
		return
	}
	if tracker[i][j] == 0 && (i == 0 || j == 0 || i == len(tracker)-1 || j == len(tracker[i])-1 ||
		tracker[i-1][j] == -1 || tracker[i+1][j] == -1 ||
		tracker[i][j-1] == -1 || tracker[i][j+1] == -1) {

		tracker[i][j] = -1

		propogate(input, tracker, i+1, j)
		propogate(input, tracker, i-1, j)
		propogate(input, tracker, i, j+1)
		propogate(input, tracker, i, j-1)
	}
}

func (p position) north(input []string, tracker [][]float64) float64 {
	if p.vertical < 1 {
		return 0
	}
	np := position{
		symbol:     input[p.vertical-1][p.horizontal],
		horizontal: p.horizontal,
		vertical:   p.vertical - 1,
		stepsTaken: p.stepsTaken + 1,
	}
	// log.Println(string(np.), string(np.symbol), np.stepsTaken)
	switch np.symbol {
	case pipeNS:
		tracker[np.vertical][np.horizontal] = np.stepsTaken
		return np.north(input, tracker)
	case pipeSE:
		tracker[np.vertical][np.horizontal] = np.stepsTaken
		return np.east(input, tracker)
	case pipeSW:
		tracker[np.vertical][np.horizontal] = np.stepsTaken
		return np.west(input, tracker)
	case start:
		tracker[np.vertical][np.horizontal] = np.stepsTaken
		return np.stepsTaken
	default:
		// log.Println("ended north start")
		return 0
	}
}

func (p position) south(input []string, tracker [][]float64) float64 {
	if p.vertical >= len(input)-1 {
		return 0
	}
	np := position{
		symbol:     input[p.vertical+1][p.horizontal],
		horizontal: p.horizontal,
		vertical:   p.vertical + 1,
		stepsTaken: p.stepsTaken + 1,
	}
	// log.Println(string(np.), string(np.symbol), np.stepsTaken)
	switch np.symbol {
	case pipeNS:
		tracker[np.vertical][np.horizontal] = np.stepsTaken
		return np.south(input, tracker)
	case pipeNE:
		tracker[np.vertical][np.horizontal] = np.stepsTaken
		return np.east(input, tracker)
	case pipeNW:
		tracker[np.vertical][np.horizontal] = np.stepsTaken
		return np.west(input, tracker)
	case start:
		tracker[np.vertical][np.horizontal] = np.stepsTaken
		return np.stepsTaken
	default:
		// log.Println("ended south start")
		return 0
	}
}

func (p position) east(input []string, tracker [][]float64) float64 {
	if p.horizontal >= len(input[0])-1 {
		return 0
	}
	np := position{
		symbol:     input[p.vertical][p.horizontal+1],
		horizontal: p.horizontal + 1,
		vertical:   p.vertical,
		stepsTaken: p.stepsTaken + 1,
	}
	// log.Println(string(np.), string(np.symbol), np.stepsTaken)
	switch np.symbol {
	case pipeEW:
		tracker[np.vertical][np.horizontal] = np.stepsTaken
		return np.east(input, tracker)
	case pipeNW:
		tracker[np.vertical][np.horizontal] = np.stepsTaken
		return np.north(input, tracker)
	case pipeSW:
		tracker[np.vertical][np.horizontal] = np.stepsTaken
		return np.south(input, tracker)
	case start:
		tracker[np.vertical][np.horizontal] = np.stepsTaken
		return np.stepsTaken
	default:
		// log.Println("ended east start")
		return 0
	}
}

func (p position) west(input []string, tracker [][]float64) float64 {
	if p.horizontal < 1 {
		return 0
	}
	np := position{
		symbol:     input[p.vertical][p.horizontal-1],
		horizontal: p.horizontal - 1,
		vertical:   p.vertical,
		stepsTaken: p.stepsTaken + 1,
	}
	// log.Println(string(np.), string(np.symbol), np.stepsTaken)
	switch np.symbol {
	case pipeEW:
		tracker[np.vertical][np.horizontal] = np.stepsTaken
		return np.west(input, tracker)
	case pipeNE:
		tracker[np.vertical][np.horizontal] = np.stepsTaken
		return np.north(input, tracker)
	case pipeSE:
		tracker[np.vertical][np.horizontal] = np.stepsTaken
		return np.south(input, tracker)
	case start:
		tracker[np.vertical][np.horizontal] = np.stepsTaken
		return np.stepsTaken
	default:
		// log.Println("ended west start")
		return 0
	}
}
