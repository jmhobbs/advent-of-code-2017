package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	numbers := parse(string(input))

	log.Println("Part A", captcha(numbers))
	log.Println("Part B", halfwayCaptcha(numbers))
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

func halfwayCaptcha(input []int) int {
	acc := 0
	length := len(input)

	for i := 0; i < length; i++ {
		j := (i + length/2) % length
		if input[i] == input[j] {
			acc = acc + input[i]
		}
	}

	return acc
}
