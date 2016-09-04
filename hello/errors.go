package main

import (
	"fmt"
	"time"
)

/*
Go programs express error state with error values.
The error type is a built-in interface similar to fmt.Stringer:
type error interface {
    Error() string
}
*/


type MyError struct{
	time time.Time
	msg string
}

func (this *MyError) Error() string{
	return this.time.String() + this.msg
}

func funcRtErr() (v int, err error) {
	return 5, &MyError{
		time.Now(),
		"something wrong",
	}
}

func errors(){
	fmt.Println("---Errors")
	

	// Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.
	v, err := funcRtErr()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
}
