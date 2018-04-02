package main

import "fmt"

// the zero function will not modify the original x variable in the main function
func zero(x int){
	x = 0
}

func main(){
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5
}
