package main

import "testing"

func TestValidPassphrase(t *testing.T) {
	/*
		- aa bb cc dd ee is valid.
		- aa bb cc dd aa is not valid - the word aa appears more than once.
		- aa bb cc dd aaa is valid - aa and aaa count as different words.
	*/
	expected := map[string]bool{
		"aa bb cc dd ee":  true,
		"aa bb cc dd aa":  false,
		"aa bb cc dd aaa": true,
	}

	for passphrase, expect := range expected {
		if validPassphrase(passphrase) != expect {
			t.Errorf("Error testing passphrase '%s', expected %v", passphrase, expect)
		}
	}
}
