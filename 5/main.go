package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jl := NewJumpList(file)
	for i := 1; true; i++ {
		if jl.Step() {
			log.Println("Part A:", i)
			break
		}
	}
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

func (jl *JumpList) Step() bool {
	tmp := jl.IPointer
	jl.IPointer += jl.Instructions[jl.IPointer]
	jl.Instructions[tmp] += 1
	return jl.IPointer > len(jl.Instructions)-1
}

func (jl *JumpList) BStep() bool {
	return jl.IPointer > len(jl.Instructions)-1
}
