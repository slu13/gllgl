package main

import(
	"fmt"
)

func typeConversions(){
/*
Unlike in C, in Go assignment between items of different type requires an explicit conversion.
*/
	fmt.Println("---Type conversions")
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(
		i,
		f,
		u)


}
