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

func TestChecksum(t *testing.T) {
	s := Spreadsheet{
		Row{5, 1, 9, 5},
		Row{7, 5, 3},
		Row{2, 4, 6, 8},
	}
	// In this example, the spreadsheet's checksum would be 8 + 4 + 6 = 18.
	cs := checksum(s)
	if 18 != cs {
		t.Errorf("Bad checksum. Expected 18, got %v", cs)
	}
}

func TestRowChecksum(t *testing.T) {
	// The first row's largest and smallest values are 9 and 1, and their difference is 8.
	result := checksumRow(Row{5, 1, 9, 5})
	if 8 != result {
		t.Errorf("Bad checksum. Expected 8, got %v", result)
	}

	// The second row's largest and smallest values are 7 and 3, and their difference is 4.
	result = checksumRow(Row{7, 5, 3})
	if 4 != result {
		t.Errorf("Bad checksum. Expected 4, got %v", result)
	}

	// The third row's difference is 6.
	result = checksumRow(Row{2, 4, 6, 8})
	if 6 != result {
		t.Errorf("Bad checksum. Expected 6, got %v", result)
	}
}

func TestEvenRowChecksum(t *testing.T) {
	s := Spreadsheet{
		Row{5, 9, 2, 8},
		Row{9, 4, 7, 3},
		Row{3, 8, 6, 5},
	}

	// In the first row, the only two numbers that evenly divide are 8 and 2; the result of this division is 4.
	result := evenRowChecksum(s[0])
	if 4 != result {
		t.Errorf("Bad checksum. Expected 4, got %v", result)
	}

	// In the second row, the two numbers are 9 and 3; the result is 3.
	result = evenRowChecksum(s[1])
	if 3 != result {
		t.Errorf("Bad checksum. Expected 3, got %v", result)
	}

	// In the third row, the result is 2.
	result = evenRowChecksum(s[2])
	if 2 != result {
		t.Errorf("Bad checksum. Expected 2, got %v", result)
	}
}

func TestEvenChecksum(t *testing.T) {
	s := Spreadsheet{
		Row{5, 9, 2, 8},
		Row{9, 4, 7, 3},
		Row{3, 8, 6, 5},
	}

	// In this example, the sum of the results would be 4 + 3 + 2 = 9.
	result := evenChecksum(s)
	if 9 != result {
		t.Errorf("Bad checksum. Expected 9, got %v", result)
	}
}
