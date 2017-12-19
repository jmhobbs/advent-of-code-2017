package main

import "log"

func main() {
	log.Println("Part A:", countSteps(265149))
}

const (
	RIGHT = iota
	UP
	LEFT
	DOWN
)

func countSteps(src int) int {
	x, y := getSpiralLocation(src)

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

func getSpiralLocation(src int) (x, y int) {
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
				continue
			}
		}

		if direction == UP {
			if y == w {
				direction = LEFT
			} else {
				y = y + 1
				continue
			}
		}

		if direction == LEFT {
			if x == -1*w {
				direction = DOWN
			} else {
				x = x - 1
				continue
			}
		}

		if direction == DOWN {
			if y == -1*w {
				direction = RIGHT
			} else {
				y = y - 1
				continue
			}
		}

		if direction == RIGHT {
			x = x + 1
		}
	}

	return
}
