package main

import "fmt"

func main() {
	// 1. Basic execution order (LIFO: Last-In, First-Out)
	defer fmt.Println("1st defer: executes last")
	defer fmt.Println("2nd defer: executes second")
	defer fmt.Println("3rd defer: executes first")

	fmt.Println("Main function body")

	for i:= 0; i < 5; i++ {
		defer fmt.Println("value: ", i)
	}
}


