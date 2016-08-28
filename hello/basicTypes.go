package main

import (
	"fmt"
)

func basicTypes(){
	fmt.Println("----Enter basicTypes testing")
	fmt.Println("-Int")
	// variables declarations can be factored into blocks
	var (
		vInt int = 5
		vInt8 int8 = 5
		vInt16 int16 = 5
		vInt32 int32 = 5
		vInt64 int64 = 5
		vMaxInt uint64 = 1<<64 -1
	)
	//  also have unsigned types uint uint8 uint16 uint32 uint64 uintptr)
	
	fmt.Println(
		vInt,
		vInt8,
		vInt16,
		vInt32,
		vInt64,
		vMaxInt)

	/*
The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
	*/

	fmt.Println("-Int alias")
	// byte == uint8
	var vByte byte = 'f'
	fmt.Println(vByte)

	// rune == int32, represent a Unicode code point
	var vRune rune = '美'
	fmt.Println(vRune)


	fmt.Println("-Bool")
	var vBool bool = true
	fmt.Println(vBool)

	fmt.Println("-String")
	var vStr string = "lushiy"
	fmt.Println(vStr)

	fmt.Println("-Float")
	var vFloat32 float32 = 4.0
	var vFloat64 float64 = 4.0
	fmt.Println(vFloat32, vFloat64)

	// also have complex64, complex128
}
