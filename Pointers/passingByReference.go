package main

import "fmt"

/*
	Passing by reference functions:

	-> In call by reference/pointer, the adderess of the variable is passed into the function call as the actual parameter.
	-> All the Operations in the function are performed on value stored at the address of the parameters.

	-> Slices are passed by reference, by default in Go.
	-> Maps are passed by reference, by default in Go.
	-> Channels are passed by reference, by default in Go.
	-> Functions are passed by reference, by default in Go.
*/
func swap( a *int, b *int) {
	*a = *a + *b
	*b = *a - *b 
	*a = *a - *b 
}
func modifySlice(s []int) {
	s[2] = 100
}

func modifyMap(m map[string]int) {
	m["three"] = 3
}
func main(){
	a := 2
	b := 3
	fmt.Println("Before swap: a =", a, "b =", b) // Output: Before swap: a = 2 b = 3
	swap(&a, &b) // call by reference
	fmt.Println("After swap: a =", a, "b =", b)  // Output: After swap: a = 3 b = 2

	slice := []int{1, 2, 3}
	fmt.Println("Before modifySlice:", slice)
	modifySlice(slice)
	fmt.Println("After modifySlice:", slice)

	mapData := make(map[string]int)
	mapData["one"] = 1
	mapData["two"] = 2
	fmt.Println("Before modifyMap:", mapData)
	modifyMap(mapData)
	fmt.Println("After modifyMap:", mapData)

}