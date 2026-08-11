package main

import "fmt"

func main() {
	mySlice := []int{}
	fmt.Println(mySlice)

	mySlice2 := make([]int, 3, 5)
	mySlice2 = append(mySlice2, 4, 6, 6, 8, 9)
	// 	But capacity is only 5.
	// So Go says:
	// "I don't have enough space."
	// It automatically:
	// Allocates a new, larger underlying array.
	// Copies the existing elements.
	// Appends the new elements.
	// Makes mySlice2 point to the new array.

	// fmt.Println(mySlice2)
	// fmt.Println(len(mySlice2)) // 8 (increased)
	// fmt.Println(cap(mySlice2)) // 10 (increased)
	// fmt.Println(mySlice2[1:4])

	deleteSlice := []int{10, 20, 30, 40, 50}

	fmt.Println("Before:", deleteSlice)

	index := 2

	deleteSlice = append(deleteSlice[:index], deleteSlice[index+1:]...)

	fmt.Println("After:", deleteSlice)


	// for loop
	for i, v := range mySlice2 {
		fmt.Printf("Index: %d value: %d type %T\n", i, v, v)
	}

}


func demoSlice()  {
	// Method 1: Slice literal
	nums := []int{10, 20, 30}
	fmt.Println("Slice Literal:", nums)

	// Method 2: Using make()
	makeSlice := make([]int, 5)
	fmt.Println("make([]int, 5):", makeSlice)

	// Method 3: make with length and capacity
	capSlice := make([]int, 3, 5)
	fmt.Println("make([]int, 3, 5):", capSlice)

	fmt.Println("\n=========== Length & Capacity ===========")

	fmt.Println("len(nums):", len(nums))
	fmt.Println("cap(nums):", cap(nums))

	fmt.Println("len(capSlice):", len(capSlice))
	fmt.Println("cap(capSlice):", cap(capSlice))

	fmt.Println("\n=========== Accessing Elements ===========")

	fmt.Println("First:", nums[0])
	fmt.Println("Last:", nums[len(nums)-1])

	fmt.Println("\n=========== Updating Elements ===========")

	nums[1] = 99
	fmt.Println(nums)

	fmt.Println("\n=========== Append ===========")

	nums = append(nums, 40)
	fmt.Println(nums)

	nums = append(nums, 50, 60, 70)
	fmt.Println(nums)

	fmt.Println("\n=========== Slicing ===========")

	data := []int{10, 20, 30, 40, 50}

	fmt.Println("Original:", data)
	fmt.Println("data[1:4]:", data[1:4])
	fmt.Println("data[:3]:", data[:3])
	fmt.Println("data[2:]:", data[2:])
	fmt.Println("data[:]:", data[:])

	fmt.Println("\n=========== Shared Underlying Array ===========")

	original := []int{1, 2, 3, 4, 5}

	sub := original[1:4]

	fmt.Println("Original:", original)
	fmt.Println("Sub:", sub)

	sub[0] = 100

	fmt.Println("After modifying sub slice")
	fmt.Println("Original:", original)
	fmt.Println("Sub:", sub)

	fmt.Println("\n=========== Copy Slice ===========")

	copySource := []int{10, 20, 30}

	copyDest := make([]int, len(copySource))

	copy(copyDest, copySource)

	copyDest[0] = 999

	fmt.Println("Source:", copySource)
	fmt.Println("Copied:", copyDest)

	fmt.Println("\n=========== Delete Element ===========")

	deleteSlice := []int{10, 20, 30, 40, 50}

	fmt.Println("Before:", deleteSlice)

	index := 2

	deleteSlice = append(deleteSlice[:index], deleteSlice[index+1:]...)

	fmt.Println("After:", deleteSlice)

	fmt.Println("\n=========== Range Loop ===========")

	for index, value := range nums {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}

	fmt.Println("\nOnly Values")

	for _, value := range nums {
		fmt.Println(value)
	}

	fmt.Println("\n=========== Nil Slice ===========")

	var nilSlice []int

	fmt.Println(nilSlice)
	fmt.Println("Is nil:", nilSlice == nil)
	fmt.Println("Length:", len(nilSlice))
	fmt.Println("Capacity:", cap(nilSlice))

	nilSlice = append(nilSlice, 10)

	fmt.Println("After append:", nilSlice)

	fmt.Println("\n=========== Append Growth ===========")

	grow := []int{}

	for i := 1; i <= 10; i++ {
		grow = append(grow, i)
		fmt.Printf("Value Added: %d | len=%d cap=%d\n", i, len(grow), cap(grow))
	}

	fmt.Println("\n=========== For Loop ===========")

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}