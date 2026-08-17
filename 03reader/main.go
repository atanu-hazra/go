package main

import (
	"bufio"   // Package bufio implements buffered I/O (useful for reading lines of text)
	"fmt"     // Package fmt implements formatted I/O
	"os"      // Package os provides platform-independent interface to OS functionality (e.g. os.Stdin)
	"strconv" // Package strconv provides conversions to and from string representations of basic data types
	"strings" // Package strings provides functions to manipulate UTF-8 encoded strings
)

func main() {
	welcome := "Welcome to our app"
	fmt.Println(welcome) // output: Welcome to our app

	// Step 1: Create a buffered reader that reads from standard input (keyboard)
	reader := bufio.NewReader(os.Stdin)

	// Step 2: Prompt and read string input
	fmt.Println("Please enter your name:") // output: Please enter your name:

	// ReadString reads until the first occurrence of delimiter '\n' (Enter key)
	nameInput, err := reader.ReadString('\n')
	// Idiomatic Go error handling: always check if err != nil
	if err != nil {
		fmt.Println("Error reading name:", err) // output (on error): Error reading name: <error_msg>
		return
	}

	// taking age input
	fmt.Println("Please enter your age:")

	// ageInput, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error reading age:", err)
	// 	return
	// }

	// ageInput = strings.TrimSpace(ageInput)

	// age, err := strconv.Atoi(ageInput)
	// if err != nil {
	// 	fmt.Println("Invalid age:", err)
	// 	return
	// }

	
	age, err := ReadAge(reader)
	if err != nil {
		fmt.Println(err) // output (on error): <error_msg>
		return
	}

	// Step 4: Display the final collected user data
	fmt.Println("Thanks for using our app, ", nameInput, age) // output: Thanks for using our app,  <name> <age>
}

// ReadAge reads user input from a buffered reader, trims newlines, and parses it to an integer.
// It demonstrates Go's standard multiple return pattern: (result, error).
func ReadAge(reader *bufio.Reader) (int, error) {
	// Read until newline character
	input, err := reader.ReadString('\n')
	if err != nil {
		// %w verb wraps the original error for error chaining
		return 0, fmt.Errorf("reading age: %w", err)
	}

	// Trim trailing '\n' and whitespace so strconv can parse the pure number
	input = strings.TrimSpace(input)

	// Atoi stands for "ASCII to Integer"
	age, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("converting age: %w", err)
	}

	// Return parsed integer and nil error indicating success
	return age, nil
}