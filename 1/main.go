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
	acc := 0

	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			acc = acc + input[i]
		}
	}

	if input[len(input)-1] == input[0] {
		acc = acc + input[0]
	}

	return acc
}
