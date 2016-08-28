package stringutil

import (
	"testing"
)

func TestReverse(t *testing.T){
	cases := []struct {
		input, expectedResult string
	}{
		{"abcde", "edcba"},
		{"abcd", "dcba"},
		{"", ""},
		{"世界", "界世"},
	}

	for _, c := range cases {
		testResult := Reverse(c.input)
		if testResult != c.expectedResult {
			t.Errorf("Reverse(%q) == %q, expected %q",
				c.input, testResult, c.expectedResult)
		}
	}
}
