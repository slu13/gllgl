package main

import (
	"fmt"
)

type myStruct struct {
	str1 string
	str2 string
	str3 string
}


/*
type Stringer interface {
    String() string
}

*/

// A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
func (this myStruct) String() string {
	return  this.str1 + ":" + this.str2 + ":" + this.str3
}

func stringsya(){
	fmt.Println("---Strings")
	v1 := myStruct{"a", "b", "c"}
	fmt.Println(v1)
	
}
