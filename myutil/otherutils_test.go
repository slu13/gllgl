package myutil

import (
	"testing"
)

func TestSwap(t *testing.T){
	cases := []struct{
		x, y, x2, y2 string
	}{
		{"xstring", "ystring", "ystring", "xstring"},
		{"", "", "", ""},
	}

	for _, c := range cases {
		x2r, y2r := Swap(c.x, c.y)
		if x2r != c.x2 || y2r != c.y2 {
			t.Errorf("Swap(%q, %q) got unexpected results)",
				c.x,c.y)
		}
	}
}

func TestSplitBy46(t *testing.T){
	cases := []struct{
		sum, x2, y2 float64
	}{
		{10.0, 4.0, 6.0},
		{100.0, 40.0, 60.0},
		{3.0, 1.2, 1.8},
	}

	for _, c := range cases{
		x, y := SplitBy46(c.sum)
		if x != c.x2 || y != c.y2 {
			t.Errorf("SplitWrong %q -> %q, %q", c.sum, x, y)
		}
	}
}
