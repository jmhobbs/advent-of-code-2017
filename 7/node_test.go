package main

import (
	"reflect"
	"testing"
)

func TestNodeMatcher(t *testing.T) {
	matches := matcher.FindAllStringSubmatch("havc (66)", -1)
	if matches == nil {
		t.Error("Didn't match leaf.")
	}

	matches = matcher.FindAllStringSubmatch("fwft (72) -> ktlj, cntj, xhth", -1)
	if matches == nil {
		t.Error("Didn't match node.")
	}
}

func TestNewNode(t *testing.T) {
	n := NewNode("havc (66)")
	if n.Name != "havc" {
		t.Errorf("Incorrect Name: %s", n.Name)
	}

	if n.Weight != 66 {
		t.Errorf("Incorrect Weight: %d", n.Weight)
	}

	if len(n.Children) != 0 {
		t.Errorf("Incorrect Children: %v", n.Children)
	}

	n = NewNode("fwft (72) -> abc, defg")
	if n.Name != "fwft" {
		t.Errorf("Incorrect Name: %s", n.Name)
	}

	if n.Weight != 72 {
		t.Errorf("Incorrect Weight: %d", n.Weight)
	}

	if !reflect.DeepEqual(n.Children, []string{"abc", "defg"}) {
		t.Errorf("Incorrect Children: %v", n.Children)
	}
}
