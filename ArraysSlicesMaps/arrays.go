package main;

import (
	"fmt"
)

/*
	Arrays in Go:
		- Fixed-size collections of elements of the same type (homogeneous).
		- Declared with a specific size.
		- Memory is allocated for all elements at once.
		- Memory layout is contiguous.
		- Syntax: var arr [size]type
*/

func main() {
	// Array declaration and initialization
	var arr [5] int = [5] int {10, 20, 30, 40, 50}

	fmt.Println("Array:", arr)
	fmt.Println("Length of array:", len(arr))
	fmt.Printf("Type of array: %T\n", arr)

	// Accessing elements
	fmt.Println("First element:", arr[0])
	fmt.Println("Third element:", arr[2])

	// Modifying elements
	arr[1] = 25
	fmt.Println("Modified array:", arr)

	// Iterating over the array
	fmt.Println("Array elements:")
	for i := 0; i < len(arr); i++ {		// len function returns the length of the array
		fmt.Print(arr[i], " ")
	}
	fmt.Println();

	// Using range to iterate over the array
	for index, value := range arr {  					// range function returns both index and value
		fmt.Printf("Element at index %d: %d\n", index, value)
	}

	// array declaration using shorthand syntax
	arr2 := [3]string{"Go", "is", "fun"}
	fmt.Println("Second array:", arr2)

	//array declaration using elipses
	arr3 := [...]float64{3.14, 1.59, 2.65, 3.58}   // size of array is inferred from the number of elements.
	fmt.Println("Third array:", arr3)

	// Multidimensional array
	var matrix [3][2] int = [3][2] int {{1, 2}, {3, 4}, {5, 6}}
	fmt.Println("Multidimensional array (matrix):", matrix)

	// Accessing elements in multidimensional array
	for i := 0; i < len(matrix); i++ {
		for j :=0 ; j < len(matrix[i]); j++ {
			fmt.Printf("Element at [%d][%d] is %d\n", i, j, matrix[i][j])
		}
	}

	//Accessing multidimensional array using range
	fmt.Println("Accessing multidimensional array using range:")
	for i, row := range matrix {
		for j, value := range row {
			fmt.Printf("Element at [%d][%d] is %d\n", i, j, value)
		}
	}
}