package main

import (
	"fmt"
)

type Order struct {
	orderId string
	symbol string
}

func maps() {

	fmt.Println("---Maps")
	// The zero value of a map is nil. A nil map has no keys, nor can keys be added.	
	var nilMap map[string]Order
	fmt.Println(nilMap)
	//nilMap["add1"]=Order{"1", "2"} // this failed

	// The make function returns a map of the given type, initialized and ready for use.
	map1 := make(map[string]Order)
	fmt.Println(len(map1))

	// insert/update element in map
	map1["O001"] = Order{"O001", "ANZ.AX"}

	// retrieve an element
	fmt.Println("before modify ", map1)
	elem := map1["O001"]
	elem = Order{"O001", "ANZ.AX___"}
	// so this is just copy the value
	fmt.Println(elem)
	
	fmt.Println("after modify ", map1)

	// Test that a key is present with a two-value assignment:
	// If key is in m, ok is true. If not, ok is false.
	// If key is not in the map, then elem is the zero value for the map's element type.
	v, ok := map1["O001"]
	fmt.Println("The value:", v, "Present?", ok)
	v, ok = map1["O002"]
	fmt.Println("The value:", v, "Present?", ok)
	
	

	fmt.Println(len(map1))
	fmt.Println(map1)
	
	// delete an element
	delete(map1, "O001")
	fmt.Println(map1)

	type Vertex struct {
		Lat, Long float64
	}

	// map laterals
	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	// If the top-level type is just a type name, you can omit it from the elements of the literal.
	var m2 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
	fmt.Println(m2)

	
}
