package main

import (
	"io"
)

func main() {
}

type Group struct {
	Parent   *Group
	Children []*Group
	Score    int
}

func ProcessStream(r io.ByteReader) Group {
	return Group{nil, []*Group{}, 0}
}

func (g Group) Count() int {
	return 0
}

func (g Group) TotalScore() int {
	return 0
}
