package mylib

import (
	"log"
	"strconv"
)

func ErrHandledAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
