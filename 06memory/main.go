package main

import "fmt"

func main() {
	// memory()
	pointers()
}

func memory() {
	// Allocates memory and returns a pointer.
	x := new(int)

	fmt.Println(x) // memory pointer
	fmt.Println(*x) // value
	fmt.Println(&x) // memory pointer

	// another way
	var value int
	y := &value
	fmt.Println(y, *y) 

	name := new(string)
	*name = "Rick"
	fmt.Println(name, *name)
}

func pointers()  {
	x := 10 // initializing a variable

	p := &x // getting the memory adress (pointer) of that variable

	fmt.Println(x, p) 

	*p = 50 // changing the value of that memory adress using pointer

	fmt.Println(x)

	// Using pass by pointer
	a := 10

	change(&a)

	fmt.Println("a value after pass by pointer change", a)
}

// Pass by Pointer
func change(x *int) {
	*x = 100
}