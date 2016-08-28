package myutil

import (
	"testing"
)

func TestAdd(t *testing.T){
	cases := []struct{
		inputX, inputY, expectedResult float64
	}{
		{5.0, 6.0, 11.0},
		{2.0, 3.0, 5.0},
	}
	for _, c := range cases {
		if Add(c.inputX, c.inputY) != c.expectedResult {
			t.Errorf("%q + %q == %q != %q",
				c.inputX, c.inputY, Add(c.inputX, c.inputY),
				c.expectedResult)
		}
	}
}
