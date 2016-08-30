package main

import (
	"fmt"
)

func playRange() {
	fmt.Println("---Range")
	// Array can be ranged
	array1 := [5]int{1,2,3,4,5}
	for index, value := range array1 {
		fmt.Println(index, value)
	}

	// Slice
	sl1 := array1[:]

	//
	fmt.Println("Values are copied")
	for index, value := range sl1 {
		fmt.Println("before", index, value)
		//value = 999 // not affect anything
		//sl1[index] = 999 // this will affect the original value
	}

	fmt.Println("full slice")
	for index, value := range sl1 {
		fmt.Println(index, value)
	}
	sl2 := array1[2:3]
	fmt.Println("2:3 slice")
	for index, value := range sl2 {
		fmt.Println(index, value)
	}


	// only need index
	//You can skip the index or value by assigning to _.
	for _, value := range sl2 {
		fmt.Println("value:", value)
	}

	// only need index
	for index, _ := range sl2 {
		fmt.Println("index:", index)
	}
}
