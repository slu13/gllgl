package main

import (
	"fmt"
)

func arrays() {
	fmt.Println("---Array")
	// array cannot be resized
	var a1 [10]int
	//a2 := [10]int
	a1[2] = 99
	fmt.Println(a1)
	
	fmt.Println(a1[2])

	a2 := [5]int{1,2,3,4,5}
	fmt.Println(a2)
	fmt.Println("---Slice")
	// A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

		
}
