package main;

import (
	"fmt"
)

/*
	Slices in Go:
	` A slice is defined as s contiguous segment of an underlying array and it provides numbered sequence of elements from that array.`
	Slices provide access to parts of an array in sequential order, they are more flexible and more powerfull than arrays.

	Slices are variable typed (elements can be added or removed)

	Slice has three major components:
	1. Pointer: points to the first element of the array that is accessible through the slice.
	2. Length: number of elements in the slice.
	3. Capacity: number of elements in the underlying array, counting from the first element in the slice.

	Declaring a slice:
	   <slice_name> := []<data_type>{<values>}
	
	** using len() and cap() functions to get length and capacity of the slice
	Example:
	   grades := []int{10,20,30,40}

	// Initalizing a slice from an array
	array[start_index : end_index]  --> end_index is exclusive
	array[start_index : ]  --> slice from start_index to the end of the array
	array[ : end_index]  --> slice from the start of the array to end_index-1
	array[ : ]  --> slice of the entire array
*/


func main() {
	var grades[10] int = [10] int {10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	// creating slices from the array
	slice1 := grades[2:5] // [30, 40, 50]
	slice2 := grades[4:]  // [50, 60, 70, 80, 90, 100]
	slice3 := grades[:3]
	slice4 := grades[:]   // entire array
	
	fmt.Println("Original grades array:", grades)
	fmt.Println("Slice1 (grades[2:5]):", slice1)
	fmt.Println("Slice2 (grades[4:]):", slice2)
	fmt.Println("Slice3 (grades[:3]):", slice3)
	fmt.Println("Slice4 (grades[:]):", slice4)

	// Length and Capacity of slices
	fmt.Printf("Length of slice1: %d, Capacity of slice1: %d\n", len(slice1), cap(slice1))
	fmt.Printf("Length of slice2: %d, Capacity of slice2: %d\n", len(slice2), cap(slice2))
}