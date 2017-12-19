package main

import (
	"log"
)

func main() {
	log.Println("Part A:", countSteps(265149))
	for i := 2; true; i++ {
		v := getSpiralValue(i)
		if v > 265149 {
			log.Println("Part B:", v)
			break
		}
	}
}

const (
	RIGHT = iota
	UP
	LEFT
	DOWN
)

type spiralCallback interface {
	callback(int, int, int)
}

type noopSpiralCallback struct{}

func (nsc *noopSpiralCallback) callback(int, int, int) {}

func countSteps(src int) int {
	x, y := getSpiralLocation(src, &noopSpiralCallback{})

	total := 0
	if x < 0 {
		total = total + -1*x
	} else {
		total = total + x
	}

	if y < 0 {
		total = total + -1*y
	} else {
		total = total + y
	}

	return total
}

func getSpiralLocation(src int, cb spiralCallback) (x, y int) {
	x = -1
	y = 0
	w := 0
	direction := RIGHT

	for i := 1; i <= src; i++ {

		if x == w && y == -1*w {
			w = w + 1
		}

		if direction == RIGHT {
			if x == w {
				direction = UP
			} else {
				x = x + 1
				cb.callback(i, x, y)
				continue
			}
		}

		if direction == UP {
			if y == w {
				direction = LEFT
			} else {
				y = y + 1
				cb.callback(i, x, y)
				continue
			}
		}

		if direction == LEFT {
			if x == -1*w {
				direction = DOWN
			} else {
				x = x - 1
				cb.callback(i, x, y)
				continue
			}
		}

		if direction == DOWN {
			if y == -1*w {
				direction = RIGHT
			} else {
				y = y - 1
				cb.callback(i, x, y)
				continue
			}
		}

		if direction == RIGHT {
			x = x + 1
			cb.callback(i, x, y)
		}
	}

	return
}

func getSpiralValue(src int) int {
	sva := NewSpiralValueAccumulator()
	x, y := getSpiralLocation(src, &sva)
	return sva.Values[Coordinates{x, y}]
}

type Coordinates struct {
	X int
	Y int
}

type spiralValueAccumulator struct {
	Values map[Coordinates]int
}

func NewSpiralValueAccumulator() spiralValueAccumulator {
	sva := spiralValueAccumulator{make(map[Coordinates]int)}
	// special case
	sva.Values[Coordinates{0, 0}] = 1
	return sva
}

func (sva *spiralValueAccumulator) callback(i, x, y int) {
	// special case
	if i == 1 {
		return
	}

	total := 0
	total += sva.Values[Coordinates{x + 1, y}]
	total += sva.Values[Coordinates{x + 1, y + 1}]
	total += sva.Values[Coordinates{x, y + 1}]
	total += sva.Values[Coordinates{x - 1, y + 1}]
	total += sva.Values[Coordinates{x - 1, y}]
	total += sva.Values[Coordinates{x - 1, y - 1}]
	total += sva.Values[Coordinates{x, y - 1}]
	total += sva.Values[Coordinates{x + 1, y - 1}]

	sva.Values[Coordinates{x, y}] = total
}
