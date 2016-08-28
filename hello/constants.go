package main

import (
	"fmt"
)

func constants(){
//Constants are declared like variables, but with the const keyword.
// Constants can be character, string, boolean, or numeric values.

	fmt.Println("---Constants")
	const charConst = '美'
	const strConst = "美国"
	const boolConst = false

	// Numeric constants
	const floatConst = 5.6453
	const(
		fc2 float32 = 5.6453
		fc3 float64 = 5.6453
		// Numeric constants are high-precision values.
		// An untyped constant takes the type needed by its context.
		Big = 1 << 60
	)
	

	fmt.Printf("%T, %v\n", charConst, charConst)
	fmt.Printf("%T, %v\n", strConst, strConst)
	fmt.Printf("%T, %v\n", boolConst, boolConst)
	fmt.Printf("%T, %v\n", floatConst, floatConst)
	fmt.Printf("%T, %v\n", fc2, fc2)
	fmt.Printf("%T, %v\n", fc3, fc3)
	fmt.Printf("%T, %v\n", Big, Big)


// Constants cannot be declared using the := syntax.
	//const inferenceConst := 21312.213

}
