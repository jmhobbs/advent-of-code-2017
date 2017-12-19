package main

import "testing"

func TestKnownMoves(t *testing.T) {
	fixtures := map[int]int{
		// Data from square 1 is carried 0 steps, since it's at the access port.
		1: 0,
		// Data from square 12 is carried 3 steps, such as: down, left, left.
		12: 3,
		// Data from square 23 is carried only 2 steps: up twice.
		23: 2,
		//Data from square 1024 must be carried 31 steps.
		1024: 31,
	}

	for v, expected := range fixtures {
		count := countSteps(v)
		if count != expected {
			t.Errorf("Expected %d for %d got %d", expected, v, count)
		}
	}
}
