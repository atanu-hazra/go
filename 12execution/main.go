package main

import "fmt"

// Go Loop & Control Flow: for, break, continue, labeled loops, and goto.
// Note: 'for' is the ONLY looping construct in Go (no while or do-while).

func main() {
	// execution1()
	// execution2()
	execution3()
}

// Demonstrates standard for loops, while-style loops, range loops, break, and continue
func execution1() {
	// 1. Standard 3-component for loop: init; condition; post
	for i := 0; i < 10; i++ {
		fmt.Println(i) // output: 0 \n 1 \n 2 \n 3 \n 4 \n 5 \n 6 \n 7 \n 8 \n 9
	}

	// 2. Infinite loop (uncomment to run):
	// for {
	// 	fmt.Println("Infinite loop")
	// }

	// 3. While-style loop in Go: for <condition>
	i := 0
	for i < 5 {
		fmt.Println(i) // output: 0 \n 1 \n 2 \n 3 \n 4
		i++
	}

	// 4. for-range loop over a slice
	numbers := []int{10, 20, 30}
	for index, value := range numbers {
		fmt.Println(index, value) // output: 0 10 \n 1 20 \n 2 30
	}

	// 5. 'break' statement: immediately terminates loop when i == 5
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i) // output: 0 \n 1 \n 2 \n 3 \n 4
	}

	// 6. 'continue' statement: skips the rest of current iteration when i == 2
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i) // output: 0 \n 1 \n 3 \n 4
	}
}

// Demonstrates labeled break and labeled continue for nested loops
func execution2() {
	// Labeled break: terminates the labeled outer loop rather than just the inner loop
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outer // breaks out of the entire 'outer' loop
			}
			fmt.Println(i, j)
			// output: 0 0 \n 0 1 \n 0 2 \n 1 0
		}
	}

	// Labeled continue: skips to the next iteration of the specified outer loop
outer2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				continue outer2 // skips remaining inner loop iterations, proceeds to next 'i'
			}
			fmt.Println(i, j)
			// output: 0 0 \n 1 0 \n 2 0
		}
	}
}

// Demonstrates goto statement (unconditional jump to label within function)
func execution3() {
	i := 0

start:
	fmt.Println(i) // output: 0 (1st), 1 (2nd), 2 (3rd), 3 (4th), 4 (5th)

	i++

	if i < 5 {
		goto start // Jumps execution back to the 'start:' label
	}
}