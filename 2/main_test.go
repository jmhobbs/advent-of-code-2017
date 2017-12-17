package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	expected := Spreadsheet{
		Row{1, 2, 3},
		Row{6, 6, 9},
		Row{4, 2, 1},
	}
	result := parse(strings.NewReader("1\t2\t3\n6\t6\t9\n4\t2\t1\n"))
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Parse failed.\n%v\n%v", expected, result)
	}
}
