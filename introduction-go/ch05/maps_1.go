package main

import "fmt"

func main(){
	// A map is an unordered collection of key-value pairs
	//(maps are also sometimes called associative arrays, hash table, or dictionaries)
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
}
