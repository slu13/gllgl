package main

import (
	"fmt"
)

func compute(fn func(int, int) int, a int, b int) int {
	return fn(a, b)
}

func getFunc(index int) func(int, int) int {
	if index == 1 {
		return func (a int, b int) int {
			return a - b
		}
	}else {
		return nil
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}




func functions() {
	fmt.Println("---FUnctions")

	addFunc := func(a, b int) int {
		return a + b
	}
	fmt.Println(compute(addFunc, 3, 5))
	fmt.Println(addFunc(3, 5))
	fmt.Println(compute(func(a, b int) int {
		return a * b
	}, 3, 5))

	fmt.Println(getFunc(1)(6, 7))
	
	fmt.Println(adder(), adder())
}
