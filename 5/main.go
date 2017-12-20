package main

import (
	"bufio"
	"io"
	"log"
	"strconv"
)

func main() {
}

type JumpList struct {
	Instructions []int
	IPointer     int
}

func NewJumpList(input io.Reader) *JumpList {
	jl := JumpList{[]int{}, 0}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err == nil {
			jl.Instructions = append(jl.Instructions, i)
		} else {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &jl
}

func (jl *JumpList) Step() {
}
