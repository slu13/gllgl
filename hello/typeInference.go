package main

import (
	"fmt"
)

func typeInference(){
	fmt.Println("---Type inference")	
/*
When the right hand side of the declaration is typed, the new variable is of that same type
*/

	var vStr string = "12312"
	vStr2 := vStr
	fmt.Printf("%T, %v\n", vStr2, vStr2)

/*
But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:
*/
	i := 5 // int
	j := 5.012 // float64
	k := 0.867 + 0.5i // complex128

	fmt.Printf("%T, %v\n", i, i)
	fmt.Printf("%T, %v\n", j, j)
	fmt.Printf("%T, %v\n", k, k)

}
