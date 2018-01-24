package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func main() {
	contents, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	root := ProcessStream(bytes.NewBuffer(contents))
	fmt.Println("Part A:", root.TotalScore())
	fmt.Println("Part B:", root.TotalGarbageCount())
}

type Group struct {
	Parent       *Group
	Children     []*Group
	Score        int
	GarbageCount int
}

func ProcessStream(r io.ByteReader) Group {
	// { - Opens group
	// } - Closes group
	// < - Opens garbage
	// > - Closes garbage
	// ! - Skip next char (in garbage)

	root_group := Group{nil, []*Group{}, 0, 0}
	current_group := &root_group
	skip_char := false
	garbage_opened := false

	for {
		b, err := r.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		if garbage_opened {
			if skip_char {
				skip_char = false
			} else if b == '>' {
				garbage_opened = false
			} else if b == '!' {
				skip_char = true
			} else {
				current_group.GarbageCount += 1
			}
			continue
		}

		if b == '<' {
			garbage_opened = true
		}

		if b == '{' {
			new_group := Group{current_group, []*Group{}, current_group.Score + 1, 0}
			current_group.Children = append(current_group.Children, &new_group)
			current_group = &new_group
		}

		if b == '}' {
			current_group = current_group.Parent
		}
	}

	return root_group
}

func (g Group) Count() int {
	total := len(g.Children)
	for _, g := range g.Children {
		total += g.Count()
	}
	return total
}

func (g Group) TotalScore() int {
	total := g.Score
	for _, g := range g.Children {
		total += g.TotalScore()
	}
	return total
}

func (g Group) TotalGarbageCount() int {
	total := g.GarbageCount
	for _, g := range g.Children {
		total += g.TotalGarbageCount()
	}
	return total
}
