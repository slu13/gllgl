package main

import (
	"fmt"
)

func ifElse() {
	i := 5

	fmt.Println("---if else")
	// standard if
	if i < 10 {
		fmt.Println("i < 10")
	}

	// if with a short statement to execute before the condition
	if i2 := 10; i2 > 10 {
		fmt.Println("v2 > 10")
	} else {
	// Variables declared inside an if short statement are also available inside any of the else blocks.
		fmt.Println("v2 <= 10")
	}

	// Variables declared by the statement are only in scope until the end of the if.
	//i2 = 9	

	fmt.Println("---Switch")
	// standard switch
	// A case body breaks automatically,
	for j := 0 ; j < 4; j++ {
		switch j {
		case 0:
			fmt.Print("0\n")
		case 1:
			fmt.Print("1\n")
		case 2:
			fmt.Print("2\n")
		default:
			fmt.Print("Default\n")
		}

	}
	// fallthrough switch
	fmt.Println("---Fallthrough Switch")
	
	for j := 0 ; j < 4; j++ {
		switch j {
		case 0:
			fmt.Print("0\n")
			fallthrough
			// with this, j==0 will result 012Default
		case 1:
			fmt.Print("1\n")
			fallthrough
		case 2:
			fmt.Print("2\n")
			fallthrough
		default:
			fmt.Print("Default\n")
		}

	}

	fmt.Println("---True Switch")
	// switch true
	// Switch without a condition is the same as switch true.
	// This construct can be a clean way to write long if-then-else chains.
	for j := 0 ; j < 6; j++ {
		switch {
		case j < 2:
			fmt.Println(j, "<2")
		case j < 4:
			fmt.Println(j, "2<=j<4")
		default:
			fmt.Println(j, ">=4")	
		}
	}
}	
