// package name should be the same as the last element of the import path
// for example:import "github.com/user/myutil"
package myutil

import (
	"math/rand"
)

func RandInt(maxInt int) int{
	return rand.Intn(maxInt)
}

