package main

import "testing"

type Test struct {
	a      string
	result string
}

var tests = []Test{
	{"fffGyhd", "fffGyhd"},
	{"aaaaaa", ""},
	{"aasaa", "s"},
	{"b", "b"},
	{"", ""},
	{"aeiou", ""},
}

func TestDeleteVowels(t *testing.T) {
	for i, y := range tests {
		result := DeleteVowels(y.a)
		if result != y.result {
			t.Errorf("#%d: DeleteVowels(%s) = %s ожидалось %s", i, y.a, result, y.result)
		}
	}
}
