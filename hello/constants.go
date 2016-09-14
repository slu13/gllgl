package main

import (
	"fmt"
)


// iota give the first one 1
// second one 2
// and you cannot be Enum1 and Enum2 at the same time
// Mutually exclusive 'enum'
const (
	Enum1 = iota
	Enum2
	Enum3
	Enum4
)


// Enum == 1
// Enum2 = 10
// Enum3 = 100
// Enum4 = 1000
// So you can have Enum2 and Enum4 at the same time == 1010
// bitfield (inclusive or) 'enum'
const (
	Enum = 1 << iota
	Enum 2
	Enum 3
	Enum 4
)

const (
	CombineE1E3 = Enum | Enum3
	CombineE2E3 = Enum2 | Enum3
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
