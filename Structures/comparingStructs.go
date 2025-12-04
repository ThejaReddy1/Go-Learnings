package main

import "fmt"

type s struct {
	x int
}
type s1 struct {
	x int
}

func main(){
	a := s{1}
	// b := s1{1}
	c := s{5}
	d := s{1}
	//fmt.Println(a==b) //invalid operation: a == b (mismatched types s and s1)
	fmt.Println(a==c)
	fmt.Println(a==d)
}