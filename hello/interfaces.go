package main

import (
	"fmt"
)

type I1 interface {
	iMethod1(v int) int
}

type S1 struct {
	v int
}

func (v *S1) iMethod1(i int) int{
	return v.v * i
}

type S2 struct {
	v int
}


// Interfaces are implemented implicitly
// S2 implements the interface I1, but we don't need to explicityly declare that it does so.
func (v S2) iMethod1(i int) int{
	return v.v * i
}


type EmptyInterface interface{

}

func interfaces(){
	fmt.Println("---Interfaces")
	var a I1
	// a.iMethod1(5) // Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
	describe(a) // (<nil>, <nil>)
	var s0 S2
	a = s0
	describe(a) // ({0}, main.S2)
	a.iMethod1(5)
	s1 := S1{7}

	// a = s1 // S1 does not implement I1 (iMethod1 method has pointer receiver)
	a = &s1
	fmt.Println(a.iMethod1(9))
	describe(a) // (&{7}, *main.S1)


	// However, if the receiver is defined as a value, s2 / &s2 are available
	s2 := S2{9}
	a = s2
	describe(a) // ({9}, main.S2)
	a = &s2
	describe(a) // (&{9}, *main.S2)

	// we can start polimorphisim here
	//var i I1 = S2{5}


	// Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.
	var i2 EmptyInterface
	i2 = 3435.435
	fmt.Println(i2)

	// Type assertions
	a = &s1
	t, ok := a.(*S1)
	fmt.Println(t, ok) // TRUE
	t2, ok := a.(S2)
	fmt.Println(t2, ok) // FALSE
	a = s2 
	t3, ok := a.(S2)
	fmt.Println(t3, ok) // TRUE
	t4, ok := a.(*S2)
	fmt.Println(t4, ok) // FALSE
	
	typeSwitches()
}

func typeSwitches(){
	
	fmt.Println("---Type switches")
	do(21)
	do("hello")
	do(true)
}

func do(i EmptyInterface) {
	// The declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword type.
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func describe(i I1) {
	fmt.Printf("(%v, %T)\n", i, i)
}
