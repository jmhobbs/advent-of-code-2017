package main

import (
	"fmt"
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

	kh := New(256)
	for _, step := range strings.Split(string(input), ",") {
		i, err := strconv.Atoi(step)
		if err != nil {
			log.Fatal(err)
		}
		kh.Knot(i)
	}

	fmt.Println("Part A:", kh.List[0]*kh.List[1])
}

type KnotHash struct {
	Size     int
	List     []int
	Position int
	SkipSize int
}

func New(size int) *KnotHash {
	kh := KnotHash{size, make([]int, size), 0, 0}
	for i := 0; i < size; i++ {
		kh.List[i] = i
	}
	return &kh
}

func (kh *KnotHash) Knot(length int) {
	// - Reverse the order of that length of elements in the list, starting with the element at the current position.
	tmp := make([]int, length)
	for i := 0; i < length; i++ {
		tmp[length-i-1] = kh.List[(kh.Position+i)%kh.Size]
	}
	for i := 0; i < length; i++ {
		kh.List[(kh.Position+i)%kh.Size] = tmp[i]
	}
	// - Move the current position forward by that length plus the skip size.
	kh.Position = (kh.Position + length + kh.SkipSize) % kh.Size
	// - Increase the skip size by one.
	kh.SkipSize += 1
}
