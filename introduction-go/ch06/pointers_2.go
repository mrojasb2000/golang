package main

import "fmt"

// Pointers reference a location in memory where a value is stored rather than the value itself.
// By using a pointer (*int), the zero function is able to modify the original variable.

// An asterisk is also used to dereference pointer variables. Dereferencing a pointer gives
// us access to the value the pointer points to.
func zero(xPtr *int){
	// When we write *xPtr = 0, we are saying "store the int 0 in the memory location xPtr refers to"
	*xPtr = 0
}

func main(){
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0
}
