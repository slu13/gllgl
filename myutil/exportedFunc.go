package myutil

func ExportedFunc() string{
	return "This is an exported func"
}

func unexportedFunc() string{
	return "This is NOT an exported func"
}
