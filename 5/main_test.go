package main

import (
	"reflect"
	"testing"
)

func TestJumpListStep(t *testing.T) {
	jl := &JumpList{[]int{0, 3, 0, 1, -3}, 0}

	steps := [][]int{
		[]int{0, 3, 0, 1, -3},
		[]int{1, 3, 0, 1},
		[]int{2, 3, 0, 1},
		[]int{2, 4, 0, 1},
		[]int{2, 4, 0, 1},
		[]int{2, 5, 0, 1},
	}

	for i, step := range steps {
		jl.Step()
		if !reflect.DeepEqual(jl.Instructions, step) {
			t.Fatalf("Bad step, %d. Expected %v got %v.", i, step, jl.Instructions)
		}
	}
}
