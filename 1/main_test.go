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
}

func TestKnownCaptchas(t *testing.T) {
	// 1122 produces a sum of 3 (1 + 2) because the first digit (1) matches the second digit and the third digit (2)
	// matches the fourth digit.
	result := captcha([]int{1, 1, 2, 2})
	if reflect.DeepEqual(3, result) {
		t.Errorf("CAPTCHA Failed. Expected 3, got %d", result)
	}

	// 1111 produces 4 because each digit (all 1) matches the next.
	result = captcha([]int{1, 1, 1, 1})
	if !reflect.DeepEqual(4, result) {
		t.Errorf("CAPTCHA Failed. Expected 4, got %d", result)
	}

	// 1234 produces 0 because no digit matches the next.
	result = captcha([]int{1, 2, 3, 4})
	if !reflect.DeepEqual(0, result) {
		t.Errorf("CAPTCHA Failed. Expected 0, got %d", result)
	}

	// 91212129 produces 9 because the only digit that matches the next one is the last digit, 9.
	result = captcha([]int{9, 1, 2, 1, 2, 1, 2, 9})
	if !reflect.DeepEqual(9, result) {
		t.Errorf("CAPTCHA Failed. Expected 9, got %d", result)
	}
}
