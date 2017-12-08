package main

import (
	"strconv"
	"strings"
)

func main() {
}

func parse(input string) []int {
	ints := []int{}
	for _, c := range strings.Split(input, "") {
		i, err := strconv.Atoi(c)
		if err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}

func captcha(input []int) int {
	return 0
}
