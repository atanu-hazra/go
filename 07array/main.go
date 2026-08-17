package main

import "fmt"

func main() {
	// 1. Array declaration with fixed size: var <name> [<size>]<type>
	// Elements are initialized to their default zero-value (0 for int).
	var nums [3]int

	nums[1] = 2

	fmt.Println(nums) // output: [0 2 0]

	// 2. Initialize array while declaring (array literal)
	age := [3]int{10, 20, 30}
	random := [4]int{} // initialized with all zero-values [0 0 0 0]
	fmt.Println(age, random) // output: [10 20 30] [0 0 0 0]

	// Default zero-values for string ("") and bool (false)
	names := [3]string{}
	boolean := [3]bool{}

	fmt.Println(names, boolean)          // output: [  ] [false false false]
	fmt.Println(len(names), len(boolean)) // output: 3 3

	// The array size is part of the type signature: [3]string != [4]string
	fmt.Printf("Variable type: %T \n", names) // output: Variable type: [3]string 

	// 3. Iterating over arrays with loops:

	// Style A: Traditional index-based for loop
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i]) // output: 0 (1st iter), 2 (2nd iter), 0 (3rd iter)
	}

	// Style B: for-range loop getting both index and value
	for index, value := range nums {
		fmt.Println(index, value) // output: 0 0 (iter 0), 1 2 (iter 1), 2 0 (iter 2)
	}

	// Style C: for-range loop ignoring index using blank identifier '_'
	for _, value := range nums {
		fmt.Println(value) // output: 0 (iter 0), 2 (iter 1), 0 (iter 2)
	}

	// 4. Array Comparison:
	// Arrays of the same type and length can be directly compared with == (compares all elements).
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}

	fmt.Println(a == b) // output: true
}
