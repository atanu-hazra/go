package main

import "fmt"

func main() {
	// 1. Expression Switch with multiple comma-separated case values.
	// Note: In Go, cases do NOT fall through by default (no explicit 'break' statement is needed).
	month := "January"
	switch month {
	case "January", "February":
		fmt.Println("Winter") // output: Winter
	case "March", "April", "May":
		fmt.Println("Spring") // output (if matched): Spring
	case "June", "July", "August":
		fmt.Println("Summer") // output (if matched): Summer
	case "September", "October", "November":
		fmt.Println("Autumn") // output (if matched): Autumn
	default:
		fmt.Println("Invalid month") // output (if default): Invalid month
	}

	// 2. Tagless / Expressionless Switch:
	// A switch without an expression is equivalent to 'switch true'.
	// It evaluates boolean conditions from top to bottom, making it a cleaner alternative to 'if-else if-else' chains.
	age := 23

	switch {
	case age < 13:
		fmt.Println("Child") // output (if matched): Child

	case age < 20:
		fmt.Println("Teenager") // output (if matched): Teenager

	case age < 60:
		fmt.Println("Adult") // output: Adult

	default:
		fmt.Println("Senior") // output (if default): Senior
	}
}
