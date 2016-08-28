package myutil

var PkgBoolExported, pgkBoolUnexported bool

func Swap(x string, y string) (string, string) {
	return y, x
}

/*
Named return values
Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.
*/
func SplitBy46(sum float64) (x, y float64){
	x = sum * (4.0/10.0)
	y = sum - x
	return
}
