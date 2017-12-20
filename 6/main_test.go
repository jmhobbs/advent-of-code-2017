package main

import (
	"reflect"
	"testing"
)

func TestFindLargestBlock(t *testing.T) {
	// Largest in any order
	m := Memory{1, 2, 5, 4, 3}
	idx := FindLargestBlock(m)
	if idx != 2 {
		t.Errorf("Incorrect largest block. Expected 2 got %d", idx)
	}

	// Largest, preferring lower index values
	m = Memory{1, 3, 5, 4, 5}
	idx = FindLargestBlock(m)
	if idx != 2 {
		t.Errorf("Incorrect largest block. Expected 2 got %d", idx)
	}
}

func TestReallocate(t *testing.T) {
	//  - The banks start with 0, 2, 7, and 0 blocks. The third bank has the most blocks, so it is chosen for redistribution.
	m := Memory{0, 2, 7, 0}

	//  - Starting with the next bank (the fourth bank) and then continuing to the first bank, the second bank, and so on, the 7 blocks are spread out over the memory banks. The fourth, first, and second banks get two blocks each, and the third bank gets one back. The final result looks like this: 2 4 1 2.
	m = Reallocate(m)
	expect := Memory{2, 4, 1, 2}
	if !reflect.DeepEqual(m, expect) {
		t.Fatalf("Bad reallocation. Expected %v, got %v", expect, m)
	}

	//  - Next, the second bank is chosen because it contains the most blocks (four). Because there are four memory banks, each gets one block. The result is: 3 1 2 3.
	m = Reallocate(m)
	expect = Memory{3, 1, 2, 3}
	if !reflect.DeepEqual(m, expect) {
		t.Fatalf("Bad reallocation. Expected %v, got %v", expect, m)
	}

	//  - Now, there is a tie between the first and fourth memory banks, both of which have three blocks. The first bank wins the tie, and its three blocks are distributed evenly over the other three banks, leaving it with none: 0 2 3 4.
	m = Reallocate(m)
	expect = Memory{0, 2, 3, 4}
	if !reflect.DeepEqual(m, expect) {
		t.Fatalf("Bad reallocation. Expected %v, got %v", expect, m)
	}

	//  - The fourth bank is chosen, and its four blocks are distributed such that each of the four banks receives one: 1 3 4 1.
	m = Reallocate(m)
	expect = Memory{1, 3, 4, 1}
	if !reflect.DeepEqual(m, expect) {
		t.Fatalf("Bad reallocation. Expected %v, got %v", expect, m)
	}

	//  - The third bank is chosen, and the same thing happens: 2 4 1 2.
	m = Reallocate(m)
	expect = Memory{2, 4, 1, 2}
	if !reflect.DeepEqual(m, expect) {
		t.Fatalf("Bad reallocation. Expected %v, got %v", expect, m)
	}
}
