package main

import (
	"reflect"
	"testing"
)

func TestCaptchaParse(t *testing.T) {
	expected := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 6, 7}
	result := parse("11223344567")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Parse failed.\n%v\n%v", expected, result)
	}

	result = parse("11223344567\n")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Parse failed on newline.\n%v\n%v", expected, result)
	}
}

func TestKnownCaptchas(t *testing.T) {
	// 1122 produces a sum of 3 (1 + 2) because the first digit (1) matches the second digit and the third digit (2)
	// matches the fourth digit.
	result := captcha([]int{1, 1, 2, 2})
	if 3 != result {
		t.Errorf("CAPTCHA Failed. Expected 3, got %d", result)
	}

	// 1111 produces 4 because each digit (all 1) matches the next.
	result = captcha([]int{1, 1, 1, 1})
	if 4 != result {
		t.Errorf("CAPTCHA Failed. Expected 4, got %d", result)
	}

	// 1234 produces 0 because no digit matches the next.
	result = captcha([]int{1, 2, 3, 4})
	if 0 != result {
		t.Errorf("CAPTCHA Failed. Expected 0, got %d", result)
	}

	// 91212129 produces 9 because the only digit that matches the next one is the last digit, 9.
	result = captcha([]int{9, 1, 2, 1, 2, 1, 2, 9})
	if 9 != result {
		t.Errorf("CAPTCHA Failed. Expected 9, got %d", result)
	}
}

func TestKnownHalfwayCaptchas(t *testing.T) {
	// 1212 produces 6: the list contains 4 items, and all four digits match the digit 2 items ahead.
	result := halfwayCaptcha([]int{1, 2, 1, 2})
	if 6 != result {
		t.Errorf("CAPTCHA Failed. Expected 6, got %d", result)
	}

	// 1221 produces 0, because every comparison is between a 1 and a 2.
	result = halfwayCaptcha([]int{1, 2, 2, 1})
	if 0 != result {
		t.Errorf("CAPTCHA Failed. Expected 0, got %d", result)
	}

	// 123425 produces 4, because both 2s match each other, but no other digit has a match.
	result = halfwayCaptcha([]int{1, 2, 3, 4, 2, 5})
	if 4 != result {
		t.Errorf("CAPTCHA Failed. Expected 4, got %d", result)
	}

	// 123123 produces 12.
	result = halfwayCaptcha([]int{1, 2, 3, 1, 2, 3})
	if 12 != result {
		t.Errorf("CAPTCHA Failed. Expected 12, got %d", result)
	}

	// 12131415 produces 4.
	result = halfwayCaptcha([]int{1, 2, 1, 3, 1, 4, 1, 5})
	if 4 != result {
		t.Errorf("CAPTCHA Failed. Expected 4, got %d", result)
	}
}
