package main

import(
	"fmt"
)

func pointers() {
	fmt.Println("-- Pointers")

	// The type *T is a pointer to a T value. Its zero value is nil.
	var p *int
	fmt.Printf("%T, %v\n", p, p)
	vInt := 5
	p = &vInt
	fmt.Printf("%T, %v\n", p, p)
	fmt.Printf("%T, %v\n", *p, *p)
	// Dereferencing
	*p = 6
	fmt.Printf("%T, %v\n", p, p)
	fmt.Printf("%T, %v\n", *p, *p)
	p2 := &vInt
	*p2 = 7
	fmt.Printf("%T, %v\n", p, p)
	fmt.Printf("%T, %v\n", *p, *p)

		
	

}


