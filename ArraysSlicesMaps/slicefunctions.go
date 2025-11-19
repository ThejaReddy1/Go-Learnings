package main;

import (
	"fmt"
);
// When we create a slice the slice gets its own numbering or index numbers.
/*
	when we change a value in the slice, the corresponding value in the underlying array also gets changed.
	Similarly, when we change a value in the underlying array, the corresponding value in the slice also gets changed.
	Slice is nothing but a reference to the underlying array.

	However, if we re-slice or append to the slice such that it exceeds its capacity, a new underlying array is created for the slice.
	This means that changes made to the new slice will not affect the original array, and vice versa.
*/
func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[:3] // slice of first three elements
	fmt.Println("Original array: ", arr)
	fmt.Println("Original slice: ", slice)

	// Modifying the slice
	slice[0] = 100
	fmt.Println("After modifying slice:")
	fmt.Println("Array: ", arr)   // arr[0] will be changed to 100
	fmt.Println("Slice: ", slice) // slice[0] will be 100


	// Re-slicing beyond capacity
	newSlice := slice[:4] // This exceeds the original slice capacity
	newSlice[3] = 300
	fmt.Println("After re-slicing and modifying newSlice:")
	fmt.Println("Array: ", arr)       // arr remains unchanged for index 3
	fmt.Println("Original Slice: ", slice) // slice remains unchanged for index 3
	fmt.Println("New Slice: ", newSlice)   // newSlice[3] will be 300

	// Appending to the slice
	// slice append function syntax: func append(slice []Type, elements...Type) []Type
	// alice := append(slice, element-1, element-2, ...., element-n)

	appendedSlice := append(slice, 400)  // If the capacity is exceeded, the capacity is doubled and a new underlying array is created
	fmt.Println("After appending to slice:")
	fmt.Println("Array: ", arr)		  // arr remains unchanged for index 1
	fmt.Println("Original Slice: ", slice)    // slice remains unchanged for index 1
	fmt.Println("Appended Slice: ", appendedSlice) // appendedSlice[1] will be 200

	// we can append a slice to another slice
	// alice := append(slice1, slice2...)
	// here slice2 is appended to slice1
	a := []int{500, 600, 700, 800, 900}
	slice_1 := a[:3] // [500, 600, 700]
	slice_2 := a[:3] // [500, 600, 700]
	
	fmt.Println("Slice 1: ", slice_1)
	fmt.Println("Slice 2: ", slice_2)
	newSlice = append(slice_1, slice_2...) // Appending the slice to another slice (... is a variadic parameter)

	fmt.Println("Appended Slice: ", newSlice) // appendedSlice will have all the elements

	// deleting an element from a slice
	// alice := append(slice[:index], slice[index+1:]...)
	sliceToDelete := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Slice for deletion: ", sliceToDelete)
	indexToDelete := 2 // deleting the element at index 2 (which is 3)
	sliceAfterDelete := append(sliceToDelete[:indexToDelete], sliceToDelete[indexToDelete+1:]...)
	fmt.Println("Slice after deletion: ", sliceAfterDelete) // should print [1, 2, 4, 5]

	// copying a slice to another slice
	// func copy(destination_slice, source_slice []Type) int
	sourceSlice := []int{10, 20, 30, 40, 50}
	destinationSlice := make([]int, 3) // creating a slice of length 3
	numCopied := copy(destinationSlice, sourceSlice) // copies first 3 elements from sourceSlice to destinationSlice
	fmt.Println("Source Slice: ", sourceSlice)
	fmt.Println("Destination Slice after copy: ", destinationSlice)
	fmt.Println("Number of elements copied: ", numCopied)

	//looping through a slice

	for i, value := range sourceSlice {
		fmt.Println("Index: ", i, "Value: ", value)
	}

	// looping through a slice without index
	for _, value := range sourceSlice {
		fmt.Println("Value: ", value)
	}
}