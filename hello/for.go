package main

import (
	"fmt"
)

func forLoops(){
	fmt.Println("---For loops")
	sum := 0

// standard for loops
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)
// The init and post statement are optional.

	i := 0
	sum = 0
	for ; i < 10; {
		sum += i
		i++	
	}
	fmt.Println("sum = ", sum)

// While loop
	i = 0
	sum = 0
	for i < 10 {
		sum += i
		i++	
	}
	fmt.Println("sum = ", sum)

// infinite loop
	i = 0
	sum = 0
	for {
		if i < 10 {
			sum += i
			i++
		} else {
			break
		}
	}
	fmt.Println("sum = ", sum)
}
