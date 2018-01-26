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

	lengths := []int{}

	for _, step := range strings.Split(string(input), ",") {
		i, err := strconv.Atoi(step)
		if err != nil {
			log.Fatal(err)
		}
		lengths = append(lengths, i)
	}

	kh := New(256, lengths)
	kh.Round()
	fmt.Println("Part A:", kh.List[0]*kh.List[1])

	ascii_lengths := []int{}
	for _, char := range strings.TrimSpace(string(input)) {
		ascii_lengths = append(ascii_lengths, int(char))
	}
	ascii_lengths = append(ascii_lengths, 17, 31, 73, 47, 23)

	kh = New(256, ascii_lengths)
	hash := kh.Hash()
	fmt.Println("Part B:", hash)
}

type KnotHash struct {
	Size     int
	List     []int
	Position int
	SkipSize int
	Lengths  []int
}

func New(size int, lengths []int) *KnotHash {
	kh := KnotHash{size, make([]int, size), 0, 0, lengths}
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

func (kh *KnotHash) Round() {
	for _, length := range kh.Lengths {
		kh.Knot(length)
	}
}

func (kh *KnotHash) Hash() string {
	for i := 0; i < 64; i++ {
		kh.Round()
	}

	hash := make([]int, 16)
	for i := 0; i < 16; i++ {
		block := 0
		for j := 0; j < 16; j++ {
			block = block ^ kh.List[i*16+j]
		}
		hash[i] = block
	}

	str := ""
	for _, i := range hash {
		str += fmt.Sprintf("%02x", i)
	}

	return str
}
