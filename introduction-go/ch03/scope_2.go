package main

import "fmt"

// Notice that we moved the variable outside of the main function.
// This means that other functions can access this variable
var x string = "Hello, World"

func main(){
	fmt.Println(x)
}

func f(){
	fmt.Println(x)
}