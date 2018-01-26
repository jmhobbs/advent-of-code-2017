package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	hg := New()
	for _, direction := range strings.Split(strings.TrimSpace(string(input)), ",") {
		hg.Move(direction)
	}
	log.Println("Part A:", hg.DistanceFromOrigin())
}

type HexGrid struct {
	Q int
	R int
}

func New() *HexGrid {
	return &HexGrid{0, 0}
}

func (hg *HexGrid) Move(direction string) {
	switch direction {
	case "n":
		hg.R -= 1
	case "ne":
		hg.R -= 1
		hg.Q += 1
	case "se":
		hg.Q += 1
	case "s":
		hg.R += 1
	case "sw":
		hg.R += 1
		hg.Q -= 1
	case "nw":
		hg.Q -= 1
	default:
		log.Fatalf("Invalid direction: %s", direction)
	}
	return
}

func (hg *HexGrid) DistanceFromOrigin() int {
	return (abs(0-hg.Q) + abs(0-hg.Q-hg.R) + abs(0-hg.R)) / 2
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
