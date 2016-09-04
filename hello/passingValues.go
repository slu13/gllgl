package main

import (
	"fmt"
)

func passArr(v []int) {
	v[0] = 999
}

func passInt(v int){
	v = 999
}

func passMap(v map[int]int){
	v[5] = 5
}

func passingValues() {
	fmt.Println("--- Passing values")

	// it's passing the slice to the func, and slice is a reference, so modifications matters
	v1 := []int{1,2,3,4,5}
	fmt.Printf("%T, %v\n", v1, v1)
	passArr(v1)
	fmt.Printf("%T, %v\n", v1, v1)
	v2 := make([]int, 50)
	fmt.Printf("%T, %v\n", v2, v2)
	passArr(v2)
	fmt.Printf("%T, %v\n", v2, v2)
	
	// passing int just copy, no effects	

	v3 := 5
	fmt.Printf("%T, %v\n", v3, v3)
	passInt(v3)
	fmt.Printf("%T, %v\n", v3, v3)


	// passing maps, it's passing the reference
	v4 := make(map[int]int)
	fmt.Printf("%T, %v\n", v4, v4)
	passMap(v4)
	fmt.Printf("%T, %v\n", v4, v4)
	
	
}
