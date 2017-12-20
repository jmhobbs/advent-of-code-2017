package main

import (
	"reflect"
	"testing"
)

func TestJumpListStep(t *testing.T) {
	jl := &JumpList{[]int{0, 3, 0, 1, -3}, 0}

	steps := [][]int{
		[]int{1, 3, 0, 1, -3},
		[]int{2, 3, 0, 1, -3},
		[]int{2, 4, 0, 1, -3},
		[]int{2, 4, 0, 1, -2},
		[]int{2, 5, 0, 1, -2},
	}

	for i, step := range steps {
		jl.Step()
		if !reflect.DeepEqual(jl.Instructions, step) {
			t.Fatalf("Bad step, %d. Expected %v got %v.", i, step, jl.Instructions)
		}
	}
}

func TestJumpListBStep(t *testing.T) {
	jl := &JumpList{[]int{0, 3, 0, 1, -3}, 0}

	for i := 0; i < 10; i++ {
		jl.BStep()
	}

	expected := []int{2, 3, 2, 3, -1}

	if !reflect.DeepEqual(jl.Instructions, expected) {
		t.Fatalf("Expected %v got %v.", expected, jl.Instructions)
	}
}
