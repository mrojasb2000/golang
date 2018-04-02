package main

import "fmt"

func main(){
	// Slice is a segment of an array. Like array, slices are indexable and have  a length.
	// x  has been created with  a length zero.
	//var x []float64

	// This creates a slice is associated with an underlying float64 array of length 5.
	// Slices are always associated with some array, and although they can never be longer
	// than the array, they can be smaller.
	//x := make([]float64, 5)

	arr := [5]float64{1,2,3,4,5}
	x := arr[0:5]
	fmt.Println(x)
}