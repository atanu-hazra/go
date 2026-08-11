package main

import "fmt"

func main()  {
	var nums [3]int

	nums[1] = 2

	fmt.Println(nums)

	// Initialize while declaring
	age := [3]int{10, 20, 30}
	random := [4]int{}
	fmt.Println(age, random)  // [10 20 30] [0 0 0 0]
	
	names := [3]string{}
	boolean := [3]bool{}

	fmt.Println(names, boolean)  // [  ] [false false false]
	fmt.Println(len(names), len(boolean))

	fmt.Printf("Variable type: %T \n", names) // [3]string 

	// for loops

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	for index, value := range nums {
		fmt.Println(index, value)
	}

	for _, value := range nums {
		fmt.Println(value)
	}

	// Arrays can be compared directly.

	a := [3]int{1,2,3}
	b := [3]int{1,2,3}

	fmt.Println(a == b)
}
