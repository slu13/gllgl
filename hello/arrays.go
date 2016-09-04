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

	a2 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(a2)
	fmt.Println("---Slice")
	// A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

	var s1 []int = a2[1:3]
	s2 := a2[1:3]
	fmt.Println(s1)
	fmt.Println(s2)
	// A slice does not store any data, it just describes a section of an underlying array.
	s2[0] = 999
	fmt.Println(a2, s1, s2)
	fmt.Printf("%T, %V\n", s2, s2)

	// Array literal
	al1 := [3]bool{true, false, true}
	fmt.Printf("%T, %V\n", al1, al1)
	// And this creates the same array as above, then builds a slice that references it
	al2 := []bool{true, false, true}
	fmt.Printf("%T, %V\n", al2, al2)
	
	cases := []struct{
		input int
		er int
	}{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(cases)
	
	// Slice defaults
	sd1 := a2[:]
	sd2 := a2[:5] // be careful here, 5 is the high bound, not the index of the last item
	sd3 := a2[0:]
	sd4 := a2[0:5]
	fmt.Println(sd1, sd2, sd3, sd4)
	fmt.Println("testing array slice cap")
	printSlice(a2[:])
	printSlice(a2[2:3])

	// Slice's length and capacity
	s := []int{1,2,3,4,5}
	printSlice(s)
	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)
	
	// Extend its length.
	s = s[1:4]
	printSlice(s)

	// Extend the length to exceed its capacity
	// runtime error: slice bounds out of rang
	// s = s[:6]
	// printSlice(s)
	
	// Nil Slice
	// The zero value of a slice is nil. A nil slice has a length and capacity of 0 and has no underlying array.
	var ns []int
	printSlice(ns)


	// Create Slice with make
	ms1 := make([]int, 5) // len 5, cap 5
	printSlice(ms1)

	ms2 := make([]int, 3, 10) // len 3, cap 10
	printSlice(ms2)


	// Slices of Slices
	ss1 := [][]int {
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	printSliceOfSlices(ss1)

	ss1[0][2] = 0
	ss1[2][1] = 0
	printSliceOfSlices(ss1)

	// Append items to slices
	as1 := make([]int, 3, 5)
	
	printSlice(as1)
	as1 = append(as1, 1, 2)
	printSlice(as1)
	as1 = append(as1, 1)
	printSlice(as1)

	aaaa2 := [5]int{1, 2,3,4,5}
	as2 := aaaa2[1:3]
	printSlice(as2)
	
	as2 = append(as2, 1, 2, 3, 4, 8, 9)
	printSlice(as2)

	

}


func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSliceOfSlices(ss1 [][]int){	
	for i := 0; i < len(ss1); i++ {
		fmt.Println(ss1[i])
	}
}
