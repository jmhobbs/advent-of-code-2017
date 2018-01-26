package main

import "testing"

func TestExample(t *testing.T) {

	// - ne,ne,ne is 3 steps away.
	hg := New()
	hg.Move("ne")
	hg.Move("ne")
	hg.Move("ne")
	if 3 != hg.DistanceFromOrigin() {
		t.Errorf("Bad distance, expected 3, got %d", hg.DistanceFromOrigin())
	}

	// - ne,ne,sw,sw is 0 steps away (back where you started).
	hg = New()
	hg.Move("ne")
	hg.Move("ne")
	hg.Move("sw")
	hg.Move("sw")
	if 0 != hg.DistanceFromOrigin() {
		t.Errorf("Bad distance, expected 0, got %d", hg.DistanceFromOrigin())
	}

	// - ne,ne,s,s is 2 steps away (se,se).
	hg = New()
	hg.Move("ne")
	hg.Move("ne")
	hg.Move("s")
	hg.Move("s")
	if 2 != hg.DistanceFromOrigin() {
		t.Errorf("Bad distance, expected 2, got %d", hg.DistanceFromOrigin())
	}

	// - se,sw,se,sw,sw is 3 steps away (s,s,sw).
	hg = New()
	hg.Move("se")
	hg.Move("sw")
	hg.Move("se")
	hg.Move("sw")
	hg.Move("sw")
	if 3 != hg.DistanceFromOrigin() {
		t.Errorf("Bad distance, expected 3, got %d", hg.DistanceFromOrigin())
	}
}

func TestAbs(t *testing.T) {
	samples := map[int]int{
		-1: 1,
		0:  0,
		1:  1,
	}

	for i, expected := range samples {
		if abs(i) != expected {
			t.Errorf("abs(%d) != %d -> %d", i, expected, abs(i))
		}
	}
}
