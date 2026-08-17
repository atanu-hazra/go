package main

import "fmt"

func main() {
	// 1. Slice literal creation: creates a dynamically sized slice
	mySlice := []int{}
	fmt.Println(mySlice) // output: []

	// 2. make([]Type, len, cap): creates slice with initial length=3 (filled with 0s) and capacity=5
	mySlice2 := make([]int, 3, 5)

	// Appending 5 elements when cap is only 5:
	// Total elements = 3 (zeros) + 5 = 8.
	// Because 8 > 5, Go automatically allocates a new underlying array with larger capacity (e.g. 10),
	// copies existing elements over, and updates the slice pointer.
	mySlice2 = append(mySlice2, 4, 6, 6, 8, 9)

	// 3. Deleting an element at a specific index from a slice:
	// Go does not have a built-in delete() for slices; use append() with slice bounds and variadic '...'
	deleteSlice := []int{10, 20, 30, 40, 50}
	fmt.Println("Before:", deleteSlice) // output: Before: [10 20 30 40 50]

	index := 2 // Remove element at index 2 (value: 30)
	deleteSlice = append(deleteSlice[:index], deleteSlice[index+1:]...)
	fmt.Println("After:", deleteSlice) // output: After: [10 20 40 50]

	// 4. for-range loop over slice:
	for i, v := range mySlice2 {
		fmt.Printf("Index: %d value: %d type %T\n", i, v, v)
		// output: Index: 0 value: 0 type int
		// output: Index: 1 value: 0 type int
		// output: Index: 2 value: 0 type int
		// output: Index: 3 value: 4 type int
		// output: Index: 4 value: 6 type int
		// output: Index: 5 value: 6 type int
		// output: Index: 6 value: 8 type int
		// output: Index: 7 value: 9 type int
	}
}

