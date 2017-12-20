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
	type Test struct {
		Description string
		Name        string
		Weight      int
		Children    []string
	}

	tests := []Test{
		Test{
			"havc (66)",
			"havc",
			66,
			[]string{},
		},
		Test{
			"fwft (72) -> abc, defg",
			"fwft",
			72,
			[]string{"abc", "defg"},
		},
	}

	for _, test := range tests {
		n, err := NewNode(test.Description)
		if err != nil {
			t.Errorf("Error creating %s: %s", test.Description, err)
		}

		if n.Name != test.Name {
			t.Errorf("Incorrect name for \"%s\". Expected %s, got %s", test.Description, test.Name, n.Name)
		}

		if n.Weight != test.Weight {
			t.Errorf("Incorrect weight for \"%s\". Expected %d, got %d", test.Description, test.Weight, n.Weight)
		}

		if !reflect.DeepEqual(n.Children, test.Children) {
			t.Errorf("Incorrect children for \"%s\". Expected %v, got %v", test.Description, test.Children, n.Children)
		}
	}
}

func TestFindRootNode(t *testing.T) {
	nodes := []*Node{
		&Node{"pbga", 66, []string{}},
		&Node{"xhth", 57, []string{}},
		&Node{"ebii", 61, []string{}},
		&Node{"havc", 66, []string{}},
		&Node{"ktlj", 57, []string{}},
		&Node{"fwft", 72, []string{"ktlj", "cntj", "xhth"}},
		&Node{"qoyq", 66, []string{}},
		&Node{"padx", 45, []string{"pbga", "havc", "qoyq"}},
		&Node{"tknk", 41, []string{"ugml", "padx", "fwft"}},
		&Node{"jptl", 61, []string{}},
		&Node{"ugml", 68, []string{"gyxo", "ebii", "jptl"}},
		&Node{"gyxo", 61, []string{}},
		&Node{"cntj", 57, []string{}},
	}

	root_node := FindRootNode(nodes)
	if root_node.Name != "tknk" {
		t.Errorf("Wrong root node. Expected tknk, got %s", root_node.Name)
	}
}
