package main

import "fmt"

func main(){
	i := 3

	switch i {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Tree")
		case 4: fmt.Println("Four")
		case 5: fmt.Println("Fi")
		default: fmt.Println("Unknown Number")
	}
}
