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

func TestSpiralLocation(t *testing.T) {
	/*
	   37  36  35  34  33  32  31
	   38  17  16  15  14  13  30
	   39  18   5   4   3  12  29
	   40  19   6   1   2  11  28
	   41  20   7   8   9  10  27
	   42  21  22  23  24  25  26
	   43  44  45  46  47  48  49
	*/

	expected := [][]int{
		[]int{1, 0, 0},
		[]int{2, 1, 0},
		[]int{3, 1, 1},
		[]int{4, 0, 1},
		[]int{5, -1, 1},
		[]int{6, -1, 0},
		[]int{7, -1, -1},
		[]int{8, 0, -1},
		[]int{9, 1, -1},
		[]int{10, 2, -1},
		[]int{11, 2, 0},
		[]int{12, 2, 1},
		[]int{13, 2, 2},
		[]int{14, 1, 2},
	}

	for _, expect := range expected {
		x, y := getSpiralLocation(expect[0])
		if x != expect[1] || y != expect[2] {
			t.Errorf("Wrong location for %d. Expected %d,%d got %d,%d", expect[0], expect[1], expect[2], x, y)
		}
	}
}

func TestSpiralValue(t *testing.T) {
	/*
		147  142  133  122   59
		304    5    4    2   57
		330   10    1    1   54
		351   11   23   25   26
		362  747  806--->   ...
	*/
	expected := map[int]int{
		1:  1,
		2:  1,
		3:  2,
		4:  4,
		5:  5,
		6:  10,
		7:  11,
		8:  23,
		9:  25,
		10: 26,
		11: 54,
		12: 57,
		13: 59,
		14: 122,
		15: 133,
		16: 142,
		17: 147,
		18: 304,
		19: 330,
		20: 351,
		21: 362,
		22: 747,
		23: 806,
	}

	for src, expect := range expected {
		v := getSpiralValue(src)
		if v != expect {
			t.Errorf("Wrong value for %d. Expected %d got %d", src, expect, v)
		}
	}
}
