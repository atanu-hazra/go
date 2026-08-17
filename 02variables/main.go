package main

import "fmt"

// Package-level (global) variable.
// Note: Capitalized first letter (GloabalInt) means it is exported (public) to other packages.
var GloabalInt int = 43

func main() {
	// 1. Explicit String declaration: var <name> <type> = <value>
	var username string = "Nova"
	fmt.Println(username)                        // output: Nova
	fmt.Printf("Variable type: %T \n", username) // output: Variable type: string 

	// 2. Boolean declaration
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)                        // output: true
	fmt.Printf("Variable type: %T \n", isLoggedIn) // output: Variable type: bool 

	// 3. Integer declaration
	var age int = 24
	fmt.Println(age)                        // output: 24
	fmt.Printf("Variable type: %T \n", age) // output: Variable type: int 

	// 4. Short variable declaration (:=) - "Walrus operator"
	// Type is inferred automatically (float64 by default for floating numbers).
	// Note: := can only be used inside functions, not at package level.
	smallFloat := 334.4555
	fmt.Println(smallFloat)                        // output: 334.4555
	fmt.Printf("Variable type: %T \n", smallFloat) // output: Variable type: float64 

	// 5. Accessing package-level variable
	fmt.Println(GloabalInt)                        // output: 43
	fmt.Printf("Variable type: %T \n", GloabalInt) // output: Variable type: int 
}