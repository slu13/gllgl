package main

import (
	"fmt"
)

func deferStatements(){
	fmt.Println("---defer")
	f1()
	defer f2()
	f3()
	
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")	
}

func f1(){
	fmt.Println("f1")
	defer f11()
}

func f2(){
	fmt.Println("f2")
	defer f21()
}

func f3(){
	fmt.Println("f3")
}

func f11(){
	fmt.Println("f11")
}

func f21(){
	fmt.Println("f21")
}
