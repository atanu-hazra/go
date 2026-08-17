package main

import (
	"bufio"   // Buffered I/O for user input
	"fmt"     // Formatted I/O
	"os"      // OS standard streams (os.Stdin)
	"strconv" // Type conversions between strings and basic data types
	"strings" // String manipulation (TrimSpace)
)

func main() {
	
	// basicConversions()
	conversions()
}

// Demonstrates reading user input from terminal and converting string to float
func basicConversions() {

	// 1. Create a buffered reader
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Rate our app (1-5):") // output: Rate our app (1-5):

	// 2. Read string input until newline
	rating, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading rating:", err) // output (on error): Error reading rating: <error_msg>
		return
	}

	// 3. strings.TrimSpace removes whitespace/newline before parsing.
	// strconv.ParseFloat converts string to float64 (bitSize: 64).
	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	if err != nil {
		fmt.Println("Error parsing input to float:", err) // output (on error): Error parsing input to float: <error_msg>
		return
	} else {
		// Arithmetic with float
		fmt.Println("Added 1 to your rating", numRating+1) // output: Added 1 to your rating 5 (e.g. if input was 4)
	}
}

// Demonstrates explicit type casting and string-to-int conversion
func conversions() {
	// Note: Go requires EXPLICIT type conversion; implicit coercion is not allowed.
	age := 24
	height := 5.11

	// 1. Convert int to float64 to perform float addition
	fmt.Println("sum", float64(age)+height) // output: sum 29.11

	// 2. Convert float64 to int (truncates/drops decimal part)
	fmt.Println(int(height)) // output: 5

	// 3. Convert string to integer using strconv.Atoi (ASCII to integer)
	strNum := "54"

	// Using the blank identifier '_' to ignore the error returned by Atoi
	num, _ := strconv.Atoi(strNum)

	fmt.Println(num) // output: 54
}
