package main

import (
	"reflect"
	"testing"
)

func TestExample(t *testing.T) {
	// The list begins as [0] 1 2 3 4 (where square brackets indicate the current position).
	kh := New(5, []int{})
	// The first length, 3, selects ([0] 1 2) 3 4 (where parentheses indicate the sublist to be reversed).
	kh.Knot(3)
	// After reversing that section (0 1 2 into 2 1 0), we get ([2] 1 0) 3 4.
	expected := []int{2, 1, 0, 3, 4}
	if !reflect.DeepEqual(kh.List, expected) {
		t.Errorf("List in wrong order: %v, expected %v", kh.List, expected)
	}
	// Then, the current position moves forward by the length, 3, plus the skip size, 0: 2 1 0 [3] 4. Finally, the skip size increases to 1.
	if kh.Position != 3 {
		t.Errorf("Position is incorrect: %d, expected 3", kh.Position)
	}
	if kh.SkipSize != 1 {
		t.Errorf("Skip size is incorrect: %d, expected 1", kh.SkipSize)
	}
	// The second length, 4, selects a section which wraps: 2 1) 0 ([3] 4.
	kh.Knot(4)
	// The sublist 3 4 2 1 is reversed to form 1 2 4 3: 4 3) 0 ([1] 2.
	expected = []int{4, 3, 0, 1, 2}
	if !reflect.DeepEqual(kh.List, expected) {
		t.Errorf("List in wrong order: %v, expected %v", kh.List, expected)
	}
	// The current position moves forward by the length plus the skip size, a total of 5, causing it not to move because it wraps around: 4 3 0 [1] 2. The skip size increases to 2.
	if kh.Position != 3 {
		t.Errorf("Position is incorrect: %d, expected 3", kh.Position)
	}
	if kh.SkipSize != 2 {
		t.Errorf("Skip size is incorrect: %d, expected 2", kh.SkipSize)
	}
	// The third length, 1, selects a sublist of a single element, and so reversing it has no effect.
	kh.Knot(1)
	expected = []int{4, 3, 0, 1, 2}
	if !reflect.DeepEqual(kh.List, expected) {
		t.Errorf("List in wrong order: %v, expected %v", kh.List, expected)
	}
	// The current position moves forward by the length (1) plus the skip size (2): 4 [3] 0 1 2. The skip size increases to 3.
	if kh.Position != 1 {
		t.Errorf("Position is incorrect: %d, expected 1", kh.Position)
	}
	if kh.SkipSize != 3 {
		t.Errorf("Skip size is incorrect: %d, expected 3", kh.SkipSize)
	}
	// The fourth length, 5, selects every element starting with the second: 4) ([3] 0 1 2. Reversing this sublist (3 0 1 2 4 into 4 2 1 0 3) produces: 3) ([4] 2 1 0.
	kh.Knot(5)
	expected = []int{3, 4, 2, 1, 0}
	if !reflect.DeepEqual(kh.List, expected) {
		t.Errorf("List in wrong order: %v, expected %v", kh.List, expected)
	}
	// Finally, the current position moves forward by 8: 3 4 2 1 [0]. The skip size increases to 4.
	if kh.Position != 4 {
		t.Errorf("Position is incorrect: %d, expected 4", kh.Position)
	}
	if kh.SkipSize != 4 {
		t.Errorf("Skip size is incorrect: %d, expected 4", kh.SkipSize)
	}
}

func TestHash(t *testing.T) {
	samples := map[string][]int{
		// The empty string becomes a2582a3a0e66e6e86e3812dcb672a272.
		"a2582a3a0e66e6e86e3812dcb672a272": []int{17, 31, 73, 47, 23},
		// AoC 2017 becomes 33efeb34ea91902bb2f59c9920caa6cd.
		"33efeb34ea91902bb2f59c9920caa6cd": []int{int('A'), int('o'), int('C'), int(' '), int('2'), int('0'), int('1'), int('7'), 17, 31, 73, 47, 23},
		// 1,2,3 becomes 3efbe78a8d82f29979031a4aa0b16a9d.
		"3efbe78a8d82f29979031a4aa0b16a9d": []int{int('1'), int(','), int('2'), int(','), int('3'), 17, 31, 73, 47, 23},
		// 1,2,4 becomes 63960835bcdc130f0b66d7ff4f6a5a8e.
		"63960835bcdc130f0b66d7ff4f6a5a8e": []int{int('1'), int(','), int('2'), int(','), int('4'), 17, 31, 73, 47, 23},
	}

	for expected, lengths := range samples {
		kh := New(256, lengths)
		hash := kh.Hash()
		if hash != expected {
			t.Errorf("Bad hash for %v.\nExpected: '%s'\n     Got: '%s'", lengths, expected, hash)
		}
	}
}
