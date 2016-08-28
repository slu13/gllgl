package main

import (
	"fmt"
)

//A struct is a collection of fields.
//(And a type declaration does what you'd expect.)
type Vertex struct {
	X int
	Y int
	z int
}

func structs() {
	fmt.Println("--- Structs")
	fmt.Println(Vertex{1,2,3})
	v := Vertex{4,5,6}
	fmt.Println(v.X, v.Y, v.z)
	fmt.Println("--- Structs pointers")
	p := &Vertex{7,8,9}

	// To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
	fmt.Println(p.X, (*p).Y, p.z)
	fmt.Println(p, *p)

	fmt.Println("--- Structs literals")
	var v1 = Vertex{1, 2, 3}
	var v2 Vertex = Vertex{1, 2, 3}
	v3 := Vertex{1, 2, 3}
	pv4 := &Vertex{1, 2, 3}

	fmt.Println(v1, v2, v3, pv4)

	//v5 := Vertex
	v6 := Vertex{}
	//v7 := Vertex{1}
	v8 := Vertex{1, 2, 3}
	// You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)
	v9 := Vertex{z:3}
	fmt.Println(v6, v8, v9)


}
