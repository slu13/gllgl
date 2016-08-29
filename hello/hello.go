package main

// "factored" import statement : is more recommended
import (
	"fmt"
//	"math/rand"
	"github.com/slu13/stringutil"
	"github.com/slu13/myutil"
)

// Alternative way of importing packages
// import "fmt"
// import "github*"

// variables with initializers, no need type if the initial value is given
var pkgVar1, pkgVar2, pkgVar3 = true, "abc", 15.04
//syntax error: non-declaration statement outside function body
//pkgVar4 := 1

func main() {
	fmt.Println("Welcome to my tour of go!")
	fmt.Println(stringutil.Reverse("abcdefg"))
	fmt.Println(stringutil.Reverse("abcdef"))
	fmt.Println("---Test myutil.RandomInt func")
	for i := 0; i < 10; i++ {
		fmt.Println(myutil.RandInt(1000))
	}
	fmt.Println("---Test myutil.Sqrt func")
	fmt.Println(myutil.Sqrt(4))
	fmt.Println(myutil.Sqrt(9))
	fmt.Println(myutil.Pi())
	fmt.Println("---Test exported/unexported func")
	fmt.Println(myutil.ExportedFunc())
	//Any "unexported" names are not accessible from outside the package.
	//fmt.Println(myutil.unexportedFunc())
	fmt.Println(myutil.PkgBoolExported)
	//fmt.Println(myutil.pkgBoolUnexported)
	var localBool bool = false
	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	localBool2 := true
	localBool = true
	fmt.Println(localBool, localBool2)
	fmt.Println(pkgVar1, pkgVar2, pkgVar3)
	
	//
	testValue()
	basicTypes()
	zeroValues()
	typeConversions()
	typeInference()
	constants()
	forLoops()
	ifElse()
	deferStatements()
	pointers()
	structs()
	arrays()

}
