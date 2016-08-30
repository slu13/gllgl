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

	// The make function returns a map of the given type, initialized and ready for use.
	map1 := make(map[string]Order)
	map1["O001"] = Order{"O001", "ANZ.AX"}

	fmt.Println(map1)
}
