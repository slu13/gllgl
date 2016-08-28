package main

import (
	"fmt"
)

func zeroValues(){
/*
The zero value is:

0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.
*/

	fmt.Println("---Zero Values")
	var(
		vInt int
		vBool bool
		vFloat32 float32
		vStr string
	)

	fmt.Println(
		vInt,
		vBool,
		vFloat32,
		vStr)
}
