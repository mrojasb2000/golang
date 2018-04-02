package main

import "fmt"

// An asterisk is also used to dereference pointer variables. Dereferencing a pointer gives
// us access to the value the pointer points to.
func one(xPtr *int){
	*xPtr = 1
}

func main(){
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}