// Comprehensive slice demonstration covering all core slice mechanics in Go
func demoSlice() {
	// Method 1: Slice literal
	nums := []int{10, 20, 30}
	fmt.Println("Slice Literal:", nums) // output: Slice Literal: [10 20 30]

	// Method 2: Using make(type, length) - initialized with default zero values
	makeSlice := make([]int, 5)
	fmt.Println("make([]int, 5):", makeSlice) // output: make([]int, 5): [0 0 0 0 0]

	// Method 3: Using make(type, length, capacity)
	capSlice := make([]int, 3, 5)
	fmt.Println("make([]int, 3, 5):", capSlice) // output: make([]int, 3, 5): [0 0 0]

	fmt.Println("\n=========== Length & Capacity ===========")
	// len() = number of elements present
	// cap() = maximum elements slice can hold before resizing underlying array
	fmt.Println("len(nums):", len(nums)) // output: len(nums): 3
	fmt.Println("cap(nums):", cap(nums)) // output: cap(nums): 3

	fmt.Println("len(capSlice):", len(capSlice)) // output: len(capSlice): 3
	fmt.Println("cap(capSlice):", cap(capSlice)) // output: cap(capSlice): 5

	fmt.Println("\n=========== Accessing Elements ===========")
	fmt.Println("First:", nums[0])           // output: First: 10
	fmt.Println("Last:", nums[len(nums)-1]) // output: Last: 30

	fmt.Println("\n=========== Updating Elements ===========")
	nums[1] = 99
	fmt.Println(nums) // output: [10 99 30]

	fmt.Println("\n=========== Append ===========")
	// append() returns a new slice header with updated length (and new backing array if resized)
	nums = append(nums, 40)
	fmt.Println(nums) // output: [10 99 30 40]

	// Appending multiple elements at once
	nums = append(nums, 50, 60, 70)
	fmt.Println(nums) // output: [10 99 30 40 50 60 70]

	fmt.Println("\n=========== Slicing ===========")
	// Slice operator: [start:end] -> includes start, excludes end
	data := []int{10, 20, 30, 40, 50}

	fmt.Println("Original:", data)   // output: Original: [10 20 30 40 50]
	fmt.Println("data[1:4]:", data[1:4]) // output: data[1:4]: [20 30 40]
	fmt.Println("data[:3]:", data[:3])   // output: data[:3]: [10 20 30]
	fmt.Println("data[2:]:", data[2:])   // output: data[2:]: [30 40 50]
	fmt.Println("data[:]:", data[:])     // output: data[:]: [10 20 30 40 50]

	fmt.Println("\n=========== Shared Underlying Array ===========")
	// Sub-slicing does NOT copy data; it shares the same underlying array memory!
	original := []int{1, 2, 3, 4, 5}
	sub := original[1:4]

	fmt.Println("Original:", original) // output: Original: [1 2 3 4 5]
	fmt.Println("Sub:", sub)           // output: Sub: [2 3 4]

	// Modifying sub also changes original because they reference the same array
	sub[0] = 100

	fmt.Println("After modifying sub slice")
	fmt.Println("Original:", original) // output: Original: [1 100 3 4 5]
	fmt.Println("Sub:", sub)           // output: Sub: [100 3 4]

	fmt.Println("\n=========== Copy Slice ===========")
	// copy(dest, src) creates an independent copy; modifying copy won't affect source
	copySource := []int{10, 20, 30}
	copyDest := make([]int, len(copySource)) // Must pre-allocate destination slice

	copy(copyDest, copySource)
	copyDest[0] = 999

	fmt.Println("Source:", copySource) // output: Source: [10 20 30]
	fmt.Println("Copied:", copyDest)   // output: Copied: [999 20 30]

	fmt.Println("\n=========== Delete Element ===========")
	deleteSlice := []int{10, 20, 30, 40, 50}
	fmt.Println("Before:", deleteSlice) // output: Before: [10 20 30 40 50]

	index := 2
	deleteSlice = append(deleteSlice[:index], deleteSlice[index+1:]...)
	fmt.Println("After:", deleteSlice) // output: After: [10 20 40 50]

	fmt.Println("\n=========== Range Loop ===========")
	for index, value := range nums {
		fmt.Printf("Index: %d Value: %d\n", index, value)
		// output: Index: 0 Value: 10
		// output: Index: 1 Value: 99
		// output: Index: 2 Value: 30
		// output: Index: 3 Value: 40
		// output: Index: 4 Value: 50
		// output: Index: 5 Value: 60
		// output: Index: 6 Value: 70
	}

	fmt.Println("\nOnly Values")
	for _, value := range nums {
		fmt.Println(value) // output: 10 \n 99 \n 30 \n 40 \n 50 \n 60 \n 70
	}

	fmt.Println("\n=========== Nil Slice ===========")
	// Uninitialized slice is nil; has len=0 and cap=0, but is safe to append to
	var nilSlice []int

	fmt.Println(nilSlice)                      // output: []
	fmt.Println("Is nil:", nilSlice == nil)   // output: Is nil: true
	fmt.Println("Length:", len(nilSlice))     // output: Length: 0
	fmt.Println("Capacity:", cap(nilSlice))   // output: Capacity: 0

	nilSlice = append(nilSlice, 10)
	fmt.Println("After append:", nilSlice)    // output: After append: [10]

	fmt.Println("\n=========== Append Growth ===========")
	// Demonstrates how Go dynamically allocates and grows capacity as needed
	grow := []int{}
	for i := 1; i <= 10; i++ {
		grow = append(grow, i)
		fmt.Printf("Value Added: %d | len=%d cap=%d\n", i, len(grow), cap(grow))
		// output: Value Added: 1 | len=1 cap=1
		// output: Value Added: 2 | len=2 cap=2
		// output: Value Added: 3 | len=3 cap=4
		// output: Value Added: 4 | len=4 cap=4
		// output: Value Added: 5 | len=5 cap=8
		// output: Value Added: 6 | len=6 cap=8
		// output: Value Added: 7 | len=7 cap=8
		// output: Value Added: 8 | len=8 cap=8
		// output: Value Added: 9 | len=9 cap=16
		// output: Value Added: 10 | len=10 cap=16
	}

	fmt.Println("\n=========== For Loop ===========")
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i]) // output: 10 \n 99 \n 30 \n 40 \n 50 \n 60 \n 70
	}
}