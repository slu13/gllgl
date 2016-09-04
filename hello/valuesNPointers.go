package main

import (
	"fmt"
)

type Arg struct {
	v int
}

func funcWithPtrArg(v *Arg) (int) {
	v.v = 999
	return v.v * 10
}

func funcWithVarArg(v Arg) (int) {
	// v is copied, assigning value take no effects
	v.v = 999
	return v.v * 10
}

func (v *Arg) methodWithPtrRec() (int) {
	v.v = 999
	return v.v * 10
}

func (v Arg) methodWithVarRec() (int) {
	// v is copied, no effect
	v.v = 999
	return v.v * 10
}

func valueNPointers(){
	fmt.Println("---Values and pointers")

	v1 := Arg{5}
	p1 := &v1

	// Functions that take a value argument must take a value of that specific type 
	// fmt.Println(funcWithPtrArg(v1))//cannot use v1 (type Arg) as type *Arg in argument to funcWithPtrArg
	funcWithPtrArg(p1)
	fmt.Println(v1)
	
	v1.v = 5
	funcWithVarArg(v1)	
	fmt.Println(v1)
	//funcWithVarArg(p1) // cannot use p1 (type *Arg) as type Arg in argument to funcWithVarArg	

	//methods with pointer receivers take either a value or a pointer as the receiver when they are called
	// Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.
	v1.v = 5
	v1.methodWithPtrRec()
	p1.methodWithPtrRec()
	fmt.Println(v1)

	v1.v = 5
	v1.methodWithVarRec()
	p1.methodWithVarRec()
	fmt.Println(v1)


	// Choosing a value or pointer receiver
	// There are two reasons to use a pointer receiver.
	// The first is so that the method can modify the value that its receiver points to.
	// The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
	// In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.
 
}
