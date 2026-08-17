package main

import "fmt"

// Function Syntax in Go:
// func functionName(param1 type1, param2 type2) (returnType1, returnType2) {
//     // body
//     return val1, val2
// }

func main() {
	// 1. Basic function call with single return value
	result := add(3, 5)
	sqrResult := square(11)
	fmt.Println(result, sqrResult) // output: 8 121

	// 2. Multiple return values with error handling (division by zero)
	divideResult, err1 := divide(12, 0)
	fmt.Println(divideResult, err1) // output: 0 cannot divide by zero

	// Successful division
	divideResult2, err2 := divide(12, 6)
	fmt.Println(divideResult2, err2) // output: 2 <nil>

	// 3. Multiple return values (sum and difference)
	resultSum, resultDiff := sumAndDiff(10, 5)
	fmt.Println(resultSum, resultDiff) // output: 15 5

	// 4. Variadic function call (accepts any number of int arguments)
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) // output: 55
}

// add takes two int parameters and returns a single int
func add(a int, b int) int {
	return a + b
}

// square calculates and returns n squared
func square(n int) int {
	return n * n
}

// divide demonstrates parameter type grouping (a, b int) and returning (result, error).
// In Go, errors are returned as values rather than throwing exceptions.
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil // nil represents no error occurred
}

// sumAndDiff returns multiple values of the same type: (sum, diff)
func sumAndDiff(a int, b int) (int, int) {
	return a + b, a - b
}

// sum is a variadic function: '...int' accepts zero or more int arguments as a slice ([]int)
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}
