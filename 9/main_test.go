package main

import (
	"bytes"
	"testing"
)

func TestExamples(t *testing.T) {
	samples := map[string]int{
		"{}":                        1,
		"{{{}}}":                    3,
		"{{},{}}":                   3,
		"{{{},{},{{}}}}":            6,
		"{<{},{},{{}}>}":            1,
		"{<a>,<a>,<a>,<a>}":         1,
		"{{<a>},{<a>},{<a>},{<a>}}": 5,
		"{{<!>},{<!>},{<!>},{<a>}}": 2,
	}

	for str, expected := range samples {
		r := bytes.NewBufferString(str)
		g := ProcessStream(r)
		count := g.Count()
		if expected != count {
			t.Errorf("'%s', expected %d groups, got %d.", str, expected, count)
		}
	}
}

func TestScore(t *testing.T) {
	samples := map[string]int{
		"{}":                            1,
		"{{{}}}":                        6,
		"{{},{}}":                       5,
		"{{{},{},{{}}}}":                16,
		"{<a>,<a>,<a>,<a>}":             1,
		"{{<ab>},{<ab>},{<ab>},{<ab>}}": 9,
		"{{<!!>},{<!!>},{<!!>},{<!!>}}": 9,
		"{{<a!>},{<a!>},{<a!>},{<ab>}}": 3,
	}

	for str, expected := range samples {
		r := bytes.NewBufferString(str)
		g := ProcessStream(r)
		score := g.TotalScore()
		if expected != score {
			t.Errorf("'%s', expected %d score, got %d.", str, expected, score)
		}
	}
}

func TestGarbageCount(t *testing.T) {
	samples := map[string]int{
		"<>": 0,
		"<random characters>": 17,
		"<<<<>":               3,
		"<{!>}>":              2,
		"<!!>":                0,
		"<!!!>>":              0,
		"<{o\"i!a,<{i<a>":     10,
	}

	for str, expected := range samples {
		r := bytes.NewBufferString(str)
		g := ProcessStream(r)
		garbage := g.TotalGarbageCount()
		if expected != garbage {
			t.Errorf("'%s', expected %d garbage, got %d.", str, expected, garbage)
		}
	}
}
