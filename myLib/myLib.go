package myLib

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ErrHandledAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func ErrHandledReadConv(s string) []string {
	a, err := os.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(a), "\n")
}

func ErrHandledRead(s string) []byte {
	a, err := os.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	return a
}

func CheckSurroundings(i, j int, slice [][]*int) *int {
	if i < 0 || j < 0 || i >= len(slice) || j >= len(slice[i]) || slice[i][j] == nil {
		return nil
	}
	return slice[i][j]
}

func IntPtr(val int) *int {
	return &val
}

/*
// make a generic err handler?

type ErrHandled[T any] interface {
	ErrHandler(T) T
}

type StringConversion struct {
	Atoi func(string) (int, error)
}

func (f StringConversion) ErrHandler(s string) int {
	a, err := f.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return a
}

type ReadFile struct {
	toInput func(string) ([]byte, error)
}

func (f ReadFile) ErrHandler(s string) []string {
	a, err := f.toInput(s)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(a), "\n")
}

// demo of use:
func main() {
	atoi := StringConversion{Atoi: strconv.Atoi}
	x := atoi.ErrHandler("5")
	log.Println(x)

	read := ReadFile{toInput: os.ReadFile}
	input := read.ErrHandler("demo_input.txt")
	log.Println(input)
}
*/
