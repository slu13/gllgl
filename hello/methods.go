package main

import (
	"fmt"
)

type Interpreter struct{
	name string
	something string
}

// A method is a function with a special receiver argument.
// With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.) 
func (v Interpreter) method1() string{
	// !!! This v is value copy, not the pointer
	v.name = "Changed"
	return v.name + v.something
}


// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
func (v *Interpreter) method3(){
	v.something = "m3something"
	v.name = "m3name"
}



/*
You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
*/

type MyInt int

func (v MyInt) method2() int {
	return int(v) * 10
}

func methods (){
	fmt.Println("---Methods")
	v := Interpreter{"i1", "something1"}
	fmt.Println(v.method1())
	fmt.Println(v)

	v.method3()
	fmt.Println(v)


	var myInt MyInt = 5
	fmt.Println(myInt.method2())

	

}
