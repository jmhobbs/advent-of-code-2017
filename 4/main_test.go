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

func TestAnagramValidity(t *testing.T) {
	/*
		- abcde fghij is a valid passphrase.
		- abcde xyz ecdab is not valid - the letters from the third word can be rearranged to form the first word.
		- a ab abc abd abf abj is a valid passphrase, because all letters need to be used when forming another word.
		- iiii oiii ooii oooi oooo is valid.
		- oiii ioii iioi iiio is not valid - any of these words can be rearranged to form any other word.
	*/
	expected := map[string]bool{
		"abcde fghij":              true,
		"abcde xyz ecdab":          false,
		"a ab abc abd abf abj":     true,
		"iiii oiii ooii oooi oooo": true,
		"oiii ioii iioi iiio":      false,
	}

	for passphrase, expect := range expected {
		if anagramValidPassphrase(passphrase) != expect {
			t.Errorf("Error testing passphrase '%s', expected %v", passphrase, expect)
		}
	}
}
