package main

import "fmt"

func main() {
	// memory()
	pointers()
}

// Demonstrates memory allocation using new() vs address-of operator (&)
func memory() {
	// new(Type) allocates zeroed memory for the type and returns a pointer (*int)
	x := new(int)

	fmt.Println(x)  // output: 0xc000012028 (memory address/pointer)
	fmt.Println(*x) // output: 0 (default zero-value of int)
	fmt.Println(&x) // output: 0xc000058020 (memory address of pointer variable x itself)

	// Another way: create a standard variable and reference its address with &
	var value int
	y := &value
	fmt.Println(y, *y) // output: 0xc000012030 0 (address, value)

	// Allocating memory for string pointer with new()
	name := new(string)
	*name = "Rick"             // Dereference to assign value
	fmt.Println(name, *name)   // output: 0xc000010230 Rick (address, value)
}

// Demonstrates pointer manipulation and pass-by-pointer
func pointers() {
	// 1. Initialize a standard variable
	x := 10

	// 2. '&' (Address-of operator) gets the memory address of x
	p := &x

	fmt.Println(x, p) // output: 10 0xc000012140 (value, pointer address)

	// 3. '*' (Dereference operator) modifies the value at the memory address directly
	*p = 50

	fmt.Println(x) // output: 50 (x is updated because p points to x)

	// 4. Pass-by-pointer in functions:
	// Passing &a allows the function to mutate the original 'a' directly.
	a := 10

	change(&a) // Pass memory address of 'a'

	fmt.Println("a value after pass by pointer change", a) // output: a value after pass by pointer change 100
}

// change accepts a pointer to an integer (*int) and mutates the original value
func change(x *int) {
	*x = 100 // Dereferences and modifies original value in memory
}