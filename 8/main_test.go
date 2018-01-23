package main

import "testing"

func TestParseInstruction(t *testing.T) {
	samples := map[string]Instruction{
		"b inc 5 if a > 1":     Instruction{"b", true, 5, "a", GREATER_THAN, 1},
		"a inc 1 if b < 5":     Instruction{"a", true, 1, "b", LESS_THAN, 5},
		"c dec -10 if a >= 1":  Instruction{"c", false, -10, "a", GREATER_THAN_EQUAL, 1},
		"c inc -20 if c == 10": Instruction{"c", true, -20, "c", EQUAL, 10},
	}

	for sample, expected := range samples {
		result, err := ParseInstruction(sample)
		if err != nil {
			t.Errorf("Error parsing '%s': %s", sample, err)
		} else if result != expected {
			t.Errorf("Parse error on '%s'.  Got %v, expected %v.", sample, result, expected)
		}
	}
}

func TestExample(t *testing.T) {
	r := make(Registers)
	// - Because a starts at 0, it is not greater than 1, and so b is not modified.
	r.Execute(Instruction{"b", true, 5, "a", GREATER_THAN, 1})
	if r["b"] != 0 {
		t.Fatal("Error executing instruction!")
	}
	// - a is increased by 1 (to 1) because b is less than 5 (it is 0).
	r.Execute(Instruction{"a", true, 1, "b", LESS_THAN, 5})
	if r["a"] != 1 {
		t.Fatal("Error executing instruction!")
	}
	// - c is decreased by -10 (to 10) because a is now greater than or equal to 1 (it is 1).
	r.Execute(Instruction{"c", false, -10, "a", GREATER_THAN_EQUAL, 1})
	if r["c"] != 10 {
		t.Fatal("Error executing instruction!")
	}
	// - c is increased by -20 (to -10) because c is equal to 10.
	r.Execute(Instruction{"c", true, -20, "c", EQUAL, 10})
	if r["c"] != -10 {
		t.Fatal("Error executing instruction!")
	}
}
